//+build wireinject

package config

import (
	"outstagram/server/controllers/authcontroller"
	"outstagram/server/controllers/usercontroller"
	"outstagram/server/repositories/userrepo"
	"outstagram/server/services/userservice"

	"github.com/google/wire"
)

func InitializeUserController() (*usercontroller.Controller, error) {
	wire.Build(usercontroller.New, userservice.New, userrepo.New, ConnectDatabase)
	return &usercontroller.Controller{}, nil
}

func InitializeAuthController() (*authcontroller.Controller, error) {
	wire.Build(authcontroller.New, userservice.New, userrepo.New, ConnectDatabase)
	return &authcontroller.Controller{}, nil
}