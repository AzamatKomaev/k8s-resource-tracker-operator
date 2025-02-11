/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"errors"
	"github.com/AzamatKomaev/k8s-resource-tracker-operator/internal/alert"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"strconv"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	tgv1 "github.com/AzamatKomaev/k8s-resource-tracker-operator/api/v1"
)

// TrackedFieldReconciler reconciles a TrackedField object
type TrackedFieldReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=tg.azamaton.ru,resources=trackedfields,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=tg.azamaton.ru,resources=trackedfields/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=tg.azamaton.ru,resources=trackedfields/finalizers,verbs=update
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch
// +kubebuilder:rbac:groups=apps,resources=deployments/status,verbs=get
// +kubebuilder:rbac:groups="",resources=services,verbs=get;list;watch
// +kubebuilder:rbac:groups="",resources=services/status,verbs=get

func (r *TrackedFieldReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	addNewStatus := func(trackedField *tgv1.TrackedField, action tgv1.ActionType, value string) error {
		newStatus := tgv1.TrackedFieldStatus{
			Time:   &metav1.Time{Time: time.Now()},
			Action: action,
			Value:  value,
		}
		trackedField.Status = append(trackedField.Status, newStatus)

		if err := r.Update(ctx, trackedField); err != nil {
			log.Error(err, "unable to update TrackedField status")
			return err
		}

		return nil
	}

	// CP для отправки уведомления в чат (tg/webhook/etc)
	var contactPointService alert.ContactPoint

	// кастомный ресурс, который нужно привести к желаемому состоянию
	var trackedField tgv1.TrackedField
	// кастомный ресурс ContactPoint
	var contactPoint tgv1.ContactPoint
	// секрет с API Token для авторизации запросов на отправку уведомлении
	var secretWithAPIToken corev1.Secret
	// массив со статусом ресурса TrackedField
	var trackedFieldStatus []tgv1.TrackedFieldStatus
	// ресурс K8s который необходимо отслеживать
	var targetResource interface{}
	// значение отслеживаемого поля
	var valueOfTrackedField string

	if err := r.Get(ctx, req.NamespacedName, &trackedField); err != nil {
		log.V(1).Info("TrackedField was deleted")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	targetContactPointLocation := types.NamespacedName{
		Name:      trackedField.Spec.ContactPoint,
		Namespace: req.Namespace,
	}

	if err := r.Get(ctx, targetContactPointLocation, &contactPoint); err != nil {
		log.V(1).Info("ContactPoint is not found by spec")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	targetResourceLocation := types.NamespacedName{
		Name:      trackedField.Spec.Target.Name,
		Namespace: trackedField.Spec.Target.Namespace,
	}

	targetSecretLocation := types.NamespacedName{
		Name:      contactPoint.Spec.ApiToken.SecretName,
		Namespace: "habr-k8s-resource-tracker-system",
	}

	if !contactPoint.Status.Ready {
		log.V(1).Info("Contact point is not ready")
		return ctrl.Result{RequeueAfter: time.Minute}, nil
	}

	if err := r.Get(ctx, targetSecretLocation, &secretWithAPIToken); err != nil {
		log.Error(err, "unable to get secret with api token by spec")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	switch trackedField.Spec.Target.Kind {
	case tgv1.DeploymentKind:
		targetResource = &v1.Deployment{}
	case tgv1.ServiceKind:
		targetResource = &corev1.Service{}
	}

	contactPointService = alert.GetContactPointByType(contactPoint,
		string(secretWithAPIToken.Data[contactPoint.Spec.ApiToken.Key]))

	obj, isOk := targetResource.(client.Object)

	if !isOk {
		return ctrl.Result{}, errors.New("cannot cast target kind to resource")
	}

	trackedFieldStatus = trackedField.Status

	if err := r.Get(ctx, targetResourceLocation, obj); err != nil {
		log.Error(err, "unable to get resource by TrackedField spec")
		lastStatus := trackedFieldStatus[len(trackedFieldStatus)-1]

		if lastStatus.Action == tgv1.ResourceDeleted {
			return ctrl.Result{RequeueAfter: time.Minute}, nil
		}

		addNewStatus(&trackedField, tgv1.ResourceDeleted, "")
		_, err = contactPointService.SendAlert("Resource kind: " + string(trackedField.Spec.Target.Kind) + "\n" +
			"Name: " + trackedField.Spec.Target.Name + "\n" +
			"Status: " + string(tgv1.ResourceDeleted) + "\n" +
			"Value: " + "-1")
		return ctrl.Result{RequeueAfter: time.Minute}, nil
	}

	switch trackedField.Spec.Field {
	case "replicas":
		valueOfTrackedField = strconv.Itoa(int(*obj.(*v1.Deployment).Spec.Replicas))
	case "image":
		valueOfTrackedField = obj.(*v1.Deployment).Spec.Template.Spec.Containers[0].Image
	case "targetPort":
		valueOfTrackedField = obj.(*corev1.Service).Spec.Ports[0].TargetPort.StrVal
	}

	if len(trackedFieldStatus) < 1 ||
		(trackedFieldStatus[len(trackedFieldStatus)-1].Value != valueOfTrackedField) {

		action := tgv1.ResourceCreated
		if len(trackedFieldStatus) > 0 {
			lastStatus := trackedFieldStatus[len(trackedFieldStatus)-1]
			if lastStatus.Action != tgv1.ResourceDeleted {
				action = tgv1.ResourceUpdated
			}
		}

		err := addNewStatus(&trackedField, action, valueOfTrackedField)
		if err != nil {
			return ctrl.Result{}, err
		}

		_, err = contactPointService.SendAlert("Resource kind: " + string(trackedField.Spec.Target.Kind) + "\n" +
			"Name: " + trackedField.Spec.Target.Name + "\n" +
			"Status: " + string(action) + "\n" +
			"Value: " + valueOfTrackedField)
		if err != nil {
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{RequeueAfter: time.Minute}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *TrackedFieldReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&tgv1.TrackedField{}).
		Named("trackedfield").
		Complete(r)
}
