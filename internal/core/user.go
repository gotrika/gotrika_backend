package core

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Name     string             `bson:"name"`
	Password string             `bson:"password"`
	Sign     string             `bson:"sign"`
	IsActive bool               `bson:"is_active"`
	IsAdmin  bool               `bson:"is_admin"`
}
