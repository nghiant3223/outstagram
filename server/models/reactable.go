package models

import "github.com/jinzhu/gorm"

type Reactable struct {
	gorm.Model
	Reacts     []React
	ReactCount uint     `gorm:"-"`
	Reactors   []string `gorm:"-"`
}