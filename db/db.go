package db

import (
	"context"
	"fmt"
	"log"

	"github.com/ecelliiit/ecell-backend/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ClientVar *mongo.Client

func Connect() *mongo.Client {
	fmt.Printf("Connecting to the MongoDB server")

	var url string = config.Cfg.MongoURL
	clientOptions := options.Client().ApplyURI(url)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	ClientVar = client
	if err != nil {
		log.Printf(`Error in connecting to mongodb: %v`, err)
		return nil
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Printf(`Error in pinging to database: %v`, err)
		return nil
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

func Disconnect() {
	ClientVar.Disconnect(context.TODO())
}
