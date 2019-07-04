package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"outstagram/server/injection"
	"outstagram/server/middlewares"
)

func ImageAPIRouter(router *gin.RouterGroup) {
	imageController, err := injection.InitializeImageController()
	if err != nil {
		log.Fatal(err.Error())
	}

	router.Use(middlewares.VerifyToken)

	router.GET("/avatar/:userID", imageController.GetImage)
	router.GET("/:imageID", imageController.GetImage)
}

