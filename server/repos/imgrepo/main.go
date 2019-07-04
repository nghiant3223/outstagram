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

func (r *ImageRepo) FindByID(id uint) (*models.Image, error) {
	var image models.Image
	if err := r.db.First(&image, id).Error; err != nil {
		return nil, err
	}
	return &image, nil
}
