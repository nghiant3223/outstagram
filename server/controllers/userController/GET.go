package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) GetAllUser(c *gin.Context) {
}

func (uc *UserController) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	user, _ := uc.service.FindByUsername(username)
	c.String(http.StatusOK, user.Password)
}
