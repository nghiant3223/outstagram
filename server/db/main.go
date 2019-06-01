package db

import (
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
	"outstagram/server/models"
)

var dbInstance *gorm.DB

func New() (*gorm.DB, error) {
	if dbInstance == nil {
		var err error

		// If on TEST mode
		if flag.Lookup("test.v") != nil {
			dbInstance, err = gorm.Open("mysql", "root:root@/outstagram?charset=utf8&parseTime=True&loc=Local")
		} else {
			dbInstance, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@/%v?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_SCHEMA")))
		}

		if err != nil {
			log.Fatal(err.Error())
		}

		dbInstance.LogMode(true)
		dbInstance.Debug()

		dbInstance.SingularTable(true)
		dbInstance.AutoMigrate(&models.User{}, &models.Comment{}, &models.Image{},
			&models.Message{}, &models.NotifBoard{}, &models.Notification{}, &models.PostImage{},
			&models.Post{}, &models.React{}, &models.Reply{}, &models.Room{}, &models.StoryBoard{}, &models.StoryImage{})
	}
	return dbInstance, nil
}
