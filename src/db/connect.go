package db

import (
	"context"
	"fmt"
	"go-todo-list/src/config"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Client {
	var MONGODB_URI = config.LoadConfig().MONGODB_URI

	var clientOptions = options.Client().ApplyURI(MONGODB_URI)

	var client, err = mongo.Connect(context.Background(), clientOptions)

	defer client.Disconnect(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to DB")

	return client
}
