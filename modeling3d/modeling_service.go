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
	
func (s *ModelingService) GetModelingInfo(ctx context.Context, modelingID string) (*repository.ModelingInfoResponse, error) {
	return s.repo.GetModelingInfo(ctx, modelingID)
}

func (s *ModelingService) PostModelingInfo(ctx context.Context, modeling *repository.ModelingInfoResponse) (*repository.ModelingInfoResponse, error) {
	if newModel.LessorID == "" || newModel.Title == "" || newModel.Address == "" || newModel.GLBFileURL == "" {
		return nil, errors.New("missing required fields")
	}
	
	// Saving to DB
	insertModel, err := s.repo.PostModelingInfo(ctx, modeling)
	if err != nil { return nil, err }

	return insertModel, nil
}

func (s *ModelingService) DeleteModelingInfo(ctx context.Context, modelingID string) error {
	return s.repo.DeleteModelingInfo(ctx, modelingID)
}
