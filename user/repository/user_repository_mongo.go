package repository

import (
	"context"

	dbConfig "github.com/Chocobone/articode_web/v2/db/config"
	
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryMongo struct {
	collection *mongo.Collection
}

type UserInfoResponse struct {
	UserID		int	   `bson:"user_id"`
	Name        string `bson:"name"`
	Email       string `bson:"email"`
}

func NewUserRepository() UserRepository {
	return &UserRepositoryMongo{
		collection: dbConfig.UserCollection,
	}
}

func (r *UserRepositoryMongo) GetUserInfo(ctx context.Context, userID int) (*UserInfoResponse, error) {
	var userInfoResponse UserInfoResponse
	err := r.collection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&userInfoResponse)
	if err != nil {
		return nil, err
	}
	return &userInfoResponse, nil
}

