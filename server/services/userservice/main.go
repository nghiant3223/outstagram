package userservice

import (
	"errors"
	"github.com/jinzhu/gorm"
	"outstagram/server/constants"
	"outstagram/server/models"
	"outstagram/server/repos/userrepo"
)

type UserService struct {
	userRepo *userrepo.UserRepo
}

func New(userRepo *userrepo.UserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) FindByID(id uint) (*models.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *UserService) FindByUsername(username string) (*models.User, error) {
	return s.userRepo.FindByUsername(username)
}

func (s *UserService) VerifyLogin(username, password string) (*models.User, error) {
	user, err := s.userRepo.FindByUsername(username)
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("username not found")
	}

	if user.Password != password {
		return nil, errors.New("username not found")
	}

	return user, nil
}

func (s *UserService) Save(user *models.User) error {
	if s.CheckExistsByID(user.ID) {
		return s.userRepo.Save(user)
	}

	return s.userRepo.Create(user)
}

func (s *UserService) Delete(id uint) error {
	return s.userRepo.DeleteByID(id)
}

func (s *UserService) CheckExistsByID(id uint) bool {
	return s.userRepo.ExistsByID(id)
}

func (s *UserService) CheckExistsByUsername(username string) bool {
	return s.userRepo.ExistsByUsername(username)
}

func (s *UserService) CheckExistsByEmail(email string) bool {
	return s.userRepo.ExistsByEmail(email)
}

func (s *UserService) GetFollowers(userID uint) []models.User {
	return s.userRepo.GetFollowers(userID)
}

func (s *UserService) GetFollowings(userID uint) []*models.User {
	return s.userRepo.GetFollowings(userID)
}

func (s *UserService) CheckFollow(follow, followed uint) (bool, error) {
	return s.userRepo.CheckFollow(follow, followed)
}

func (s *UserService) Follow(following, follower uint) error {
	hasFollowed, err := s.CheckFollow(following, follower)
	if err != nil {
		return err
	}

	if hasFollowed {
		return errors.New(constants.AlreadyExist)
	}

	return s.userRepo.Follow(following, follower)
}

func (s *UserService) Unfollow(following, follower uint) error {
	hasFollowed, err := s.CheckFollow(following, follower)
	if err != nil {
		return err
	}

	if !hasFollowed {
		return errors.New(constants.NotExist)
	}

	return s.userRepo.Unfollow(following, follower)
}

func (s *UserService) GetPostFeed(userID uint) []uint {
	return s.userRepo.GetPostFeed(userID)
}

func (s *UserService) GetFollowingsWithAffinity(userID uint) []*models.User {
	return s.userRepo.GetFollowingsWithAffinity(userID)
}
