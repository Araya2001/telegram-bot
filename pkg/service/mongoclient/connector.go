package mongoclient

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type Connection struct {
}

type Connector interface {
	GetConnection() (*mongo.Client, func(), error)
}

func (c Connection) GetConnection() (*mongo.Client, func(), error) {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, nil, err
	}
	return client, func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}, nil
}
