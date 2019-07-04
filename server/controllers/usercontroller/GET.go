package usercontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"outstagram/server/dtos/userdtos"
	"outstagram/server/utils"
)

func (uc *Controller) GetAllUser(c *gin.Context) {
}

func (uc *Controller) GetUsersInfo(c *gin.Context) {
	username := c.Param("userID")
	if username == "" {
		utils.ResponseWithError(c, http.StatusBadRequest, "Username must be provided", nil)
		return
	}

	user, err := uc.userService.FindByUsername(username)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.ResponseWithError(c, http.StatusNotFound, "Not found", err.Error())
			return
		}

		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving story board", err.Error())
		return
	}

	authUserID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs verifyToken middleware")
	}

	var resBody userdtos.GetUserResponse

	resBody.ID = user.ID
	resBody.Fullname = user.Fullname
	resBody.Username = user.Username
	resBody.FollowerCount = len(uc.userService.GetFollowers(user.ID))
	resBody.FollowingCount = len(uc.userService.GetFollowings(user.ID))

	isMe := authUserID == user.ID
	if !isMe {
		ok, err := uc.userService.CheckFollow(authUserID, user.ID)
		if err != nil {
			utils.ResponseWithError(c, http.StatusOK, "Error while retrieving user", err.Error())
			return
		}

		resBody.Followed = utils.NewBoolPointer(ok)
	}

	resBody.IsMe = isMe
	utils.ResponseWithSuccess(c, http.StatusOK, "Retrieve user's info successfully", resBody)
}

func (uc *Controller) GetUserStoryBoard(c *gin.Context) {
	userID, err := utils.StringToUint(c.Param("userID"))
	if err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid userID", err.Error())
		return
	}

	var resBody userdtos.GetStoryBoardResponse

	userStoryBoardDTO, err := uc.storyBoardService.GetUserStoryBoardDTO(userID)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.ResponseWithError(c, http.StatusNotFound, "Not found", err.Error())
			return
		}

		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving story board", err.Error())
		return
	}

	userStoryBoardDTO.IsMy = false
	resBody.StoryBoard = userStoryBoardDTO
	utils.ResponseWithSuccess(c, http.StatusOK, "Get user's storyboard successfully", resBody)
}
