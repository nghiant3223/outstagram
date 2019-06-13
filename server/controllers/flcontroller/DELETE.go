package flcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"outstagram/server/constants"
	"outstagram/server/utils"
)

func (fc *Controller) RemoveFollow(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs verifyToken middleware")
	}

	followingID, err := utils.StringToUint(c.Param("followingID"))
	if err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid parameter", err.Error())
		return
	}

	if err := fc.userService.Unfollow(userID, followingID); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.ResponseWithError(c, http.StatusNotFound, "User not found", err.Error())
			return
		}

		if err.Error() == constants.NotExisted {
			utils.ResponseWithError(c, http.StatusNotFound, "Not followed yet", err.Error())
			return
		}

		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while following user", err.Error())
		return
	}

	utils.ResponseWithSuccess(c, http.StatusCreated, "Follow user successfully", nil)
}

