package authcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"outstagram/server/dtos"
	"outstagram/server/utils"
)

func (ac *Controller) GetMe(c *gin.Context) {
	userID, _ := c.Get("userID")

	user, err := ac.userService.FindByID(utils.GetUserIDFromToken(userID))
	if gorm.IsRecordNotFoundError(err) {
		utils.ResponseWithError(c, http.StatusNotFound, "User not found", nil)
		return
	}

	result := dtos.GetMeResponse{
		Username:     user.Username,
		Fullname:     user.Fullname,
		Email:        user.Email,
		LastLogin:    user.LastLogin,
		StoryBoardID: user.StoryBoardID,
		NotifBoardID: user.NotifBoardID,
		Phone:        user.Phone,
		Gender:       user.Gender,
	}

	utils.ResponseWithSuccess(c, http.StatusOK, "Get user successfully", result)
}
