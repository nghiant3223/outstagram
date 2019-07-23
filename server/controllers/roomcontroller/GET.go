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

func (rc *Controller) GetRoom(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs verifyToken middleware")
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
		utils.ResponseWithError(c, http.StatusNotFound, "Room not found", nil)
		return
	}

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.ResponseWithError(c, http.StatusNotFound, "Room not found", err.Error())
			return
		}

		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while fetching recent rooms", err.Error())
		return
	}

	dtoRoom := room.ToDTO(userID)
	utils.ResponseWithSuccess(c, http.StatusOK, "Fetching recent rooms successfully", dtoRoom)
}

func (rc *Controller) GetRecentRooms(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs verifyToken middleware")
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
