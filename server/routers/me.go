package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"outstagram/server/injection"
	"outstagram/server/middlewares"
)

func MeAPIRouter(router *gin.Engine, routerGroup *gin.RouterGroup) {
	meController, err := injection.InitializeMeController()
	if err != nil {
		log.Fatal(err.Error())
	}

	routerGroup.Use(middlewares.VerifyToken(true))

	routerGroup.GET("", meController.GetMe)
	routerGroup.PATCH("", meController.UpdateUser)

	routerGroup.GET("/newsfeed", meController.GetNewsFeed)
	routerGroup.GET("/storyfeed", meController.GetStoryFeed)
	routerGroup.GET("/follow-suggestions", meController.GetFollowSuggestion)
	routerGroup.GET("/posts", middlewares.RedirectToDuplicateRoute(router, "/api/users/%v/posts"))
}
