package modeling3d

import (
	"time"

	"github.com/chocobone/articode_web/db/model"
	"github.com/chocobone/articode_web/util"
	"github.com/gin-gonic/gin"
)

type ModelingHandler struct {
	Service *ModelingService
}

// 3D Model Saving/Addition
func (h *ModelingHandler) CreateModel(c *gin.Context) {

	var newModel model.Modeling3D
	if err := c.ShouldBindJSON(&newModel); err != nil {
		util.RespondBadRequest(c, "Invalid request body")
		return
	}

	// Recording Creating&Updating time
	newModel.CreatedAt = time.Now()
	newModel.UpdatedAt = time.Now()

	// Calling Service hierarchy
	insertedModel, err := h.Service.CreateModel(newModel)
	if err != nil {
		util.RespondInternalError(c, "Failed to save 3D model")
		return
	}

	// Success
	util.RespondSuccess(c, map[string]interface{}{
		"status": "success",
		"model": map[string]interface{}{
			"id":           insertedModel.ID.Hex(),
			"glb_file_url": insertedModel.GLBFileURL,
		},
	})
}
