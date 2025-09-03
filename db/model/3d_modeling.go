package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type modeling struct {
	ID 			primitive.ObjectID	`bson:"_id,omitempty" json:"id,omitempty"`
	Owner   	int					`bson:"owner" json:"owner"`
	ModelName	string 				`bson:"model_name" json:"model_name"`
	Directory	string				`bson:"directory" json:"directory"`
}