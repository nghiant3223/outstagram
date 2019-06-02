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
	for _, user := range users {
		ur.db.Model(&user).Related(&user.NotifBoard, "NotifBoard")
		ur.db.Model(&user).Related(&user.StoryBoard, "StoryBoard")
	}
	return users, nil
}

func (ur *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	if err := ur.db.Find(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}
	ur.db.Model(&user).Related(&user.NotifBoard, "NotifBoard")
	ur.db.Model(&user).Related(&user.StoryBoard, "StoryBoard")
	return &user, nil
}

func (ur *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	if err := ur.db.Find(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	ur.db.Model(&user).Related(&user.NotifBoard, "NotifBoard")
	ur.db.Model(&user).Related(&user.StoryBoard, "StoryBoard")
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

func (ur *UserRepository) Create(user *models.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}

	ur.db.Create(&models.NotifBoard{UserID:user.ID})
	ur.db.Create(&models.StoryBoard{UserID:user.ID})
	return nil
}

func (ur *UserRepository) CreateAll(users []*models.User) error {
	for _, user := range users {
		err := ur.db.Create(&user).Error
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

func (ur *UserRepository) ExistsByEmail(email string) bool {
	var user models.User
	return !ur.db.Where("email = ?", email).First(&user).RecordNotFound()
}

func (ur *UserRepository) GetFollowers(userID uint) []models.User {
	var users []models.User
	ur.db.Raw(`SELECT user.* FROM user INNER JOIN follows ON user_follow_id = user.id WHERE follows.user_followed_id = ?`, userID).Scan(&users)
	return users
}

func (ur *UserRepository) GetFollowings(userID uint) []models.User {
	var users []models.User
	ur.db.Raw(`SELECT user.* FROM user INNER JOIN follows ON user_followed_id = user.id WHERE follows.user_follow_id = ?`, userID).Scan(&users)
	return users
}