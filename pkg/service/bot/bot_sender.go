package bot

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"telegram-bot/pkg/service/message"
	"telegram-bot/pkg/service/mongoclient"
)

type Message struct {
	Message string `json:"message"`
	IdChat  int    `json:"idChat"`
}

type Sender interface {
	Send() error
}

func (m Message) Send() error {

	msgreq := message.Setter(message.SendMessageRequest{
		ChatId: m.IdChat,
		Text:   m.Message,
	})

	mc := mongoclient.Connection{}
	mw := mongoclient.Writer(mongoclient.WriteDocument{
		Database:   "telegram_bot",
		Collection: "telegram_bot_collection",
		SingleData: bson.D{{"chat_id", m.IdChat}, {"message", m.Message}},
	})
	_, err := mw.SetOneDocument(mc)
	if err != nil {
		log.Fatal("Error saving document: " + err.Error())
		return err
	}
	msg := message.Sender(message.SendMessageResponse{})
	resp, err := msg.SendMessage(msgreq)
	fmt.Println(resp)
	if err != nil {
		log.Fatal("Error saving document: " + err.Error())
		return err
	}
	return nil

}
