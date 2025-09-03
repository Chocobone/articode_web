package user

import (
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, userHandler *UserHandler) {
	//user routes
	r.GET("/api/users/myinfo", userHandler.GetUserInfo)
}

func NewUserRoutes(r *gin.Engine, userHandler *UserHandler) {
	r.GET("/api/users/new", userHandler.NewUserHandler)
}
