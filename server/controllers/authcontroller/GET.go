package authcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"outstagram/server/dtos"
	"outstagram/server/utils"
)

func (ac *Controller) GetMe(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatalf("This route need VerifyToken middleware")
	}

	user, err := ac.userService.FindByID(userID)
	if gorm.IsRecordNotFoundError(err) {
		utils.ResponseWithError(c, http.StatusNotFound, "User not found", nil)
		return
	}

	result := dtos.GetMeResponse{
		Username:     user.Username,
		Fullname:     user.Fullname,
		Email:        user.Email,
		LastLogin:    user.LastLogin,
		StoryBoardID: user.StoryBoard.ID,
		NotifBoardID: user.NotifBoard.ID,
		Phone:        user.Phone,
		Gender:       user.Gender,
		NumFollower:  len(ac.userService.GetFollowers(userID)),
		NumFollowing: len(ac.userService.GetFollowings(userID)),
	}

	utils.ResponseWithSuccess(c, http.StatusOK, "Get user successfully", result)
}
