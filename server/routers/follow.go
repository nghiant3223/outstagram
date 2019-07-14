package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"outstagram/server/injection"
	"outstagram/server/middlewares"
)

func FollowAPIRouter(router *gin.Engine, routerGroup *gin.RouterGroup) {
	followController, err := injection.InitializeFollowController()
	if err != nil {
		log.Fatal(err.Error())
	}

	routerGroup.Use(middlewares.VerifyToken(true))

	routerGroup.POST("/:followingID", followController.CreateFollow)

	routerGroup.DELETE("/:followingID", followController.RemoveFollow)
}

