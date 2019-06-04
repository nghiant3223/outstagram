package userrepo

import (
	"outstagram/server/models"

	"github.com/jinzhu/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func New(dbConnection *gorm.DB) *UserRepo {
	return &UserRepo{db: dbConnection}
}

func (r *UserRepo) FindAll() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	for _, user := range users {
		r.db.Model(&user).Related(&user.NotifBoard, "NotifBoard")
		r.db.Model(&user).Related(&user.StoryBoard, "StoryBoard")
	}

	return users, nil
}

func (r *UserRepo) FindByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.db.Find(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}

	r.db.Model(&user).Related(&user.NotifBoard, "NotifBoard")
	r.db.Model(&user).Related(&user.StoryBoard, "StoryBoard")
	return &user, nil
}

func (r *UserRepo) FindByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.Find(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	r.db.Model(&user).Related(&user.NotifBoard, "NotifBoard")
	r.db.Model(&user).Related(&user.StoryBoard, "StoryBoard")
	return &user, nil
}

func (r *UserRepo) Save(user *models.User) error {
	err := r.db.Save(user).Error
	if err != nil {
		return err
	}

	r.db.Create(&models.NotifBoard{UserID: user.ID})
	r.db.Create(&models.StoryBoard{UserID: user.ID})
	return nil
}

func (r *UserRepo) SaveAll(users []*models.User) error {
	for _, user := range users {
		err := r.db.Save(&user).Error
		if err != nil {
			return err
		}

		r.db.Create(&models.NotifBoard{UserID: user.ID})
		r.db.Create(&models.StoryBoard{UserID: user.ID})
	}

	return nil
}

func (r *UserRepo) DeleteByID(id uint) error {
	var user models.User
	if err := r.db.Where("id = ?", id).Find(&user).Error; err != nil {
		return err
	}

	r.db.Delete(&user)
	return nil
}

func (r *UserRepo) DeleteAll(ids []uint) error {
	for _, id := range ids {
		var user models.User
		if err := r.db.Where("id = ?", id).Find(&user).Error; err != nil {
			return err
		}
		r.db.Delete(&user)
	}

	return nil
}

func (r *UserRepo) ExistsByID(id uint) bool {
	var user models.User
	return !r.db.First(&user, id).RecordNotFound()
}

func (r *UserRepo) ExistsByUsername(username string) bool {
	var user models.User
	return !r.db.Where("username = ?", username).First(&user).RecordNotFound()
}

func (r *UserRepo) ExistsByEmail(email string) bool {
	var user models.User
	return !r.db.Where("email = ?", email).First(&user).RecordNotFound()
}

func (r *UserRepo) GetFollowers(userID uint) []models.User {
	var users []models.User
	r.db.Raw(`SELECT user.* FROM user INNER JOIN follows ON user_follow_id = user.id WHERE follows.user_followed_id = ?`, userID).Scan(&users)
	return users
}

func (r *UserRepo) GetFollowings(userID uint) []models.User {
	var users []models.User
	r.db.Raw(`SELECT user.* FROM user INNER JOIN follows ON user_followed_id = user.id WHERE follows.user_follow_id = ?`, userID).Scan(&users)
	return users
}
