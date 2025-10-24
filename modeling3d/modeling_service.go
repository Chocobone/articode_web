package modeling3d

import (
	"context"
	"log"

	"github.com/chocobone/articode_web/modeling3d/repository"
)

type ModelingService struct {
	repo repository.ModelingRepository
}

func NewModelingService(repo repository.ModelingRepository) *ModelingService {
	if repo == nil {
		log.Fatal("Modeling Repository cannot be nil")
	}
	return &ModelingService{repo: repo}
}

// 준형햄한테 물어보기 / POST metohd
// func (s *ModelingService) PostModelingInfo(newModel repository.ModelingInfoResponse) (*repository.ModelingInfoResponse, error) {
// 	ctx = c.Context
// 	// Required Parameters validation
// 	if newModel.LessorID == "" || newModel.Title == "" || newModel.Address == "" || newModel.USDZFileURL == "" {
// 		return nil, errors.New("missing required fields")
// 	}

// 	// GLB File conversion logic

// 	// Saving to DB
// 	insertedModel, err := s.repo.PostModelingInfo(newModel) //POST
// 	if err != nil {
// 		return nil, err
// 	}

// 	return insertedModel, nil
// }

func (s *ModelingService) GetModelingInfo(ctx context.Context, modelingID string) (*repository.ModelingInfoResponse, error) {
	return s.repo.GetModelingInfo(ctx, modelingID)
}

func (s *ModelingService) PostModelingInfo(ctx context.Context, modeling *repository.ModelingInfoResponse) (*repository.ModelingInfoResponse, error) {
	return s.repo.PostModelingInfo(ctx, modeling)
}

func (s *ModelingService) DeleteModelingInfo(ctx context.Context, modelingID string) error {
	return s.repo.DeleteModelingInfo(ctx, modelingID)
}
