package models

import (
	"github.com/jinzhu/gorm"
	"outstagram/server/dtos/dtomodels"
)

// Comment entity
type Comment struct {
	gorm.Model
	Content       *string `gorm:"not null"`
	CommentableID uint
	ReactableID   uint
	UserID        uint
	User          User
	Replies       []Reply
	ReplyCount    int      `gorm:"-"`
	Reactors      []string `gorm:"-"`
	ReactCount    int      `gorm:"-"`
}

func (c *Comment) ToDTO() dtomodels.Comment {
	return dtomodels.Comment{
		ID:            c.ID,
		Content:       c.Content,
		ReplyCount:    c.ReplyCount,
		CreatedAt:     c.CreatedAt,
		OwnerFullname: c.User.Fullname,
		OwnerID:       c.UserID,
		ReactCount:    c.ReactCount,
		Reactors:      c.Reactors,
		CommentableID: c.CommentableID,
		ReactableID:   c.ReactableID}
}
