package rctcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"outstagram/server/models"
	"outstagram/server/utils"
)

func (rc *Controller) CreateReaction(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs VerifyToken middleware")
	}

	reactableID, err := utils.StringToUint(c.Param("rctableID"))
	if err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid parameter", err.Error())
		return
	}

	if err := rc.checkUserAuthorizationForReactable(reactableID, userID); err != nil {
		utils.ResponseWithError(c, err.StatusCode, err.Message, err.Data)
		return
	}

	if reacted := rc.reactableService.CheckUserReaction(userID, reactableID); reacted {
		utils.ResponseWithError(c, http.StatusConflict, "Already reacted", nil)
		return
	}

	react := models.React{ReactableID: reactableID, UserID: userID}
	if err := rc.reactService.Save(&react); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.ResponseWithError(c, http.StatusNotFound, "UserID or ReactableID not found", err.Error())
			return
		}

		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while saving react", err.Error())
		return
	}

	utils.ResponseWithSuccess(c, http.StatusCreated, "Create react successfully", nil)
}
