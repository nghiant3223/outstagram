package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"outstagram/server/models"
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

func ResponseWithAppError(c *gin.Context, err *models.AppError) {
	c.JSON(err.Code, gin.H{"status": "error", "message": err.Message, "id": err.ID, "params": err.Params})
}

func ResponseWithBadRequest(c *gin.Context, additionalError error) {
	c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": additionalError.Error(), "id": "bad_request", "params": nil})
}

// ResponseWithSuccess responses request with a success
func ResponseWithSuccess(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, gin.H{"status": "success", "message": message, "data": data})
}
