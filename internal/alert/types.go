package alert

type ContactPoint interface {
	SendAlert(string) (interface{}, error)
}

type TelegramSendMessageBody struct {
	Text   string `json:"text"`
	ChatID int    `json:"chat_id"`
}

type TelegramSendMessageResponse struct {
	Ok     bool              `json:"ok"`
	Result SendMessageResult `json:"result"`
}

type SendMessageResult struct {
	MessageID int         `json:"message_id"`
	From      interface{} `json:"from"`
	Chat      interface{} `json:"chat"`
	Date      int64       `json:"date"`
	Text      string      `json:"text"`
}
