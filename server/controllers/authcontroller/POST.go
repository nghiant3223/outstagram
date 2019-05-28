package authcontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"outstagram/server/dtos"
	"outstagram/server/utils"
)

func (ac *Controller) Login(c *gin.Context) {

}

func (ac *Controller) Register(c *gin.Context) {
	var reqBody dtos.RegisterRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		utils.AbortRequestWithError(c, http.StatusBadRequest, "Some required fields is missing")
		return
	}
	// TODO: Create user in database
	//user := models.User{Username:reqBody.Username, Password: reqBody.Password, Email:reqBody.Email}
	utils.AbortRequestWithSuccess(c, http.StatusCreated, "Create user successfully", nil)
}
