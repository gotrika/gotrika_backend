package repository

import "go.mongodb.org/mongo-driver/mongo"

type Repositories struct{}

// NewRepositories: ini all repos
func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{}
}
