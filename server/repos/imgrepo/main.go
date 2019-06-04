package imgrepo

import (
	"github.com/jinzhu/gorm"
	"outstagram/server/models"
)

type ImageRepo struct {
	db *gorm.DB
}

func New(dbConnection *gorm.DB) *ImageRepo {
	return &ImageRepo{db: dbConnection}
}

func (r *ImageRepo) Save(image *models.Image) error {
	return r.db.Save(image).Error
}
