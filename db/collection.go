package db

import (
	"fmt"

	"github.com/ecelliiit/ecell-backend/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type Collection struct {
	Subscriber *mongo.Collection
}

var CollectionVar *Collection

func LoadCollections() *Collection {
	collection := &Collection{}
	collection.Subscriber = ClientVar.Database(config.Cfg.Database).Collection(config.Cfg.Collection.Subscriber)
	CollectionVar = collection
	fmt.Println("Collections Loaded")
	return collection
}
