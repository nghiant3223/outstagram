package db

import (
	"flag"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"log"
	"outstagram/server/models"
	"outstagram/server/pkg/configutils"
)

var dbConn *gorm.DB

func New() (*gorm.DB, error) {
	if dbConn == nil {
		configutils.LoadConfiguration("outstagram", "main", "configs")

		var err error

		// If on TEST mode
		if flag.Lookup("test.v") != nil {
			dbConn, err = gorm.Open("mysql", "root:root@tcp(172.24.21.56:33060)/outstagram?charset=utf8mb4&parseTime=True&loc=Local")
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
				&models.Story{},
				&models.User{},
				&models.Viewable{})
		}
	}

	return dbConn, nil
}
