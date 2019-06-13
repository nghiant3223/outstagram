package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"outstagram/server/configs"
)

func AuthAPIRouter(router *gin.RouterGroup) {
	authController, err := configs.InitializeAuthController()
	if err != nil {
		log.Fatal(err.Error())
	}

	router.POST("/register", authController.Register)
	router.POST("/login", authController.Login)
}
