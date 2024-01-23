package utils

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConverIDtoObjectId(idParam string) (primitive.ObjectID, error) {
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return primitive.ObjectID{}, errors.New("invalid id param")
	}
	return id, nil
}
