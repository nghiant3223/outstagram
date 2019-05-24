package main

import (
	"log"
	"os"
	"fmt"
	
	"outstagram/server/managers"
	"outstagram/server/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Cannot load .env file")
	}

	router := gin.Default()

	go managers.HubInstance.Run(managers.StoryManagerInstance.WSMux)
	router.GET("/ws", func(c *gin.Context) { managers.ServeWs(c.Writer, c.Request) })

	apiRouter := router.Group("/api")
	{
		routers.UserAPIRouter(apiRouter.Group("/user"))
		routers.StoryAPIRouter(apiRouter.Group("/story"))
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		router.Run(":3000")
		return
	}
	router.Run(fmt.Sprintf(":%v", PORT))
}