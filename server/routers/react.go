package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"outstagram/server/injection"
	"outstagram/server/middlewares"
)

func ReactAPIRouter(router *gin.Engine, routerGroup *gin.RouterGroup) {
	reactController, err := injection.InitializeReactController()
	if err != nil {
		log.Fatal(err.Error())
	}

	routerGroup.GET("/:rctableID", middlewares.VerifyToken(true), reactController.GetReactions)
	routerGroup.POST("/:rctableID", middlewares.VerifyToken(true), reactController.CreateReaction)
	routerGroup.DELETE("/:rctableID", middlewares.VerifyToken(true), reactController.RemoveReaction)
}
