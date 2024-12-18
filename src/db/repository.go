package db

import "go.mongodb.org/mongo-driver/mongo"

var repository *mongo.Collection

func Repository() *mongo.Collection {
	var dbClient = Connect()

	repository = dbClient.Database("db").Collection("todos")

	return repository
}
