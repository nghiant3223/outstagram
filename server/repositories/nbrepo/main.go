package nbrepo

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

func (br *NotifBoardRepo) Save(board *models.NotifBoard) error {
	return br.db.Create(board).Error
}
