package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"outstagram/server/injection"
	"outstagram/server/middlewares"
)

func StoryAPIRouter(router *gin.Engine, routerGroup *gin.RouterGroup) {
	storyController, err := injection.InitializeStoryController()
	if err != nil {
		log.Fatal(err.Error())
	}

	routerGroup.Use(middlewares.VerifyToken(true))

	routerGroup.GET("", storyController.GetMyStoryBoard)
	routerGroup.POST("", storyController.CreateStory)

	routerGroup.POST("/:storyID/views", storyController.ViewStory)
}
