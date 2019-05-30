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

func (ur *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	if err := ur.db.Find(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) Save(user *models.User) error {
	return ur.db.Save(user).Error
}

func (ur *UserRepository) SaveAll(users []*models.User) error {
	for _, user := range users {
		err := ur.db.Save(&user).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (ur *UserRepository) DeleteByID(id uint) error {
	var user models.User
	if err := ur.db.Where("id = ?", id).Find(&user).Error; err != nil {
		return err
	}
	ur.db.Delete(&user)
	return nil
}

func (ur *UserRepository) DeleteAll(ids []uint) error {
	for _, id := range ids {
		var user models.User
		if err := ur.db.Where("id = ?", id).Find(&user).Error; err != nil {
			return err
		}
		ur.db.Delete(&user)
	}
	return nil
}

func (ur *UserRepository) ExistsByID(id uint) bool {
	var user models.User
	return !ur.db.First(&user, id).RecordNotFound()
}

func (ur *UserRepository) ExistsByUsername(username string) bool {
	var user models.User
	return !ur.db.Where("username = ?", username).First(&user).RecordNotFound()
}

func (ur *UserRepository) ExistsByEmail(username string) bool {
	var user models.User
	return !ur.db.Where("email = ?", username).First(&user).RecordNotFound()
}
