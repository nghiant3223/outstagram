package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"outstagram/server/injection"
)

func ImageStaticRouter(router *gin.RouterGroup) {
	imageController, err := injection.InitializeImageController()
	if err != nil {
		log.Fatal(err.Error())
	}

	router.GET("/avatars/:userID", imageController.GetUserAvatar)
	router.GET("/others/:imageID", imageController.GetImage)
}
