package authcontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"outstagram/server/dtos"
	"outstagram/server/models"
	"outstagram/server/utils"
)

func (ac *Controller) Login(c *gin.Context) {

}

func (ac *Controller) Register(c *gin.Context) {
	var reqBody dtos.RegisterRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		utils.AbortRequestWithError(c, http.StatusBadRequest, "Some required fields is missing", nil)
		return
	}

	if ac.userService.CheckExistsByUsername(reqBody.Username) {
		utils.AbortRequestWithError(c, http.StatusConflict, "Username already used", nil)
		return
	}

	if ac.userService.CheckExistsByEmail(reqBody.Username) {
		utils.AbortRequestWithError(c, http.StatusConflict, "Email already used", nil)
		return
	}

	notifBoard := models.NotifBoard{}
	storyBoard := models.StoryBoard{}
	ac.nbService.Save(&notifBoard)
	ac.sbService.Save(&storyBoard)
	newUser := models.User{Username: reqBody.Username, Password: reqBody.Password, Email: reqBody.Email, NotifBoardID: notifBoard.ID, StoryBoardID: storyBoard.ID}

	if err := ac.userService.Save(&newUser); err != nil {
		utils.AbortRequestWithError(c, http.StatusInternalServerError, "Fail to create user", err.Error())
		return
	}

	utils.AbortRequestWithSuccess(c, http.StatusCreated, "Create user successfully", nil)
}
