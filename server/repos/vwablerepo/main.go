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
	var count int
	r.db.Raw("SELECT 1 FROM views WHERE user_id = ? AND viewable_id = ?", userID, viewableID).Count(&count)

	if count < 1 {
		if err := r.db.Exec("INSERT INTO views(user_id, viewable_id) VALUES (?, ?)", userID, viewableID).Error; err != nil {
			return err
		}
	}

	if err := r.db.Model(models.Post{}).Where("viewable_id = ?", viewableID).Update("popularity", gorm.Expr("popularity + ?", 1)).Error; err != nil {
		return err
	}

	return nil
}
