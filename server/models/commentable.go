package models

import "github.com/jinzhu/gorm"

type Commentable struct {
	gorm.Model
	Comments     []Comment
	CommentCount int `gorm:"-"`

	// These two fields are used for GetVisibility in CommentableRepo.GetVisibility
	Post         Post
	PostImage    PostImage
}
