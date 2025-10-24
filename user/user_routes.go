package user

import (
	"github.com/gin-gonic/gin"
)

func GetUserRoutes(r *gin.Engine, userHandler *UserHandler) {
	//user routes
	r.GET("/api/users/", userHandler.GetUserInfo)
}

func PostUserRoutes(r *gin.Engine, userHandler *UserHandler) {
	r.POST("/api/users/", userHandler.PostUserInfo)
}

func DeleteUserRoutes(r *gin.Engine, userHandler *UserHandler) {
	r.DELETE("/api/users/", userHandler.DeleteUserInfo)
}
