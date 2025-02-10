package alert

import (
	tgv1 "github.com/AzamatKomaev/k8s-resource-tracker-operator/api/v1"
	v1 "github.com/AzamatKomaev/k8s-resource-tracker-operator/api/v1"
)

func GetContactPointByType(cp v1.ContactPoint, apiToken string) ContactPoint {
	var contactPointService ContactPoint

	contactPointService = &TelegramContactPoint{
		ChatID:   cp.Spec.TelegramSpec.ChatId,
		APIToken: apiToken,
	}

	if cp.Spec.Type == tgv1.WebhookType {
		contactPointService = &WebhookContactPoint{
			URL:        cp.Spec.WebhookSpec.Url,
			APIToken:   apiToken,
			HeaderName: cp.Spec.WebhookSpec.HeaderName,
		}
	}

	return contactPointService
}
