package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"outstagram/server/configs"
	"outstagram/server/middlewares"
)

func FollowAPIRouter(router *gin.RouterGroup) {
	followController, err := configs.InitializeFollowController()
	if err != nil {
		log.Fatal(err.Error())
	}

	router.Use(middlewares.VerifyToken)

	router.POST("/:followingID", followController.CreateFollow)

	router.DELETE("/:followingID", followController.RemoveFollow)
}

