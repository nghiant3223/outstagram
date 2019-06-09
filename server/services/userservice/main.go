package userservice

import (
	"outstagram/server/models"
	"outstagram/server/repos/userrepo"
)

type UserService struct {
	userRepo *userrepo.UserRepo
}

func New(userRepo *userrepo.UserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) FindByUsername(username string) (*models.User, error) {
	return s.userRepo.FindByUsername(username)
}

func (s *UserService) FindByID(id uint) (*models.User, error) {
	return s.userRepo.FindByID(id)
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

func (s *UserService) GetFollowings(userID uint) []models.User {
	return s.userRepo.GetFollowings(userID)
}
