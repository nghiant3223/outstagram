package userrepo

import (
	"outstagram/server/models"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func New(dbConnection *gorm.DB) *UserRepository {
	return &UserRepository{db: dbConnection}
}

func (ur *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	if err := ur.db.Find(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
