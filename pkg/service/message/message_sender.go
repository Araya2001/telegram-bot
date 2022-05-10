package message

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Sender interface {
	SendMessage(setter Setter) (SendMessageResponse, error)
}

type SendMessageResponse struct {
	Ok     bool `json:"ok"`
	Result struct {
		MessageId int `json:"message_id"`
		From      struct {
			Id        int64  `json:"id"`
			IsBot     bool   `json:"is_bot"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
		} `json:"from"`
		Chat struct {
			Id        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date int    `json:"date"`
		Text string `json:"text"`
	} `json:"result"`
}

func (smr SendMessageResponse) SendMessage(setter Setter) (SendMessageResponse, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	jsonreq, err := json.Marshal(setter.GetSendMessageRequest())
	if err != nil {
		return SendMessageResponse{}, err
	}
	req, err := http.NewRequestWithContext(context.Background(),
		http.MethodPost, "https://api.telegram.org/bot5323227239:AAFHNeIOyWoup5pD0F8eqamVRLrjDhGsWz0/sendMessage",
		bytes.NewBuffer(jsonreq))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return SendMessageResponse{}, err
	}
	res, err := client.Do(req)

	fmt.Println(res)
	fmt.Println(setter.GetSendMessageRequest())
	if err != nil {
		return SendMessageResponse{}, err
	}
	jsonData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return SendMessageResponse{}, err
	}
	data := smr
	_ = json.Unmarshal(jsonData, &data)
	return data, nil
}
