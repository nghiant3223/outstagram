package models

import (
	"github.com/jinzhu/gorm"
)

// Image entity
type Image struct {
	gorm.Model
	TinyURL     string // 32x32 - for avatar
	SmallURL    string // 100 - for small image in `Discover`
	MediumURL   string // 200 - for `My Images` (grid of images)
	BigURL      string // 300 - for `My post` (list of posts), large image in `Discover`
	HugeURL     string // 500 - for `Postfeed`
	OriginalURL string
}
