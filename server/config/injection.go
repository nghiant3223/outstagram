//+build wireinject

package config

import (
	"outstagram/server/controllers/usercontroller"
	userrepo "outstagram/server/repositories"
	userservice "outstagram/server/services/usersservice"

	"github.com/google/wire"
)

func InitializeUserController() (*usercontroller.Controller, error) {
	wire.Build(usercontroller.New, userservice.New, userrepo.New, ConnectDatabase)
	return &usercontroller.Controller{}, nil
}
