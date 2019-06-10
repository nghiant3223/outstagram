package usercontroller

import (
	"net/http"
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

	user, _ := uc.service.FindByID(uint(userID))
	utils.ResponseWithSuccess(c, http.StatusOK, "Retrieve user's info successfully", user)
}
