package db

import (
	"flag"
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
			dbInstance, err = gorm.Open("mysql", os.ExpandEnv("${DB_USERNAME}:${DB_PASSWORD}@/${DB_SCHEMA}?charset=utf8&parseTime=True&loc=Local"))
		}

		if err != nil {
			log.Fatal(err.Error())
		}

		dbInstance.LogMode(true)
		dbInstance.Debug()

		dbInstance.SingularTable(true)
		dbInstance.AutoMigrate(
			&models.Comment{},
			&models.Commentable{},
			&models.Image{},
			&models.Message{},
			&models.Notification{},
			&models.NotifBoard{},
			&models.Post{},
			&models.PostImage{},
			&models.React{},
			&models.Reactable{},
			&models.Reply{},
			&models.Room{},
			&models.StoryBoard{},
			&models.StoryImage{},
			&models.User{},
			&models.Viewable{})
	}

	return dbInstance, nil
}
