package utils

import (
	"github.com/gin-gonic/gin"
	"time"
)

func NewStringPointer(str string) *string {
	return &str
}

func NewTimePointer(t time.Time) *time.Time {
	return &t
}

func RetrieveUserID(c *gin.Context) (uint, bool) {
	userID, ok := c.Get("userID")
	if !ok {
		return 0, false
	}

	return uint(userID.(float64)), true
}
