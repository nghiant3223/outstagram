package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StoryAPIRouter(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Story")
	})
}
