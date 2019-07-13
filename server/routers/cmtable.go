package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"outstagram/server/injection"
	"outstagram/server/middlewares"
)

func CommentableAPIRouter(router *gin.RouterGroup) {
	commentableController, err := injection.InitializeCommentableController()
	if err != nil {
		log.Fatal(err.Error())
	}

	router.GET("/:cmtableID/comments", middlewares.VerifyToken(true), commentableController.GetComments)
	router.POST("/:cmtableID/comments", middlewares.VerifyToken(true), commentableController.CreateComment)

	router.GET("/:cmtableID/comments/:cmtID/replies", middlewares.VerifyToken(true), commentableController.GetCommentReplies)
	router.POST("/:cmtableID/comments/:cmtID/replies", middlewares.VerifyToken(true), commentableController.CreateCommentReplies)
}
