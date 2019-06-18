package db

import (
	"flag"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"log"
	"outstagram/server/models"
)

var dbConn *gorm.DB

func New() (*gorm.DB, error) {
	if dbConn == nil {
		var err error

		// If on TEST mode
		if flag.Lookup("test.v") != nil {
			dbConn, err = gorm.Open("mysql", "root:root@/outstagram?charset=utf8&parseTime=True&loc=Local")
		} else {
			dbConn, err = gorm.Open(viper.GetString("db.dialect"), viper.GetString("db.url"))
		}

		if err != nil {
			log.Fatal(err.Error())
		}

		dbConn.SingularTable(true)

		if viper.GetString("env") != "production" {
			dbConn.LogMode(true)
			dbConn.Debug()

			dbConn.AutoMigrate(
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
	}

	return dbConn, nil
}
