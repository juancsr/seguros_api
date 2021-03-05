package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//CTX Context variable
var CTX = context.TODO()

var client *mongo.Client

// var collection *mongo.Collection

// init connection
func init() {
	var err error
	URI := "mongodb://juancsr:juancsr@localhost:27016/"
	clientOptions := options.Client().ApplyURI(URI)
	client, err = mongo.Connect(CTX, clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(CTX, nil)
	if err != nil {
		log.Fatal(err)
	}
}

//GetCollection cliente de db
func GetCollection(collectionName string) *mongo.Collection {
	collection := client.Database("seguros").Collection(collectionName)
	return collection
}

// CloseConnection cierra la conexión con la BD
func CloseConnection() {
	defer func() {
		if err := client.Disconnect(CTX); err != nil {
			panic(err)
		}
	}()
}
