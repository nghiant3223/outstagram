package roomcontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"outstagram/server/dtos/roomdtos"
	"outstagram/server/models"
	"outstagram/server/utils"
)

func (rc *Controller) CreateRoom(c *gin.Context) {
	var req roomdtos.CreateRoomRequest
	var res roomdtos.CreateRoomResponse

	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs VerifyToken middleware")
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid argument", err.Error())
		return
	}

	if len(req.MemberIDs) != 1 {
		fmt.Println(req.MemberIDs)
		utils.ResponseWithError(c, http.StatusNotImplemented, "Number of member is not supported yet", nil)
		return
	}

	room, err := rc.roomService.CreateDualRoom(req.MemberIDs[0], userID)
	if err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Cannot create room", err.Error())
		return
	}

	if req.FirstMessage != nil {
		fmt.Println("created")
		message := models.Message{Content: req.FirstMessage.Content, Type: req.FirstMessage.Type, UserID: userID}
		err := rc.roomService.CreateMessage(room.ID, &message)
		if err != nil {
			utils.ResponseWithError(c, http.StatusNotImplemented, "Cannot create message", err.Error())
			return
		}
	}

	createdRoom, _ := rc.roomService.GetRoomByID(room.ID)
	dtoRoom := createdRoom.ToDTO(userID)
	res.Room = dtoRoom
	utils.ResponseWithSuccess(c, http.StatusCreated, "Room created", res)
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
	if err := rc.roomService.CreateMessage(room.ID, &message); err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Cannot save message", err.Error())
		return
	}
	res.Message = message.ToDTO()
	utils.ResponseWithSuccess(c, http.StatusCreated, "Create message successfully", res)
}
