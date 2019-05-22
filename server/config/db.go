package config

import (
	"outstagram/server/entities"

	"github.com/jinzhu/gorm"
)

// ConnectDatabase connects application to the database server
func ConnectDatabase() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:root@/outstagram?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	db.AutoMigrate(&entities.User{})
	return db, err
}
