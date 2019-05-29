package sbrepo

import (
	"github.com/jinzhu/gorm"
	"outstagram/server/models"
)

type StoryBoardRepo struct {
	db *gorm.DB
}

func New(dbConnection *gorm.DB) *StoryBoardRepo {
	return &StoryBoardRepo{db: dbConnection}
}

func (br *StoryBoardRepo) Save(board *models.StoryBoard) error {
	return br.db.Create(board).Error
}
