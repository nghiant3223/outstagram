package userservice

import (
	"outstagram/server/models"
	"outstagram/server/repositories/userrepo"
)

type UserService struct {
	userRepo *userrepo.UserRepository
}

func New(userRepository *userrepo.UserRepository) *UserService {
	return &UserService{userRepo: userRepository}
}

func (us *UserService) FindByUsername(username string) (*models.User, error) {
	return us.userRepo.FindByUsername(username)
}

func (us *UserService) FindByID(id uint) (*models.User, error) {
	return us.userRepo.FindByID(id)
}

func (us *UserService) Save(user *models.User) error {
	if us.CheckExistsByID(user.ID) || us.CheckExistsByUsername(user.Username) {
		if err := us.userRepo.Save(user); err != nil {
			return err
		}
		return nil
	}

	return us.userRepo.Create(user)
}

func (us *UserService) Delete(id uint) error {
	return us.userRepo.DeleteByID(id)
}

func (us *UserService) CheckExistsByID(id uint) bool {
	return us.userRepo.ExistsByID(id)
}

func (us *UserService) CheckExistsByUsername(username string) bool {
	return us.userRepo.ExistsByUsername(username)
}

func (us *UserService) CheckExistsByEmail(email string) bool {
	return us.userRepo.ExistsByEmail(email)
}
