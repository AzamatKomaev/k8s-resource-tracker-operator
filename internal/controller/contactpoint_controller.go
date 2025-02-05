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
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	tgv1 "github.com/AzamatKomaev/k8s-resource-tracker-operator/api/v1"
)

// ContactPointReconciler reconciles a ContactPoint object
type ContactPointReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=tg.azamaton.ru,resources=contactpoints,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=tg.azamaton.ru,resources=contactpoints/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=tg.azamaton.ru,resources=contactpoints/finalizers,verbs=update
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch

func (r *ContactPointReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	var contactPoint tgv1.ContactPoint

	if err := r.Get(ctx, req.NamespacedName, &contactPoint); err != nil {
		log.V(1).Info("ContactPoint was deleted")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log.V(1).Info("Namespace: " + contactPoint.Namespace)
	log.V(1).Info("Type: " + string(contactPoint.Spec.Type))

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ContactPointReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&tgv1.ContactPoint{}).
		Named("contactpoint").
		Complete(r)
}
