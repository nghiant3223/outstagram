package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"outstagram/server/configs"
	"outstagram/server/middlewares"
)

func MeAPIRouter(router *gin.RouterGroup) {
	meController, err := configs.InitializeMeController()
	if err != nil {
		log.Fatal(err.Error())
	}

	router.Use(middlewares.VerifyToken)

	router.GET("/",  meController.GetMe)
	router.GET("/feed", meController.GetNewsFeed)
}