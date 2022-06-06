package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"log"
	"os"
	"telegram-test-bot/pkg/dt/webhook"
	"telegram-test-bot/pkg/service/message"
	"telegram-test-bot/pkg/service/mongoclient"
)
import "net/http"

func main() {
	envFile := os.Getenv("ENV_FILE")
	if err := godotenv.Load(envFile); err != nil {
		log.Println("No .env file found")
	}
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/webhook", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "recieved"})
		jsonData, err := ioutil.ReadAll(c.Request.Body)

		if err != nil {
			fmt.Println(err)
		}

		data := webhook.WebhookRequest{}

		_ = json.Unmarshal(jsonData, &data)

		fmt.Println(data)

		if data.Message.Text == "Dar Respuesta" {
			msgreq := message.Setter(message.SendMessageRequest{
				ChatId:           2090444260,
				Text:             "Dando respuesta",
				ReplyToMessageId: data.Message.MessageId,
			})
			fmt.Println(msgreq)
			msg := message.Sender(message.SendMessageResponse{})
			_, err := msg.SendMessage(msgreq)
			if err != nil {
				panic(err)
			}
		}

	})

	r.GET("/hello-world", func(c *gin.Context) {

		msgreq := message.Setter(message.SendMessageRequest{
			ChatId: 2090444260,
			Text:   "Hello World",
		})
		mc := mongoclient.Connection{}
		mw := mongoclient.Writer(mongoclient.WriteDocument{
			Database:   "test_document",
			Collection: "test_collection",
			SingleData: bson.D{{"test_key", "test_value"}},
		})
		result, err := mw.SetOneDocument(mc)
		fmt.Println(result)
		fmt.Println(msgreq)
		msg := message.Sender(message.SendMessageResponse{})
		sendMessage, err := msg.SendMessage(msgreq)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": sendMessage},
		)

	})
	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
