package modeling3d

import (
	"errors"

	"github.com/chocobone/articode_web/db/model"
)

type ModelingService struct {
	Repo *ModelingRepository
}

func (s *ModelingService) CreateModel(newModel model.Modeling3D) (*model.Modeling3D, error) {
	// Required Parameters validation
	if newModel.LessorID == "" || newModel.Title == "" || newModel.Address == "" || newModel.USDZFileURL == "" {
		return nil, errors.New("missing required fields")
	}

	// GLB File conversion logic

	// Saving to DB
	insertedModel, err := s.Repo.InsertModel(newModel)
	if err != nil {
		return nil, err
	}

	return insertedModel, nil
}
