package routers

import (
	"outstagram/server/config"

	"github.com/gin-gonic/gin"
)

func UserAPIRouter(router *gin.RouterGroup) {
	userController, err := config.InitializeUserController()
	if err != nil {
		panic(err)
	}
	router.GET("/:username", userController.GetUserPassword)
}
