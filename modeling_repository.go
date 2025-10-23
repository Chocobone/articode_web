package modeling3d

import (
	"context"
	"time"

	"github.com/Chocobone/articode_web/v2/db/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ModelingRepository struct {
	Collection *mongo.Collection
}

// Adding 3d Model to MongoDB
func (r *ModelingRepository) InsertModel(newModel model.Modeling3D) (*model.Modeling3D, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.Collection.InsertOne(ctx, newModel)
	if err != nil {
		return nil, err
	}

	newModel.ID = result.InsertedID.(primitive.ObjectID)
	return &newModel, nil
}