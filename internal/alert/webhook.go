package alert

import "errors"

type WebhookContactPoint struct {
	URL        string
	APIToken   string
	HeaderName string
}

func (cp *WebhookContactPoint) SendAlert(message string) (interface{}, error) {
	return nil, errors.New("webhook alerting is not available yet")
}
