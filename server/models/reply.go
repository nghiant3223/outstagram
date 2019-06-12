package models

import (
	"github.com/jinzhu/gorm"
	"outstagram/server/dtos/dtomodels"
)

// Reply entity
type Reply struct {
	gorm.Model
	Content     *string `gorm:"not null"`
	Comment Comment
	CommentID   uint
	UserID      uint
	User        User
	ReactableID uint
	Reactors    []string `gorm:"-"`
	ReactCount  int      `gorm:"-"`
}

func (r *Reply) ToDTO() dtomodels.Reply {
	return dtomodels.Reply{
		ID:            r.ID,
		Content:       r.Content,
		CreatedAt:     r.CreatedAt,
		OwnerID:       r.UserID,
		OwnerFullname: r.User.Fullname,
		Reactors:      r.Reactors,
		ReactCount:    r.ReactCount}
}
