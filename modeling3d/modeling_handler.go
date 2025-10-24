package modeling3d

import (
	"github.com/chocobone/articode_web/modeling3d/repository"
	"github.com/chocobone/articode_web/util"
	"github.com/gin-gonic/gin"
)

type ModelingHandler struct {
	service *ModelingService
}

func NewModelingHandler(service *ModelingService) *ModelingHandler {
	return &ModelingHandler{service: service}
}

// GET /api/models/:id
func (h *ModelingHandler) GetModelingInfo(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	if id == "" {
		util.RespondBadRequest(c, "Bad Request")
		return
	}

	m, err := h.service.GetModelingInfo(ctx, id)
	if err != nil {
		util.RespondInternalError(c, err.Error())
		return
	}

	util.RespondSuccess(c, m)
}

// POST /api/users/3d
func (h *ModelingHandler) PostModelingInfo(c *gin.Context) {
	ctx := c.Request.Context()

	var newModel repository.ModelingInfoResponse
	if err := c.ShouldBindJSON(&newModel); err != nil {
		util.RespondBadRequest(c, "Invalid request body")
		return
	}

	created, err := h.service.PostModelingInfo(ctx, &newModel)
	if err != nil {
		util.RespondInternalError(c, err.Error())
		return
	}

	util.RespondSuccess(c, created)
}

// DELETE /api/models/:id
func (h *ModelingHandler) DeleteModelingInfo(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	if id == "" {
		util.RespondBadRequest(c, "Bad Request")
		return
	}

	if err := h.service.DeleteModelingInfo(ctx, id); err != nil {
		util.RespondInternalError(c, err.Error())
		return
	}

	util.RespondSuccess(c, gin.H{"deleted": true})
}

// import (
// 	"time"

// 	"github.com/chocobone/articode_web/db/model"
// 	"github.com/chocobone/articode_web/util"
// 	"github.com/gin-gonic/gin"
// )

// type ModelingHandler struct {
// 	Service *ModelingService
// }

// // 3D Model Saving/Addition
// func (h *ModelingHandler) PostModelingInfo(c *gin.Context) {

// 	var newModel model.Modeling3D
// 	if err := c.ShouldBindJSON(&newModel); err != nil {
// 		util.RespondBadRequest(c, "Invalid request body")
// 		return
// 	}

// 	// Recording Creating&Updating time
// 	newModel.CreatedAt = time.Now()
// 	newModel.UpdatedAt = time.Now()

// 	// Calling Service hierarchy
// 	insertedModel, err := h.Service.PostModelingInfo(newModel)
// 	if err != nil {
// 		util.RespondInternalError(c, "Failed to save 3D model")
// 		return
// 	}

// 	// Success
// 	util.RespondSuccess(c, map[string]interface{}{
// 		"status": "success",
// 		"model": map[string]interface{}{
// 			"id":           insertedModel.ID.Hex(),
// 			"glb_file_url": insertedModel.GLBFileURL,
// 		},
// 	})
// }
