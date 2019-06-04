package postcontroller

import (
	"github.com/gin-gonic/gin"
	"outstagram/server/utils"
)

func (pc *Controller) GetMyPosts(c *gin.Context) {
	utils.ResponseWithSuccess(c, 200, "Fuck", nil)
}
