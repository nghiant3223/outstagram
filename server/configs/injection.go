//+build wireinject

package configs

import (
	"outstagram/server/controllers/authcontroller"
	"outstagram/server/controllers/postcontroller"
	"outstagram/server/controllers/usercontroller"
	"outstagram/server/db"
	"outstagram/server/repos/imgrepo"
	"outstagram/server/repos/notifbrepo"
	"outstagram/server/repos/postimgrepo"
	"outstagram/server/repos/postrepo"
	"outstagram/server/repos/storybrepo"
	"outstagram/server/repos/userrepo"
	"outstagram/server/services/imgservice"
	"outstagram/server/services/notifbservice"
	"outstagram/server/services/postimgservice"
	"outstagram/server/services/postservice"
	"outstagram/server/services/storybservice"
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

		notifbservice.New,
		notifbrepo.New,

		storybservice.New,
		storybrepo.New,

		db.New)
	return &authcontroller.Controller{}, nil
}

func InitializePostController() (*postcontroller.Controller, error) {
	wire.Build(
		postcontroller.New,

		postservice.New,
		postrepo.New,

		postimgservice.New,
		postimgrepo.New,

		imgservice.New,
		imgrepo.New,

		db.New)
	return &postcontroller.Controller{}, nil
}
