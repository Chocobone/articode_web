package modeling3d

import (
	"github.com/gin-gonic/gin"
	"github.com/Chocobone/articode_web/v2/auth"
)

func RegisterRoutes(router *gin.Engine, handler *ModelingHandler) {
	api := router.Group("/api")
	api.Use(auth.JWTMiddleware()) // JWT Tokken check middleware

	{
		api.POST("/users/3d", handler.CreateModel) // API POST
	}
}