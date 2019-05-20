package main

import (
	"outstagram/server/managers"
	"outstagram/server/routers"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := gin.Default()

	go managers.HubInstance.Run(managers.StoryManagerInstance.WSMux)
	router.GET("/ws", func(c *gin.Context) { managers.ServeWs(c.Writer, c.Request) })

	apiRouter := router.Group("/api")
	{
		routers.UserAPIRouter(apiRouter.Group("/user"))
		routers.StoryAPIRouter(apiRouter.Group("/story"))
	}

	router.Run(":3000")
}
