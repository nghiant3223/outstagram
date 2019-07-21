package routers

import (
	"log"
	"outstagram/server/injection"
	"outstagram/server/middlewares"

	"github.com/gin-gonic/gin"
)

func UserAPIRouter(router *gin.Engine, routerGroup *gin.RouterGroup) {
	userController, err := injection.InitializeUserController()
	if err != nil {
		log.Fatal(err.Error())
	}

	//userController.InitNewsfeed()

	routerGroup.Use(middlewares.VerifyToken(true))

	routerGroup.GET("/:userID", userController.GetUsersInfo)

	routerGroup.GET("/:userID/posts", userController.GetUserPosts)

	routerGroup.GET("/:userID/storyboard", userController.GetUserStoryBoard)
}
