package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"outstagram/server/injection"
)

func ImageStaticRouter(router *gin.Engine, routerGroup *gin.RouterGroup) {
	imageController, err := injection.InitializeImageController()
	if err != nil {
		log.Fatal(err.Error())
	}

	routerGroup.GET("/avatars/:userID", imageController.GetUserAvatar)
	routerGroup.GET("/others/:imageID", imageController.GetImage)
}
