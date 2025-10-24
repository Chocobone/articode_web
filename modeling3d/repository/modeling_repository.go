package repository

import (
	"context"
)

type ModelingRepository interface {
	//Modeling methods
	GetModelingInfo(ctx context.Context, ModelingID int) (*ModelingInfoResponse, error)

	//add Modeling's model
	PostModelingInfo(ctx context.Context, Modeling *ModelingInfoResponse) (*ModelingInfoResponse, error)

	// delete Modeling info
	DeleteModelingInfo(ctx context.Context, ModelingID int) error
}
