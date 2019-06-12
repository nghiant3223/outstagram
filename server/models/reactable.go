package models

import "github.com/jinzhu/gorm"

type Reactable struct {
	gorm.Model
	Reacts     []React
	ReactCount uint     `gorm:"-"`
	Reactors   []string `gorm:"-"`

	// These 4 fields are used for GetVisibility in ReactableRepo.GetVisibility
	Post       Post
	PostImage  PostImage
	Comment    Comment
	Reply      Reply
}
