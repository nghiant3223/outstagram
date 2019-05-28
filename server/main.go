package main

import (
	"fmt"
	"log"
	"os"

	"outstagram/server/managers"
	"outstagram/server/routers"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Cannot load .env file")
	}

	router := gin.Default()

	go managers.HubInstance.Run(managers.StoryManagerInstance.WSMux)
	router.GET("/ws", managers.ServeWs)

	apiRouter := router.Group("/api")
	{
		routers.UserAPIRouter(apiRouter.Group("/user"))
		routers.StoryAPIRouter(apiRouter.Group("/story"))
		routers.AuthAPIRouter(apiRouter.Group("/auth"))
	}	

	PORT := os.Getenv("PORT")
	if PORT == "" {
		err := router.Run(":3000")
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	err := router.Run(fmt.Sprintf(":%v", PORT))
	if err != nil {
		log.Fatal(err.Error())
	}
}