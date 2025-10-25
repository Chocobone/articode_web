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
	
	// query parameter
	title := c.Query("title")
	address := c.Query("address")
	category := c.Query("category")

	// call service func
	models, err := h.service.GetModelingInfo(ctx, title, address, category)
	if err != nil {
		util.RespondInternalError(c, err.Error())
		return
	}

	util.RespondSuccess(c, models)
}

// POST /api/users/3d
func (h *ModelingHandler) PostModelingInfo(c *gin.Context) {
	ctx := c.Request.Context()

	var newModel repository.ModelingInfoResponse
	if err := c.ShouldBindJSON(&newModel); err != nil {
		util.RespondBadRequest(c, "Invalid request body") // 400
		return
	}

	created, err := h.service.PostModelingInfo(ctx, &newModel)
	if err != nil {
		util.RespondInternalError(c, err.Error()) // 500
		return
	}

	util.RespondSuccess(c, created)
}

// DELETE /api/models/:id
func (h *ModelingHandler) DeleteModelingInfo(c *gin.Context) {
	ctx := c.Request.Context()

	modelingID := c.Param("id")
	if modelingID == "" {
		util.RespondBadRequest(c, "Bad Request") //400 error
		return
	}

	userID, err := util.ExtractUserIDFromJWT(c.Request)
	if err != nil {
		util.RespondUnauthorized(c, "Invaild token") // 401 error
		return
	}

	// call service
	err := h.service.DeleteModelingInfo(ctx, modelingID, userID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			util.RespondNotFound(c, "Model not found or not authorized") // id는 잘 들어왔지만 model을 못찾았거나 로그인된 이용자의 id랑 맞지 않을 때
			return
		}
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
