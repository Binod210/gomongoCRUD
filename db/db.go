package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func db(connString string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(connString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return client, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return client, err
	}
	log.Println("Successfully connected to MongoDB!")
	return client, nil
}
