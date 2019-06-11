package models

import "github.com/jinzhu/gorm"

type Commentable struct {
	gorm.Model
	Comments     []Comment
	CommentCount int `gorm:"-"`
	Post         Post
	PostImage    PostImage
}
