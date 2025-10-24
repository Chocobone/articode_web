package repository

import (
	"context"
	"time"

	dbConfig "github.com/chocobone/articode_web/db/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ModelingRepositoryMongo struct {
	collection *mongo.Collection
}

type ModelingInfoResponse struct {
	LessorID     string    `bson:"lessor_id" json:"lessor_id"`
	TenantID     *string   `bson:"tenant_id,omitempty" json:"tenant_id,omitempty"`
	Title        string    `bson:"title" json:"title"`
	Address      string    `bson:"address" json:"address"`
	USDZFileURL  string    `bson:"usdz_file_url" json:"usdz_file_url"`
	GLBFileURL   string    `bson:"glb_file_url" json:"glb_file_url"`
	ThumbnailURL *string   `bson:"thumbnail_url,omitempty" json:"thumbnail_url,omitempty"`
	Category     *string   `bson:"category,omitempty" json:"category,omitempty"`
	Description  *string   `bson:"description,omitempty" json:"description,omitempty"`
	CreatedAt    time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at" json:"updated_at"`
}

func NewModelingRepository() ModelingRepository {
	return &ModelingRepositoryMongo{
		collection: dbConfig.ModelingCollection,
	}
}

func (r *ModelingRepositoryMongo) GetModelingInfo(ctx context.Context, modelingID string) (*ModelingInfoResponse, error) {
	var modelingInfoResponse ModelingInfoResponse
	err := r.collection.FindOne(ctx, bson.M{"modeling_id": modelingID}).Decode(&modelingInfoResponse)
	if err != nil {
		return nil, err
	}
	return &modelingInfoResponse, nil
}

// ...existing code...
// Adding 3d Model to MongoDB
func (r *ModelingRepositoryMongo) PostModelingInfo(ctx context.Context, model *ModelingInfoResponse) (*ModelingInfoResponse, error) {
	_, err := r.collection.InsertOne(ctx, model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

// ...existing code...
func (r *ModelingRepositoryMongo) DeleteModelingInfo(ctx context.Context, modelingID string) error {
	res, err := r.collection.DeleteOne(ctx, bson.M{"modeling_id": modelingID})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
