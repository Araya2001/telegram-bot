package mongoclient

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Writer interface {
	SetOneDocument(connector Connector) ([]byte, error)
	SetMultipleDocuments(connector Connector) ([]byte, error)
}

type WriteDocument struct {
	Database     string
	Collection   string
	SingleData   bson.D
	MultipleData []interface{}
}

func (d WriteDocument) SetOneDocument(connector Connector) ([]byte, error) {
	client, closer, err := connector.GetConnection()
	if err != nil {
		return nil, err
	}
	defer closer()
	coll := client.Database(d.Database).Collection(d.Collection)
	if d.SingleData != nil {
		result, err := coll.InsertOne(context.TODO(), d.SingleData)
		if err == mongo.ErrNoDocuments {
			log.Fatalf("Couldn't Insert Document on Database: %s on collection: %s "+
				"with the following Document: %s \n", d.Database, d.Collection, d.SingleData)
			return nil, err
		}
		jsonData, err := json.Marshal(result)
		if err != nil {
			return nil, err
		}
		return jsonData, nil
	}
	log.Fatal("Couldn't Process Document Insertion request, please check your SingleData")
	return nil, err
}

func (d WriteDocument) SetMultipleDocuments(connector Connector) ([]byte, error) {
	client, closer, err := connector.GetConnection()
	if err != nil {
		return nil, err
	}
	defer closer()
	coll := client.Database(d.Database).Collection(d.Collection)
	if d.MultipleData != nil {
		result, err := coll.InsertMany(context.TODO(), d.MultipleData)
		if err == mongo.ErrNoDocuments {
			log.Fatalf("Couldn't Insert Documents on Database: %s on collection: %s "+
				"with the following Documents: %s \n", d.Database, d.Collection, d.MultipleData)
			return nil, err
		}
		jsonData, err := json.Marshal(result)
		if err != nil {
			return nil, err
		}
		return jsonData, nil
	}
	log.Fatal("Couldn't Process Document Insertion request, please check your MultipleData")
	return nil, err
}
