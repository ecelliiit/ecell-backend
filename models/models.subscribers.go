package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Subscriber struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName *string            `json:"first_name,omitempty" bson:"first_name"`
	LastName  *string            `json:"last_name,omitempty" bson:"last_name"`
	Email     *string            `json:"email,omitempty" bson:"email"`
	Contacted bool               `json:"contacted" bson:"contacted,default:false"`
}
