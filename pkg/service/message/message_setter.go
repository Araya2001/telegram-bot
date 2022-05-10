package message

type Setter interface {
	GetSendMessageRequest() SendMessageRequest
	newSendMessageRequest() SendMessageRequest
}

type SendMessageRequest struct {
	ChatId                   int           `json:"chat_id"`
	Text                     string        `json:"text"`
	ParseMode                string        `json:"parse_mode"`
	Entities                 []interface{} `json:"entities"`
	DisableWebPagePreview    bool          `json:"disable_web_page_preview"`
	DisableNotification      bool          `json:"disable_notification"`
	ProtectContent           bool          `json:"protect_content"`
	ReplyToMessageId         int           `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool          `json:"allow_sending_without_reply"`
	ReplyMarkup              struct {
	} `json:"reply_markup"`
}

func (smr SendMessageRequest) GetSendMessageRequest() SendMessageRequest {
	return smr
}

func (smr SendMessageRequest) newSendMessageRequest() SendMessageRequest {

	return SendMessageRequest{
		ChatId:                   0,
		Text:                     "",
		ParseMode:                "",
		Entities:                 nil,
		DisableWebPagePreview:    false,
		DisableNotification:      false,
		ProtectContent:           false,
		ReplyToMessageId:         0,
		AllowSendingWithoutReply: false,
		ReplyMarkup:              struct{}{},
	}

}
