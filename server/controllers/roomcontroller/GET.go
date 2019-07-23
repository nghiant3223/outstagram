package roomcontroller

import (
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

	if room == nil {
		utils.AbortRequestWithError(c, http.StatusNotFound, "Room not found", nil)
		return
	}

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.AbortRequestWithError(c, http.StatusNotFound, "Room not found", err.Error())
			return
		}

		utils.AbortRequestWithError(c, http.StatusInternalServerError, "Error while fetching recent rooms", err.Error())
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

func (rc *Controller) GetMessages(c *gin.Context) {

}
