package modeling3d

import (
	"github.com/gin-gonic/gin"
)

func GetModelRoutes(r *gin.Engine, modelHandler *ModelingHandler) {
	r.GET("/api/users/3d", modelHandler.GetModelingInfo)
}

func PostUserRoutes(r *gin.Engine, modelHandler *ModelingHandler) {
	r.POST("/api/users/3d", modelHandler.PostModelingInfo)
}

func DeleteUserRoutes(r *gin.Engine, modelHandler *ModelingHandler) {
	r.DELETE("/api/users/3d", modelHandler.DeleteModelingInfo)
}
