package rctcontroller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"outstagram/server/dtos/rctdtos"
	"outstagram/server/utils"
)

func (rc *Controller) GetReactions(c *gin.Context) {
	audienceUserID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs verifyToken middleware")
	}

	reactableID, err := utils.StringToUint(c.Param("rctableID"))
	if err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid parameter", err.Error())
		return
	}

	var reqBody rctdtos.GetReactionsRequest
	var resBody rctdtos.GetReactionsResponse

	if err := c.ShouldBindQuery(&reqBody); err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid query parameter", err.Error())
		return
	}

	if reqBody.Offset == 0 && reqBody.Limit == 0 {
		resBody.Reactors = rc.reactableService.GetReactorDTOs(reactableID, audienceUserID, 10, 0)
	} else {
		resBody.Reactors = rc.reactableService.GetReactorDTOs(reactableID, audienceUserID, reqBody.Limit, reqBody.Offset)
	}

	resBody.ReactCount = rc.reactableService.GetReactCount(reactableID)

	utils.ResponseWithSuccess(c, http.StatusOK, "Fetch reactions successfully", resBody)
}
