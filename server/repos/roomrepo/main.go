package roomrepo

import (
	"github.com/jinzhu/gorm"
	"outstagram/server/constants"
	"outstagram/server/models"
)

type RoomRepo struct {
	db *gorm.DB
}

func New(dbConnection *gorm.DB) *RoomRepo {
	return &RoomRepo{db: dbConnection}
}

func (r *RoomRepo) Save(room *models.Room) error {
	return r.db.Save(room).Error
}

func (r *RoomRepo) FindByID(id uint) (*models.Room, error) {
	var room models.Room

	err := r.db.First(&room, id).Error
	if err != nil {
		return nil, err
	}

	r.db.Model(&room).Related(&room.Messages)
	if messageCount := len(room.Messages); messageCount > 0 {
		room.LatestMessage = room.Messages[messageCount-1]
		if startIdx := messageCount - constants.MessageCount; startIdx >= 0 {
			room.Messages = room.Messages[startIdx:]
		}
	}

	for _, message := range room.Messages {
		var user models.User
		r.db.Model(message).Related(&user)
		message.User = user
	}

	r.db.Model(&room).Association("Members").Find(&room.Members)
	return &room, nil
}

func (r *RoomRepo) CheckDualRoomExist(userAID, userBID uint) bool {
	var count int
	query := "SELECT 1 FROM joins A INNER JOIN joins B ON A.room_id = B.room_id WHERE A.user_id = ? AND B.user_id = ?"
	r.db.Raw(query, userAID, userBID).Count(&count)
	return count > 0
}

func (r *RoomRepo) FindByPartnerID(userID, partnerID uint) *models.Room {
	if userID == partnerID {
		return nil
	}

	var room models.Room
	query := `
	SELECT room.* 
	FROM joins A INNER JOIN joins B INNER JOIN room 
		ON A.room_id = room.id WHERE 
		A.room_id = B.room_id AND A.user_id <> B.user_id AND A.user_id = ? AND B.user_id = ?`
	r.db.Raw(query, userID, partnerID).Scan(&room)
	r.db.Model(&room).Related(&room.Messages)

	if messageCount := len(room.Messages); messageCount > 0 {
		room.LatestMessage = room.Messages[messageCount-1]
		if startIdx := messageCount - constants.MessageCount; startIdx >= 0 {
			room.Messages = room.Messages[startIdx:]
		}
	}

	for _, message := range room.Messages {
		var user models.User
		r.db.Model(message).Related(&user)
		message.User = user
	}

	r.db.Model(&room).Association("Members").Find(&room.Members)
	return &room
}

func (r *RoomRepo) SaveMessage(message *models.Message) error {
	if err := r.db.Save(message).Error; err != nil {
		return err
	}

	var user models.User
	r.db.Model(message).Related(&user)
	message.User = user
	return nil
}

func (r *RoomRepo) RecentRooms(userID uint) ([]*models.Room, error) {
	var messages []*models.Message
	var rooms []*models.Room

	query := `
	SELECT MessagesInRoomWhichUserJoins.*
	FROM
	(SELECT message.* FROM message INNER JOIN joins ON message.room_id = joins.room_id WHERE joins.user_id = ? GROUP BY id) AS MessagesInRoomWhichUserJoins
	NATURAL JOIN
	(SELECT room_id, max(created_at) as created_at FROM message GROUP BY room_id) AS RoomWithItsLatestMessage
	GROUP BY room_id
	ORDER BY created_at DESC`

	r.db.Raw(query, userID).Scan(&messages)
	for _, message := range messages {
		room, err := r.FindByID(message.RoomID)
		if err != nil {
			return nil, err
		}
		r.db.Model(&room).Association("Members").Find(&room.Members)
		r.db.Model(&message).Related(&message.User)
		room.LatestMessage = message
		rooms = append(rooms, room)
	}
	return rooms, nil
}

func (r *RoomRepo) GetRoomMessages(id uint) (*models.Room, error) {
	var room models.Room

	if err := r.db.First(&room, id).Error; err != nil {
		return nil, err
	}

	r.db.Model(&room).Related(&room.Messages)
	messages := room.Messages
	for i := 0; i < len(messages); i++ {
		r.db.Model(&messages[i]).Related(&messages[i].User)
	}

	return &room, nil
}

func (r *RoomRepo) GetRoomMessagesWithLimit(id, limit, offset uint) (*models.Room, error) {
	var room models.Room

	if err := r.db.First(&room, id).Error; err != nil {
		return nil, err
	}

	query := `
	SELECT * 
		FROM (SELECT * FROM message WHERE room_id = ? ORDER BY created_at DESC LIMIT ? OFFSET ?) AS reversed
	ORDER BY created_at ASC
	`
	if err := r.db.Raw(query, id, limit, offset).Scan(&room.Messages).Error;
		err != nil {
		return nil, err
	}

	messages := room.Messages
	for i := 0; i < len(messages); i++ {
		r.db.Model(&messages[i]).Related(&messages[i].User)
	}

	return &room, nil
}
