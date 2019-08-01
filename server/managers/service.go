package managers

import (
	"outstagram/server/db"
	"outstagram/server/dtos/dtomodels"
	"outstagram/server/models"
	"outstagram/server/repos/userrepo"
	"outstagram/server/services/userservice"
)

type APIProvider interface {
	GetUserFollowers(userID uint) []*dtomodels.SimpleUser
	GetUserRoomIDs(userID uint) []uint
	UpdateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
}

type LocalAPIProvider struct {
	userService *userservice.UserService
}

func NewLocalAPIProvider() *LocalAPIProvider {
	dbConn, _ := db.New()
	userRepo := userrepo.New(dbConn)
	userService := userservice.New(userRepo)
	return &LocalAPIProvider{userService: userService}
}

func (l *LocalAPIProvider) GetUserFollowers(userID uint) []*dtomodels.SimpleUser {
	var dtoFollowers []*dtomodels.SimpleUser

	followers := l.userService.GetFollowers(userID)
	for _, follower := range followers {
		simpleDTO := follower.ToSimpleDTO()
		dtoFollowers = append(dtoFollowers, &simpleDTO)
	}

	return dtoFollowers
}

func (l *LocalAPIProvider) GetUserRoomIDs(userID uint) []uint {
	roomIDs, err := l.userService.GetUserRoomIDs(userID)
	if err != nil {
		return nil
	}
	return roomIDs
}

func (l *LocalAPIProvider) UpdateUser(user *models.User) error {
	return l.userService.Save(user)
}

func (l *LocalAPIProvider) GetUserByID(userID uint) (*models.User, error) {
	user, err := l.userService.FindByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
