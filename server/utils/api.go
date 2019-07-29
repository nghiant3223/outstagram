package utils

import (
	"github.com/gin-gonic/gin"
)

// AbortRequestWithError abort request with error, request stops at middleware in which this function is called
func AbortRequestWithError(c *gin.Context, code int, message string, data interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"status": "error", "message": message, "data": data})
}

// AbortRequestWithSuccess abort request with success, request stops at middleware in which this function is called
func AbortRequestWithSuccess(c *gin.Context, code int, message string, data interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"status": "success", "message": message, "data": data})
}

// ResponseWithError responses request with an error
func ResponseWithError(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, gin.H{"status": "error", "message": message, "description": data})
}

// ResponseWithSuccess responses request with a success
func ResponseWithSuccess(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, gin.H{"status": "success", "message": message, "data": data})
}
