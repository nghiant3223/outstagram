package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"outstagram/server/managers"
	"outstagram/server/routers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Cannot load .env file")
	}

	router := gin.Default()

	go managers.HubInstance.Run(managers.StoryManagerInstance.WSMux)
	router.GET("/ws", managers.ServeWs)

	if os.Getenv("APP_ENV") == "production" {
		router.Use(static.Serve("/", static.LocalFile("../client/build", true)))
		router.Use(func(c *gin.Context) {
			c.Next()

			if c.Writer.Status() == http.StatusNotFound {
				c.File("../client/build/index.html")
			}
		})
	}

	router.Use(static.Serve("/images", static.LocalFile("./images", true)))

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

	if os.Getenv("PROTOCOL") == "https" {
		PORT := os.Getenv("PORT")
		if PORT == "" {
			err := router.RunTLS(":3000", "./cert/cert.pem", "./cert/key.pem")
			if err != nil {
				log.Fatal(err.Error())
			}
			return
		}

		if err := router.RunTLS(fmt.Sprintf(":%v", PORT), "./cert/cert.pem", "./cert/key.pem"); err != nil {
			log.Fatal(err.Error())
		}
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
