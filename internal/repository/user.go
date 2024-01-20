package repository

import (
	"context"
	"errors"
	"math/rand"

	"github.com/gotrika/gotrika_backend/internal/core"
	"github.com/gotrika/gotrika_backend/internal/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type UserRepo struct {
	collection *mongo.Collection
}

func NewUserRepo(db *mongo.Database) *UserRepo {
	return &UserRepo{
		collection: db.Collection(core.UserCollectioName),
	}
}

func generateUserSign(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (r *UserRepo) CreateUser(ctx context.Context, userDTO *dto.CreateUserDTO) (string, error) {
	var user core.User
	err := r.collection.FindOne(ctx, bson.M{"username": userDTO.Username}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			res, err := r.collection.InsertOne(ctx, core.User{
				Username: userDTO.Username,
				Password: userDTO.HashedPassword,
				IsActive: userDTO.IsActive,
				IsAdmin:  userDTO.IsAdmin,
				Name:     userDTO.Name,
				Sign:     generateUserSign(25),
			})
			if err != nil {
				return "", err
			}
			userID := res.InsertedID.(primitive.ObjectID)
			return userID.Hex(), nil
		} else {
			return "", core.ErrUserAlreadyExists
		}
	}
	return "", core.ErrUserAlreadyExists
}

func (r *UserRepo) GetUserById(ctx context.Context, userID string) (*core.User, error) {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("invalid id param format")
	}
	var user core.User
	if err := r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetUserByUsername(ctx context.Context, username string) (*core.User, error) {
	var user core.User
	if err := r.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
