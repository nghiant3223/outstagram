package vwablerepo

import (
	"github.com/jinzhu/gorm"
	"outstagram/server/models"
)

type ViewableRepo struct {
	db *gorm.DB
}

func New(dbConnection *gorm.DB) *ViewableRepo {
	return &ViewableRepo{db: dbConnection}
}

func (r *ViewableRepo) IncrementView(userID, viewableID uint) error {
	if err := r.db.Raw("INSERT INTO views(user_id, viewable_id) VALUES (?, ?)", userID, viewableID).Error; err != nil {
		return err
	}

	if err := r.db.Model(models.Post{}).Update("numViewed", gorm.Expr("numViewed + ?", 1)).Error; err != nil {
		return err
	}

	return nil
}
