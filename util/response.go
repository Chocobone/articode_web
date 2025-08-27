package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//200 OK
func RespondSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"data":	    data,
	})
}

//400 Bad Request
func RespondBadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"success":	false,
		"error":	message,
	})
}

//401 Unauthorized
func RespondUnauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"error":	message,
	})
}

//404 Not Found
func RespondNotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{
		"success":	false,
		"error":	message,
	})
}

//500 Interval Server Error
func RespondIntervalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusIntervalServerError, gin.H{
		"success":	false,
		"error":	message,
	})
}