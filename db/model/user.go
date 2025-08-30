package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID 		primitive.ObjectID	`bson:"_id,omitempty" json:"id,omitempty"`
	UserID  int					`bson:"user_id" json:"user_id"`
	Name	string 				`bson:"name" json:"name"`
	Email	string				`bson:"email" json:"email"`
}

