package alert

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const TELEGRAM_BASE_API_URL = "https://api.telegram.org/bot%s"

type TelegramContactPoint struct {
	ChatID   int
	APIToken string
}

func (cp *TelegramContactPoint) SendAlert(message string) (interface{}, error) {
	if cp == nil {
		return nil, errors.New("TelegramContactPoint receiver is nil")
	}

	telegramAPIUrl := fmt.Sprintf(TELEGRAM_BASE_API_URL+"/sendMessage", cp.APIToken)
	body, err := json.Marshal(TelegramSendMessageBody{Text: message, ChatID: cp.ChatID})

	if err != nil {
		return nil, errors.New("failed to marshal request body")
	}

	req, err := http.NewRequest("POST", telegramAPIUrl, bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.New("failed to create HTTP request")
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != 200 {
		fmt.Println("Status code is " + strconv.Itoa(resp.StatusCode))
		return nil, errors.New("HTTP request failed")
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, errors.New("failed to read response body")
	}

	sendMessageResponse := TelegramSendMessageResponse{}
	err = json.Unmarshal(responseBody, &sendMessageResponse)
	if err != nil {
		return nil, errors.New("failed to unmarshal response")
	}

	return sendMessageResponse, nil
}
