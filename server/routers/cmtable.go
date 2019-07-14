package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"outstagram/server/injection"
	"outstagram/server/middlewares"
)

func CommentableAPIRouter(router *gin.Engine, routerGroup *gin.RouterGroup) {
	commentableController, err := injection.InitializeCommentableController()
	if err != nil {
		log.Fatal(err.Error())
	}

	routerGroup.GET("/:cmtableID/comments", middlewares.VerifyToken(true), commentableController.GetComments)
	routerGroup.POST("/:cmtableID/comments", middlewares.VerifyToken(true), commentableController.CreateComment)

	routerGroup.GET("/:cmtableID/comments/:cmtID/replies", middlewares.VerifyToken(true), commentableController.GetCommentReplies)
	routerGroup.POST("/:cmtableID/comments/:cmtID/replies", middlewares.VerifyToken(true), commentableController.CreateCommentReplies)
}
