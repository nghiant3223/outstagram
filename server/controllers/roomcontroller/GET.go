package roomcontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"outstagram/server/dtos/roomdtos"
	"outstagram/server/models"
	"outstagram/server/utils"
)

func (rc *Controller) CheckRoomExist(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs VerifyToken middleware")
	}

	var room *models.Room
	var err error

	idOrUsername, err := utils.StringToUint(c.Param("idOrUsername"))
	if err != nil {
		room, err = rc.roomService.GetRoomByPartnerUsername(userID, c.Param("idOrUsername"))
	} else {
		room, err = rc.roomService.GetRoomByID(idOrUsername)
	}

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.AbortRequestWithError(c, http.StatusNotFound, "Room not found", err.Error())
			return
		}

		utils.AbortRequestWithError(c, http.StatusInternalServerError, "Error while fetching recent rooms", err.Error())
		return
	}

	if room == nil {
		fmt.Println("> here")
		utils.AbortRequestWithError(c, http.StatusNotFound, "2 users' room not found", &gin.H{"type": "room_not_created"})
		return
	}

	c.Set("room", room)
}

func (rc *Controller) GetRoom(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs VerifyToken middleware")
	}

	sRoom, ok := c.Get("room")
	if !ok {
		log.Fatal("This route needs CheckRoomExist middleware")
	}

	room, _ := sRoom.(*models.Room)
	dtoRoom := room.ToDTO(userID)
	utils.ResponseWithSuccess(c, http.StatusOK, "Fetching recent rooms successfully", dtoRoom)
}

func (rc *Controller) GetRecentRooms(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs VerifyToken middleware")
	}

	var res roomdtos.GetRecentRoomResponse

	rooms, err := rc.roomService.GetRecentRooms(userID)
	if err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while fetching recent rooms", err.Error())
		return
	}

	for _, room := range rooms {
		// Set contact name to the name of the partner
		dtoRoom := room.ToDTO(userID)
		res.Rooms = append(res.Rooms, dtoRoom)
	}

	utils.ResponseWithSuccess(c, http.StatusOK, "Fetching recent rooms successfully", res)
}

func (rc *Controller) GetRoomMessages(c *gin.Context) {
	sRoom, ok := c.Get("room")
	if !ok {
		log.Fatal("This route needs CheckRoomExist middleware")
	}

	var req roomdtos.GetMessagesRequest
	var res roomdtos.GetMessageResponse

	var foundRoom *models.Room
	var err error

	if err := c.ShouldBindQuery(&req); err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid query parameter", err.Error())
		return
	}

	room, _ := sRoom.(*models.Room)
	// If limit and offset are not specified
	if req.Offset == 0 && req.Limit == 0 {
		foundRoom, err = rc.roomService.GetRoomMessages(room.ID)
	} else {
		foundRoom, err = rc.roomService.GetRoomMessagesWithLimit(room.ID, req.Limit, req.Offset)
	}

	if err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while fetching recent rooms", err.Error())
		return
	}

	for _, message := range foundRoom.Messages {
		res.Messages = append(res.Messages, message.ToDTO())
	}

	res.RoomID = foundRoom.ID
	utils.ResponseWithSuccess(c, http.StatusOK, "Fetching recent rooms successfully", res)
}
