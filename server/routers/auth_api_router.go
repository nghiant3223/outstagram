package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"outstagram/server/config"
	"outstagram/server/middlewares"
)

func AuthAPIRouter(router *gin.RouterGroup) {
	authController, err := config.InitializeAuthController()
	if err != nil {
		log.Fatal(err.Error())
	}

	router.GET("/me", middlewares.VerifyToken, authController.GetMe)

	router.POST("/register", authController.Register)
	router.POST("/login", authController.Login)
}
