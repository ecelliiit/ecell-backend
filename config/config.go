package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type MongoCollection struct {
	Subscriber string
}

type Config struct {
	Secret     string
	PORT       string
	MongoURL   string
	Database   string
	Collection *MongoCollection
}

var Cfg *Config

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error in loading env from file: %v", err)
	}

	Cfg = &Config{
		Secret:   os.Getenv("SECRET"),
		MongoURL: os.Getenv("MONGO_URI"),
		PORT:     os.Getenv("PORT"),
		Database: os.Getenv("MONGO_DATABASE"),
		Collection: &MongoCollection{
			Subscriber: os.Getenv("MONGO_COLLECTION_SUBSCRIBER"),
		},
	}

	fmt.Printf("Env loaded successfully")
}
