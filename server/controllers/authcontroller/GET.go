package authcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"outstagram/server/dtos/authdtos"
	"outstagram/server/utils"
)

func (ac *Controller) GetMe(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs verifyToken middleware")
	}

	user, err := ac.userService.FindByID(userID)
	if gorm.IsRecordNotFoundError(err) {
		utils.ResponseWithError(c, http.StatusNotFound, "User not found", nil)
		return
	}

	result := authdtos.GetMeResponse{
		FollowerCount:  len(ac.userService.GetFollowers(userID)),
		FollowingCount: len(ac.userService.GetFollowings(userID)),
		NotifBoardID:   user.NotifBoard.ID,
		StoryBoardID:   user.StoryBoard.ID,
	}
	if err := copier.Copy(&result, &user); err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while copying from model to reponse body", err.Error())
		return
	}

	utils.ResponseWithSuccess(c, http.StatusOK, "Fetch user successfully", result)
}
