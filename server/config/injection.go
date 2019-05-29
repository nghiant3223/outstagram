//+build wireinject

package config

import (
	"outstagram/server/controllers/authcontroller"
	"outstagram/server/controllers/usercontroller"
	"outstagram/server/db"
	"outstagram/server/repositories/nbrepo"
	"outstagram/server/repositories/sbrepo"
	"outstagram/server/repositories/userrepo"
	"outstagram/server/services/nbservice"
	"outstagram/server/services/sbservice"
	"outstagram/server/services/userservice"

	"github.com/google/wire"
)

func InitializeUserController() (*usercontroller.Controller, error) {
	wire.Build(
		usercontroller.New,

		userservice.New,
		userrepo.New,

		db.New)
	return &usercontroller.Controller{}, nil
}

func InitializeAuthController() (*authcontroller.Controller, error) {
	wire.Build(
		authcontroller.New,

		userservice.New,
		userrepo.New,

		nbservice.New,
		nbrepo.New,

		sbservice.New,
		sbrepo.New,

		db.New)
	return &authcontroller.Controller{}, nil
}
