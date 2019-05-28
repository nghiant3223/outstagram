package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"outstagram/server/models"
)

var dbInstance *gorm.DB

func New() (*gorm.DB, error) {
	if dbInstance == nil {
		dbInstance, err := gorm.Open("mysql", fmt.Sprintf("%v:%v@/%v?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_SCHEMA")))
		if err != nil {
			log.Fatal(err.Error())
		}
		dbInstance.SingularTable(true)
		dbInstance.AutoMigrate(&models.User{}, &models.Comment{}, &models.Image{},
			&models.Message{}, &models.NotifBoard{}, &models.Notification{}, &models.PostImage{},
			&models.Post{}, &models.React{}, &models.Reply{}, &models.Room{}, &models.StoryBoard{}, &models.StoryImage{})

	}
	return dbInstance, nil
}