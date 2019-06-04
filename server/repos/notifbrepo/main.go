package notifbrepo

import (
	"github.com/jinzhu/gorm"
	"outstagram/server/models"
)

type NotifBoardRepo struct {
	db *gorm.DB
}

func New(dbConnection *gorm.DB) *NotifBoardRepo {
	return &NotifBoardRepo{db: dbConnection}
}

func (r *NotifBoardRepo) Save(board *models.NotifBoard) error {
	return r.db.Create(board).Error
}
