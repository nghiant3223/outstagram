package userservice

import (
	"outstagram/server/models"
	userrepo "outstagram/server/repositories"
)

type UserService struct {
	repo *userrepo.UserRepository
}

func New(userRepository *userrepo.UserRepository) *UserService {
	return &UserService{repo: userRepository}
}

func (us *UserService) FindAll() ([]models.User, error) {
	return us.repo.FindAll()
}

func (us *UserService) FindByUsername(username string) (*models.User, error) {
	return us.repo.FindByUsername(username)
}
