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
		routers.MeAPIRouter(apiRouter.Group("/me"))
		routers.AuthAPIRouter(apiRouter.Group("/auth"))
		routers.UserAPIRouter(apiRouter.Group("/users"))
		routers.PostAPIRouter(apiRouter.Group("/posts"))
		routers.StoryAPIRouter(apiRouter.Group("/stories"))
		routers.FollowAPIRouter(apiRouter.Group("/follows"))
		routers.ReactAPIRouter(apiRouter.Group("/reactions"))
		routers.CommentableAPIRouter(apiRouter.Group("/commentable"))
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
