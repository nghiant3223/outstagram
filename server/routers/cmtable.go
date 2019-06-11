package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"outstagram/server/configs"
	"outstagram/server/middlewares"
)

func CommentableAPIRouter(router *gin.RouterGroup) {
	commentableController, err := configs.InitializeCommentableController()
	if err != nil {
		log.Fatal(err.Error())
	}

	router.GET("/:cmtableID/comments", middlewares.VerifyToken, commentableController.GetComments)

	router.POST("/:cmtableID/comments", middlewares.VerifyToken, commentableController.CreateComment)
}
