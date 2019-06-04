package routers

import (
	"log"
	"outstagram/server/configs"

	"github.com/gin-gonic/gin"
)

func UserAPIRouter(router *gin.RouterGroup) {
	userController, err := configs.InitializeUserController()
	if err != nil {
		log.Fatal(err.Error())
	}

	router.GET("/:username", userController.GetUserPassword)
}
