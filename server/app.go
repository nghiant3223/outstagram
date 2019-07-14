package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
	"outstagram/server/managers"
	"outstagram/server/routers"
)

func main() {
	router := gin.Default()

	go managers.Hub.Run(managers.StoryManager.WSMux)
	router.GET("/ws", managers.ServeWs)

	if os.Getenv("ENV") == "production" {
		router.Use(static.Serve("/", static.LocalFile("./client-build", true)))

		router.NoMethod(func(c *gin.Context) {
			c.File("./client-build/index.html")
		})

		router.NoRoute(func(c *gin.Context) {
			c.File("./client-build/index.html")
		})
	}

	router.Use(static.Serve("/images", static.LocalFile("./images", true)))

	apiRouter := router.Group("/api")
	{
		routers.MeAPIRouter(router, apiRouter.Group("/me"))
		routers.AuthAPIRouter(router, apiRouter.Group("/auth"))
		routers.UserAPIRouter(router, apiRouter.Group("/users"))
		routers.PostAPIRouter(router, apiRouter.Group("/posts"))
		routers.StoryAPIRouter(router, apiRouter.Group("/stories"))
		routers.FollowAPIRouter(router, apiRouter.Group("/follows"))
		routers.ReactAPIRouter(router, apiRouter.Group("/reactables"))
		routers.CommentableAPIRouter(router, apiRouter.Group("/commentables"))
	}

	staticRouter := router.Group("/static")
	{
		routers.ImageStaticRouter(router, staticRouter.Group("/images"))
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		err := router.Run(":3000")
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	if err := router.Run(fmt.Sprintf(":%v", PORT)); err != nil {
		log.Fatal(err.Error())
	}
}
