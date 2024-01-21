package core

import "go.mongodb.org/mongo-driver/bson/primitive"

type Site struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty"`
	OwnerId     primitive.ObjectID   `bson:"owner_id,omitempty"`
	Name        string               `bson:"name"`
	URL         string               `bson:"url"`
	AccessUsers []primitive.ObjectID `bson:"access_users"`
	CounterCode string               `bson:"counter_code"`
}
