package rctrepo

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"outstagram/server/models"
	"strings"
)

type ReactRepo struct {
	db *gorm.DB
}

func New(dbConnection *gorm.DB) *ReactRepo {
	return &ReactRepo{db: dbConnection}
}

func (r *ReactRepo) Save(react *models.React) error {
	return r.db.Save(&react).Error
}

func (r *ReactRepo) Find(projector map[string]interface{}) ([]models.React, error) {
	var fields []string
	var values []interface{}
	var reacts []models.React

	for key, v := range projector {
		fields = append(fields, fmt.Sprintf("%v = ?", key))
		values = append(values, v)
	}

	projections := strings.Join(fields, " AND ")
	if err := r.db.Where(projections, values...).Find(&reacts).Error; err != nil {
		return nil, err
	}

	return reacts, nil
}

func (r *ReactRepo) FindByID(id uint) (*models.React, error) {
	var react models.React

	if err := r.db.First(&react, id).Error; err != nil {
		return nil, err
	}

	return &react, nil
}

func (r *ReactRepo) Delete(react *models.React) error {
	return r.db.Delete(react).Error
}
