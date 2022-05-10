package webhook

type WebhookRequest struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	MessageId   int         `json:"message_id"`
	From        MessageFrom `json:"from"`
	MessageChat MessageChat `json:"chat"`
	Date        int         `json:"date"`
	Text        string      `json:"text"`
}

type MessageFrom struct {
	Id           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	LanguageCode string `json:"language_code"`
}

type MessageChat struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	ChatType  string `json:"type"`
}
