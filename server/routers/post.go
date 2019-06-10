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
	router.GET("/:postID", postController.GetPost)
	router.GET("/:postID/comments", postController.GetPostComments)
	router.GET("/:postID/comments/:cmtID/replies", postController.GetCommentReplies)


	router.POST("/", postController.CreatePost)
	router.POST("/:postID/comments", postController.CreatePostComment)
	router.POST("/:postID/views", postController.ViewPost)
	router.POST("/:postID/comments/:cmtID/replies", postController.CreateCommentReply)
}
