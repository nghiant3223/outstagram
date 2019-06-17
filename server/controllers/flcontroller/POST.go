package flcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"outstagram/server/constants"
	"outstagram/server/utils"
)

func (fc *Controller) CreateFollow(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs verifyToken middleware")
	}

	followingID, err := utils.StringToUint(c.Param("followingID"))
	if err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid parameter", err.Error())
		return
	}

	if userID == followingID {
		utils.ResponseWithError(c, http.StatusConflict, "Cannot follow yourself", nil)
		return
	}

	if err := fc.userService.Follow(userID, followingID); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.ResponseWithError(c, http.StatusNotFound, "User not found", err.Error())
			return
		}

		if err.Error() == constants.AlreadyExist {
			utils.ResponseWithError(c, http.StatusNotFound, "Already followed", err.Error())
			return
		}

		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while following user", err.Error())
		return
	}

	utils.ResponseWithSuccess(c, http.StatusCreated, "Follow user successfully", nil)
}
