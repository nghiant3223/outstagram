package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"outstagram/server/injection"
	"outstagram/server/middlewares"
)

func StoryAPIRouter(router *gin.RouterGroup) {
	storyController, err := injection.InitializeStoryController()
	if err != nil {
		log.Fatal(err.Error())
	}

	router.Use(middlewares.VerifyToken)

	router.GET("", storyController.GetMyStoryBoard)

	router.POST("", storyController.CreateStory)
	router.POST("/:storyID/views", storyController.ViewStory)
}
