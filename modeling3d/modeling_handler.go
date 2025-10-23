package modeling3d

import (
    "encoding/json"
    "net/http"
    "time"

    "github.com/Chocobone/articode_web/v2/db/model"
    "github.com/Chocobone/articode_web/v2/util"
)

type ModelingHandler struct {
	Service *ModelingService
}

// 3D Model Saving/Addition
func (h *ModelingHandler) CreateModel(w http.ResponseWriter, r *http.Request) {
	var newModel model.Modeling3D
	if err := json.NewDecoder(r.Body).Decode(&newModel); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Recording Creating&Updating time
	newModel.CreatedAt = time.Now()
	newModel.UpdatedAt = time.Now()

	// Calling Service hierarchy
	insertedModel, err := h.Service.CreateModel(newModel)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "Failed to save 3D model")
		return
	}

	// Success
	util.RespondWithJSON(w, http.StatusCreated, map[string]interface{}{
		"status": "success",
		"model": map[string]interface{}{
			"id":           insertedModel.ID.Hex(),
			"glb_file_url": insertedModel.GLBFileURL,
		},
	})
}