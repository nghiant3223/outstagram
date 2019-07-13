package routers

import (
	"log"
	"outstagram/server/injection"
	"outstagram/server/middlewares"

	"github.com/gin-gonic/gin"
)

func UserAPIRouter(router *gin.RouterGroup) {
	userController, err := injection.InitializeUserController()
	if err != nil {
		log.Fatal(err.Error())
	}

	// TODO: Handle this route for case that does not require verifyToken
	router.GET("/:userID", middlewares.VerifyToken(true), userController.GetUsersInfo)
	router.GET("/:userID/posts", middlewares.VerifyToken(false), userController.GetUsersInfo)
	router.GET("/:userID/storyboard", userController.GetUserStoryBoard)
}
