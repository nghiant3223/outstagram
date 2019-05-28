//+build wireinject

package config

import (
	"outstagram/server/controllers/authcontroller"
	"outstagram/server/controllers/usercontroller"
	"outstagram/server/db"
	"outstagram/server/repositories/userrepo"
	"outstagram/server/services/userservice"

	"github.com/google/wire"
)

func InitializeUserController() (*usercontroller.Controller, error) {
	wire.Build(usercontroller.New, userservice.New, userrepo.New, db.New)
	return &usercontroller.Controller{}, nil
}

func InitializeAuthController() (*authcontroller.Controller, error) {
	wire.Build(authcontroller.New, userservice.New, userrepo.New, db.New)
	return &authcontroller.Controller{}, nil
}