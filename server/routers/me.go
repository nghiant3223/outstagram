package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"outstagram/server/injection"
	"outstagram/server/middlewares"
)

func MeAPIRouter(router *gin.RouterGroup) {
	meController, err := injection.InitializeMeController()
	if err != nil {
		log.Fatal(err.Error())
	}

	router.Use(middlewares.VerifyToken)

	router.GET("",  meController.GetMe)
	router.GET("/newsfeed", meController.GetNewsFeed)
	router.GET("/storyfeed", meController.GetStoryFeed)
}