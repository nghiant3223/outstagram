//+build wireinject

package injection

import (
	"outstagram/server/controllers/authcontroller"
	"outstagram/server/controllers/cmtablecontroller"
	"outstagram/server/controllers/flcontroller"
	"outstagram/server/controllers/mecontroller"
	"outstagram/server/controllers/postcontroller"
	"outstagram/server/controllers/rctcontroller"
	"outstagram/server/controllers/storycontroller"
	"outstagram/server/controllers/usercontroller"
	"outstagram/server/db"
	"outstagram/server/repos/cmtablerepo"
	"outstagram/server/repos/cmtrepo"
	"outstagram/server/repos/imgrepo"
	"outstagram/server/repos/notifbrepo"
	"outstagram/server/repos/postimgrepo"
	"outstagram/server/repos/postrepo"
	"outstagram/server/repos/rctablerepo"
	"outstagram/server/repos/rctrepo"
	"outstagram/server/repos/replyrepo"
	"outstagram/server/repos/storybrepo"
	"outstagram/server/repos/userrepo"
	"outstagram/server/repos/vwablerepo"
	"outstagram/server/services/cmtableservice"
	"outstagram/server/services/cmtservice"
	"outstagram/server/services/imgservice"
	"outstagram/server/services/notifbservice"
	"outstagram/server/services/postimgservice"
	"outstagram/server/services/postservice"
	"outstagram/server/services/rctableservice"
	"outstagram/server/services/rctservice"
	"outstagram/server/services/replyservice"
	"outstagram/server/services/storybservice"
	"outstagram/server/services/userservice"
	"outstagram/server/services/vwableservice"

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

		vwablerepo.New,
		vwableservice.New,

		userservice.New,
		userrepo.New,

		postservice.New,
		postrepo.New,

		postimgservice.New,
		postimgrepo.New,

		imgservice.New,
		imgrepo.New,

		cmtableservice.New,
		cmtablerepo.New,

		rctableservice.New,
		rctablerepo.New,

		db.New)
	return &postcontroller.Controller{}, nil
}

func InitializeReactController() (*rctcontroller.Controller, error) {
	wire.Build(
		rctcontroller.New,

		rctservice.New,
		rctrepo.New,

		rctableservice.New,
		rctablerepo.New,

		cmtablerepo.New,

		userservice.New,
		userrepo.New,

		db.New)

	return &rctcontroller.Controller{}, nil
}

func InitializeCommentableController() (*cmtablecontroller.Controller, error) {
	wire.Build(
		cmtablecontroller.New,

		cmtableservice.New,
		cmtablerepo.New,

		cmtservice.New,
		cmtrepo.New,

		userservice.New,
		userrepo.New,

		rctableservice.New,
		rctablerepo.New,

		replyservice.New,
		replyrepo.New,

		db.New)

	return &cmtablecontroller.Controller{}, nil
}

func InitializeMeController() (*mecontroller.Controller, error) {
	wire.Build(
		mecontroller.New,

		userservice.New,
		userrepo.New,

		postservice.New,
		postrepo.New,

		cmtableservice.New,
		cmtablerepo.New,

		storybservice.New,
		storybrepo.New,

		rctableservice.New,
		rctablerepo.New,

		imgservice.New,
		imgrepo.New,

		db.New)

	return &mecontroller.Controller{}, nil
}

func InitializeFollowController() (*flcontroller.Controller, error) {
	wire.Build(
		flcontroller.New,

		userservice.New,
		userrepo.New,

		db.New)

	return &flcontroller.Controller{}, nil
}

func InitializeStoryController() (*storycontroller.Controller, error) {
	wire.Build(
		storycontroller.New,

		imgservice.New,
		imgrepo.New,

		vwableservice.New,
		vwablerepo.New,

		storybservice.New,
		storybrepo.New,

		userservice.New,
		userrepo.New,

		db.New)

	return &storycontroller.Controller{}, nil
}
