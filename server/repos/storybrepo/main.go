package storybrepo

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

func (r *StoryBoardRepo) Save(board *models.StoryBoard) error {
	return r.db.Create(board).Error
}
