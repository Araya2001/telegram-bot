package mongoclient

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Reader interface {
	GetOneDocument(connector Connector) ([]byte, error)
	GetMultipleDocuments(connector Connector) ([]byte, error)
}

type ReadDocument struct {
	Database     string
	Collection   string
	Keyword      string
	KeywordValue string
}

func (d ReadDocument) GetOneDocument(connector Connector) ([]byte, error) {
	client, closer, err := connector.GetConnection()
	if err != nil {
		return nil, err
	}
	defer closer()
	coll := client.Database(d.Database).Collection(d.Collection)
	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{d.Keyword, d.KeywordValue}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		log.Fatalf("No document was found with the keyword: %s and value: %s\n", d.Keyword, d.KeywordValue)
		return nil, err
	}
	jsonData, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func (d ReadDocument) GetMultipleDocuments(connector Connector) ([]byte, error) {
	client, closer, err := connector.GetConnection()
	if err != nil {
		return nil, err
	}
	defer closer()
	coll := client.Database(d.Database).Collection(d.Collection)

	cursor, err := coll.Find(context.TODO(), bson.D{{d.Keyword, d.KeywordValue}})
	if err == mongo.ErrNoDocuments {
		log.Fatalf("No document was found with the keyword: %s and value: %s\n", d.Keyword, d.KeywordValue)
		return nil, err
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatalf("No documents where found with the keyword: %s and value: %s\n", d.Keyword, d.KeywordValue)
		return nil, err
	}
	jsonData, err := json.Marshal(results)
	return jsonData, nil
}
