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

	apiRouter := router.Group("/api")
	{
		routers.MeAPIRouter(apiRouter.Group("/me"))
		routers.AuthAPIRouter(apiRouter.Group("/auth"))
		routers.UserAPIRouter(apiRouter.Group("/users"))
		routers.PostAPIRouter(apiRouter.Group("/posts"))
		routers.StoryAPIRouter(apiRouter.Group("/stories"))
		routers.FollowAPIRouter(apiRouter.Group("/follows"))
		routers.ReactAPIRouter(apiRouter.Group("/reactions"))
		routers.CommentableAPIRouter(apiRouter.Group("/commentable"))
	}

	staticRouter := router.Group("/static")
	{
		routers.ImageStaticRouter(staticRouter.Group("/images"))
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
