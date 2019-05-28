package userservice

import (
	"outstagram/server/models"
	userrepo2 "outstagram/server/repositories/userrepo"
)

type UserService struct {
	repo *userrepo2.UserRepository
}

func New(userRepository *userrepo2.UserRepository) *UserService {
	return &UserService{repo: userRepository}
}

func (us *UserService) FindAll() ([]models.User, error) {
	return us.repo.FindAll()
}

func (us *UserService) FindByUsername(username string) (*models.User, error) {
	return us.repo.FindByUsername(username)
}
