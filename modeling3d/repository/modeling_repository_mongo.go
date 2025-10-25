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
	TenantID     *string   `bson:"tenant_id,omitempty" json:"tenant_id,omitempty"` // No required
	Title        string    `bson:"title" json:"title"`
	Address      string    `bson:"address" json:"address"`
	USDZFileURL  string    `bson:"usdz_file_url" json:"usdz_file_url"` // No required
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

// Get 3D Model
func (r *ModelingRepositoryMongo) GetModelingInfo(ctx context.Context, title, address, category string) ([]*ModelingInfoResponse, error) {
	filter := bson.M{}
	if title != "" {
		filter["title"] = bson.M{"$regex": title, "$options": "i"} // case-free
	}
	if address != "" {
		filter["address"] = bson.M{"$regex": address, "$options": "i"}
	}
	if category != "" {
		filter["category"] = category
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []*ModelingInfoResponse
	for cursor.Next(ctx) {
		var model ModelingInfoResponse
		if err := cursor.Decode(&model); err != nil {
			return nil, err
		}
		results = append(results, &model)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

// Post 3D Model
func (r *ModelingRepositoryMongo) PostModelingInfo(ctx context.Context, model *ModelingInfoResponse) (*ModelingInfoResponse, error) {
	_, err := r.collection.InsertOne(ctx, model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

// Delete 3D Model
func (r *ModelingRepositoryMongo) DeleteModelingInfo(ctx context.Context, modelingID string, userID string) error {
	res, err := r.collection.DeleteOne(ctx, bson.M{
		"_id":			modelingID, //MongoDB 문서 ID
		"lessor_id":	userID, // JWT에서 가져온 로그인한 사용자의 ID
	})

	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil // 삭제가 정상적으로 되었으면 error = nil 반환
}
