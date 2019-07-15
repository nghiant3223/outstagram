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

	routerGroup.GET("/specific/:postID", postController.GetPost)
	routerGroup.PUT("/specific/:postID", postController.UpdatePost)
	routerGroup.POST("/", postController.CreatePost)

	routerGroup.GET("/images/:postImageID", postController.GetPostImage)
	routerGroup.PUT("/images/:postImageID", postController.UpdatePostImage)

	routerGroup.POST("/:postID/views", postController.ViewPost)
}
