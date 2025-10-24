package modeling3d

import (
	"github.com/gin-gonic/gin"
)

func RegisterModelRoutes(r *gin.Engine, modelHandler *ModelingHandler) {
	r.GET("/api/users/3d", modelHandler.CreateModel)
}
