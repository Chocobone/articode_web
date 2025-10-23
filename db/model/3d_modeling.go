package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Modeling3D struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	LessorID     string             `bson:"lessor_id" json:"lessor_id"`
	TenantID     *string            `bson:"tenant_id,omitempty" json:"tenant_id,omitempty"`
	Title        string             `bson:"title" json:"title"`
	Address      string             `bson:"address" json:"address"`
	USDZFileURL  string             `bson:"usdz_file_url" json:"usdz_file_url"`
	GLBFileURL   string             `bson:"glb_file_url" json:"glb_file_url"`
	ThumbnailURL *string            `bson:"thumbnail_url,omitempty" json:"thumbnail_url,omitempty"`
	Category     *string            `bson:"category,omitempty" json:"category,omitempty"`
	Description  *string            `bson:"description,omitempty" json:"description,omitempty"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
}
