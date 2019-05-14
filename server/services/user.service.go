package services

import (
	"outstagram/server/entities"
	"outstagram/server/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{repo: userRepository}
}

func (us *UserService) FindAll() ([]entities.User, error) {
	return us.repo.FindAll()
}

func (us *UserService) FindByUsername(username string) (*entities.User, error) {
	return us.repo.FindByUsername(username)
}
