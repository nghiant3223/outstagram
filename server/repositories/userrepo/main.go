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

func (ur *UserRepository) FindById(id uint) (*models.User, error) {
	var user models.User
	if err := ur.db.Find(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) Save(user *models.User) {
	ur.db.Create(user)
}

func (ur *UserRepository) SaveAll(users []*models.User) {
	for _, user := range users {
		ur.db.Create(user)
	}
}

func (ur *UserRepository) DeleteByID(id uint) {
	ur.db.Where("id = ?", id).Delete(&models.User{})
}

func (ur *UserRepository) DeleteAll(ids []uint) {
	for _, id := range ids {
		ur.db.Where("id = ?", id).Delete(&models.User{})
	}
}

func (ur *UserRepository) ExistsById(id uint) bool {
	var user models.User
	return !ur.db.First(&user, id).RecordNotFound()
}
