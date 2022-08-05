package repository

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/ecelliiit/ecell-backend/db"
	"github.com/ecelliiit/ecell-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateSubscriber(subscriber *models.Subscriber) (*models.Subscriber, error) {
	if !primitive.IsValidObjectID(subscriber.ID.Hex()) || subscriber.ID.IsZero() {
		subscriber.ID = primitive.NewObjectID()
	}
	subscriber.Contacted = false

	_, err := db.CollectionVar.Subscriber.InsertOne(context.TODO(), subscriber)
	if err != nil {
		fmt.Printf("Error in creating subscriber in DB: %v", err)
		return nil, err
	}

	return subscriber, nil
}

func GetAllSubscribers() (*[]models.Subscriber, error) {
	subscribers := []models.Subscriber{}
	filter := bson.M{}

	cursor, err := db.CollectionVar.Subscriber.Find(context.TODO(), filter)
	if err != nil {
		log.Printf("Error in finding all subscribers from database: %v", err)
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		subscriber := models.Subscriber{}
		if err := cursor.Decode(&subscriber); err != nil {
			log.Printf("Error in decoding subscriber %v", err)
			return nil, err
		}
		subscribers = append(subscribers, subscriber)
	}

	return &subscribers, nil
}

func GetSubscriberById(id primitive.ObjectID) (*models.Subscriber, error) {
	if !primitive.IsValidObjectID(id.Hex()) {
		return nil, errors.New("invalid Object Id")
	}

	subscriber := &models.Subscriber{}
	filter := bson.M{
		"_id": id,
	}

	result := db.CollectionVar.Subscriber.FindOne(context.TODO(), filter)
	if result.Err() == mongo.ErrNoDocuments {
		return nil, mongo.ErrNoDocuments
	}

	if err := result.Decode(subscriber); err != nil {
		fmt.Printf("Error in decoding result in GetSubscriberbyID: %v", err)
		return nil, err
	}

	return subscriber, nil
}

func DeleteSubscriberByID(id primitive.ObjectID) error {
	filter := bson.M{
		"_id": id,
	}
	_, err := db.CollectionVar.Subscriber.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("error in deleting subscriber from database")
		return err
	}
	return err
}

func MarkAsContacted(id primitive.ObjectID) (*models.Subscriber, error) {
	filter := bson.M{
		"_id": id,
	}
	update := bson.M{
		"$set": bson.M{
			"contacted": true,
		},
	}

	//updating the document with matching id
	result, err := db.CollectionVar.Subscriber.UpdateOne(context.TODO(), filter, update, nil)
	if err != nil {
		fmt.Printf("Error in marking subscriber id=%v as completed: %v", id.Hex(), err)
		return nil, err
	}
	if result.MatchedCount > 0 {
		fmt.Println("Document updated successfully now fetching the updated document")
	}

	//fetching updated document
	subscriber, err := GetSubscriberById(id)
	if err != nil {
		fmt.Printf("Error in fetching updated subscriber: %v", err)
		return nil, err
	}

	return subscriber, nil
}
