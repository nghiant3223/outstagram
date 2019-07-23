package models

import (
	"github.com/jinzhu/gorm"
	"outstagram/server/dtos/dtomodels"
)

// Message entity
type Message struct {
	gorm.Model
	Content string `gorm:"not null"`
	Type    int8   `gorm:"default:0"`
	RoomID  uint
	UserID  uint
	User    User
}

func (m *Message) ToDTO() dtomodels.Message {
	return dtomodels.Message{
		ID:        m.ID,
		Type:      m.Type,
		RoomID:    m.RoomID,
		CreatedAt: m.CreatedAt,
		Content:   m.Content,
		AuthorID:  m.User.ID,
	}
}
