package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"outstagram/server/configs"
	"outstagram/server/middlewares"
)

func PostAPIRouter(router *gin.RouterGroup) {
	postController, err := configs.InitializePostController()
	if err != nil {
		log.Fatal(err.Error())
	}

	router.Use(middlewares.VerifyToken)

	router.GET("/", postController.GetPosts)
	router.POST("/", postController.CreatePost)

	router.GET("/:postID/comments", postController.GetPostComments)
	router.POST("/:postID/comments", postController.CreatePostComment)

	router.GET("/:postID/comments/:cmtID/replies", postController.GetCommentReplies)
	router.POST("/:postID/comments/:cmtID/replies", postController.CreateCommentReply)

	router.GET("/:postID", postController.GetPost)

	router.POST("/:postID/views", postController.ViewPost)
}
