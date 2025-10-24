package modeling3d

import (
	"context"
	"errors"
	"log"

	"github.com/chocobone/articode_web/db/model"
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

// 준형햄한테 물어보기
func (s *ModelingService) CreateModel(newModel model.Modeling3D) (*model.Modeling3D, error) {
	// Required Parameters validation
	if newModel.LessorID == "" || newModel.Title == "" || newModel.Address == "" || newModel.USDZFileURL == "" {
		return nil, errors.New("missing required fields")
	}

	// GLB File conversion logic

	// Saving to DB
	insertedModel, err := s.repo.InsertModel(newModel)
	if err != nil {
		return nil, err
	}

	return insertedModel, nil
}

func (s *ModelingService) GetModelingInfo(ctx context.Context, modelingID string) (*model.Modeling3D, error) {
	return s.repo.GetModelingInfo(ctx, modelingID)
}

func (s *ModelingService) PostModelingInfo(ctx context.Context, modeling *repository.ModelingRepository) (*model.Modeling3D, error) {
	return s.repo.GetModelingInfo(ctx, modeling)
}

func (s *ModelingService) DeleteModelingInfo(ctx context.Context, modelingID string) error {
	return s.repo.DeleteModelingInfo(ctx, modelingID)
}
