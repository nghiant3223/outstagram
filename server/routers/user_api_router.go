package routers

import (
	"log"
	"outstagram/server/config"

	"github.com/gin-gonic/gin"
)

func UserAPIRouter(router *gin.RouterGroup) {
	userController, err := config.InitializeUserController()
	if err != nil {
		log.Fatal(err.Error())
	}
	router.GET("/:username", userController.GetUserPassword)
}
