package models

import (
	"github.com/jinzhu/gorm"
)

// Image entity
type Image struct {
	gorm.Model
	Mini   string
	Tiny   string
	Small  string
	Medium string
	Big    string
	Huge   string
	Origin string
}
