package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"outstagram/server/injection"
	"outstagram/server/middlewares"
)

func AuthAPIRouter(router *gin.Engine, routerGroup *gin.RouterGroup) {
	authController, err := injection.InitializeAuthController()
	if err != nil {
		log.Fatal(err.Error())
	}

	routerGroup.POST("/register", authController.Register)

	routerGroup.POST("/login", authController.Login)

	routerGroup.POST("/logout", middlewares.VerifyToken(true), authController.Logout)
}
