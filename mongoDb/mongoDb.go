package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Client *mongo.Client

func ConnectDB(uri string) error {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	Client = client
	return nil
}

func DisconnectDB() {
	if Client != nil {
		if err := Client.Disconnect(context.Background()); err != nil {
			log.Print(err)
		}
	}
}
