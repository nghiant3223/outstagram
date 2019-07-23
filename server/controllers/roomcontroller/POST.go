package roomcontroller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"outstagram/server/dtos/roomdtos"
	"outstagram/server/models"
	"outstagram/server/utils"
)

func (rc *Controller) CreateRoom(c *gin.Context) {

}

func (rc *Controller) CreateMessage(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs VerifyToken middleware")
	}

	sRoom, ok := c.Get("room")
	if !ok {
		log.Fatal("This route needs CheckRoomExist middleware")
	}

	var req roomdtos.CreateMessageRequest
	var res roomdtos.CreateMessageResponse

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid message", err.Error())
		return
	}

	room, _ := sRoom.(*models.Room)
	message := models.Message{UserID: userID, Content: req.Content, Type: req.Type}
	if err := rc.roomService.CreateMessage(room.ID, &message); err != nil{
		utils.ResponseWithError(c, http.StatusInternalServerError, "Cannot save message", err.Error())
		return
	}
	res.Message = message.ToDTO()
	utils.ResponseWithSuccess(c, http.StatusCreated, "Create message successfully", res)
}
