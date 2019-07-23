package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"outstagram/server/injection"
	"outstagram/server/middlewares"
)

func RoomAPIRouter(router *gin.Engine, routerGroup *gin.RouterGroup) {
	roomController, err := injection.InitializeRoomController()
	if err != nil {
		log.Fatal(err.Error())
	}

	routerGroup.Use(middlewares.VerifyToken(true))

	routerGroup.GET("/", roomController.GetRecentRooms)
	routerGroup.POST("/", roomController.CreateRoom)

	routerGroup.GET("/:idOrUsername", roomController.CheckRoomExist, roomController.GetRoom)

	routerGroup.GET("/:idOrUsername/messages", roomController.CheckRoomExist, roomController.GetRoomMessages)
	routerGroup.POST("/:idOrUsername/messages", roomController.CheckRoomExist, roomController.CreateMessage)
}
