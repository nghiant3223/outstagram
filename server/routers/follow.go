package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"outstagram/server/injection"
	"outstagram/server/middlewares"
)

func FollowAPIRouter(router *gin.RouterGroup) {
	followController, err := injection.InitializeFollowController()
	if err != nil {
		log.Fatal(err.Error())
	}

	router.Use(middlewares.VerifyToken(true))

	router.POST("/:followingID", followController.CreateFollow)

	router.DELETE("/:followingID", followController.RemoveFollow)
}

