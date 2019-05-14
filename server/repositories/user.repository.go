package repositories

import (
	"outstagram/server/entities"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(dbConnection *gorm.DB) *UserRepository {
	return &UserRepository{db: dbConnection}
}

func (ur *UserRepository) FindAll() ([]entities.User, error) {
	var users []entities.User
	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) FindByUsername(username string) (*entities.User, error) {
	var user entities.User
	if err := ur.db.Find(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
