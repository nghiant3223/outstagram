package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"outstagram/server/injection"
	"outstagram/server/middlewares"
)

func ReactAPIRouter(router *gin.RouterGroup) {
	reactController, err := injection.InitializeReactController()
	if err != nil {
		log.Fatal(err.Error())
	}

	router.POST("/:rctableID", middlewares.VerifyToken, reactController.CreateReaction)

	router.DELETE("/:rctableID", middlewares.VerifyToken, reactController.RemoveReaction)
}
