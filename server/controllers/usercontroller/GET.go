package usercontroller

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"outstagram/server/dtos/userdtos"
	"outstagram/server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (uc *Controller) GetAllUser(c *gin.Context) {
}

func (uc *Controller) GetUsersInfo(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("userID"), 10, 32)
	if err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid userID", err.Error())
		return
	}

	user, _ := uc.userService.FindByID(uint(userID))
	utils.ResponseWithSuccess(c, http.StatusOK, "Retrieve user's info successfully", user)
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
