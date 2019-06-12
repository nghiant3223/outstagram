package rctablerepo

import (
	"github.com/jinzhu/gorm"
	"outstagram/server/models"
)

type ReactableRepo struct {
	db *gorm.DB
}

func New(dbConnection *gorm.DB) *ReactableRepo {
	return &ReactableRepo{db: dbConnection}
}

func (r *ReactableRepo) GetReacts(id uint) ([]models.React, error) {
	var reactable models.Reactable
	var reacts []models.React

	if err := r.db.First(&reactable, id).Error; err != nil {
		return nil, err
	}

	r.db.Model(&reactable).Related(&reacts)
	for i := 0; i < len(reacts); i++ {
		r.db.Model(&reacts[i]).Related(&reacts[i].User)
	}

	return reacts, nil
}

func (r *ReactableRepo) GetReactsWithLimit(id uint, limit uint, offset uint) (*models.Reactable, error) {
	var reactable models.Reactable

	if err := r.db.First(&reactable, id).Error; err != nil {
		return nil, err
	}

	if err := r.db.Where("reactable_id = ?", reactable.ID).
		Offset(offset).
		Limit(limit).
		Find(&reactable.Reacts).
		Error; err != nil {
		return nil, err
	}

	reacts := reactable.Reacts
	for i := 0; i < len(reacts); i++ {
		r.db.Model(&reacts[i]).Related(&reacts[i].User)
	}

	return &reactable, nil
}

func (r *ReactableRepo) GetReactCount(reactableID uint) int {
	var count int
	r.db.Model(&models.React{}).Where("reactable_id = ?", reactableID).Count(&count)

	return count
}
