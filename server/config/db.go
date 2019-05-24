package config

import (
	"fmt"
	"os"

	"outstagram/server/models"

	"github.com/jinzhu/gorm"
)

// ConnectDatabase connects application to the database server
func ConnectDatabase() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%v:%v@/%v?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_SCHEMA")))
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	db.AutoMigrate(&models.User{}, &models.Comment{}, &models.Image{}, 
		&models.Message{}, &models.NotifBoard{}, &models.Notification{}, &models.PostImage{}, 
		&models.Post{}, &models.React{}, &models.Reply{}, &models.Room{}, &models.StoryBoard{}, &models.StoryImage{})
	return db, err
}
