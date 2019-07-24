package managers

import (
	"outstagram/server/db"
	"outstagram/server/dtos/dtomodels"
	"outstagram/server/repos/userrepo"
	"outstagram/server/services/userservice"
)

type APIProvider interface {
	GetUserFollowers(uint) []*dtomodels.SimpleUser
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
