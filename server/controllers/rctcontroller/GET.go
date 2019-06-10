package rctcontroller

import "github.com/gin-gonic/gin"

func (rc *Controller) RemoveReaction(c *gin.Context) {
	c.String(200, "ok")
}
