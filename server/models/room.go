package models

import (
	"github.com/jinzhu/gorm"
	"outstagram/server/dtos/dtomodels"
)

// Room entity
type Room struct {
	gorm.Model
	Name          string
	Type          bool    `gorm:"default:false"`
	Members       []*User `gorm:"many2many:joins"`
	Messages      []*Message
	LatestMessage *Message `gorm:"-"`
	ImageID       uint
}

func (r *Room) ToDTO(userID uint) dtomodels.Room {
	var dtoUsers []*dtomodels.User
	for _, user := range r.Members {
		dtoUser := user.ToUserDTO()
		dtoUsers = append(dtoUsers, &dtoUser)
	}

	var dtoMessages []*dtomodels.Message
	for _, message := range r.Messages {
		dtoMessage := message.ToDTO()
		dtoMessages = append(dtoMessages, &dtoMessage)
	}

	dtoRoom := dtomodels.Room{
		ID:       r.ID,
		Type:     r.Type,
		Name:     r.Name,
		Members:  dtoUsers,
		Messages: dtoMessages,
	}

	if r.LatestMessage != nil {
		dtoMessage := r.LatestMessage.ToDTO()
		dtoRoom.LatestMessage = &dtoMessage
	}

	if !r.Type {
		if r.Members[0].ID != userID {
			partner := r.Members[0].ToUserDTO()
			dtoRoom.Partner = &partner
		} else {
			partner := r.Members[1].ToUserDTO()
			dtoRoom.Partner = &partner
		}
		dtoRoom.ImageID = 0
	} else {
		dtoRoom.ImageID = r.ImageID
		dtoRoom.Partner = nil
	}

	return dtoRoom
}
