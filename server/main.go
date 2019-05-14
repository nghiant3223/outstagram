package main

import (
	"net/http"
	usercontroller "outstagram/server/controllers/userController"
	"outstagram/server/entities"
	"outstagram/server/repositories"
	"outstagram/server/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@/test?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.SingularTable(true)
	db.AutoMigrate(&entities.User{})

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := usercontroller.NewUserController(userService)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello")
	})

	router.GET("/:username", userController.GetUserByUsername)

	router.Run("localhost:5000")
}
