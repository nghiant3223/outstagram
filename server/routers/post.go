package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"outstagram/server/injection"
	"outstagram/server/middlewares"
)

func PostAPIRouter(router *gin.Engine, routerGroup *gin.RouterGroup) {
	postController, err := injection.InitializePostController()
	if err != nil {
		log.Fatal(err.Error())
	}

	routerGroup.Use(middlewares.VerifyToken(true))

	routerGroup.POST("/", postController.CreatePost)

	routerGroup.GET("/:postID", postController.GetPost)

	routerGroup.POST("/:postID/views", postController.ViewPost)
}
