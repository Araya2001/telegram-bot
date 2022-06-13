package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"log"
	"os"
	"telegram-bot/pkg/dt/webhook"
	"telegram-bot/pkg/service/bot"
	"telegram-bot/pkg/service/message"
	"telegram-bot/pkg/service/mongoclient"
	"time"
)
import "net/http"

func main() {
	envFile := os.Getenv("ENV_FILE")
	if err := godotenv.Load(envFile); err != nil {
		log.Println("No .env file found")
	}
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		ExposeHeaders:    []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/", func(context *gin.Context) {

	})

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

	r.POST("/send-message", func(c *gin.Context) {
		jsonData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Fatal(err)
		}
		data := bot.Message{}
		_ = json.Unmarshal(jsonData, &data)
		fmt.Println(data)
		s := bot.Sender(bot.Message{
			Message: data.Message,
			IdChat:  data.IdChat,
		})
		err = s.Send()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Error sending message",
			})
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"message": "Message was sent succesfully",
		})
	})
	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
