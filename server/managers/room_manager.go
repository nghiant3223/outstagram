package managers

import (
	"outstagram/server/db"
	"outstagram/server/repos/userrepo"
	"outstagram/server/services/userservice"
)

// storyManager manages WebSocket events related to Story
type roomManager struct {
	Hub         *hub
	userService *userservice.UserService
}

// NewStoryManager returns new storyManager
func NewRoomManager(hub *hub) *roomManager {
	dbConn, _ := db.New()
	userRepo := userrepo.New(dbConn)
	userService := userservice.New(userRepo)
	return &roomManager{Hub: hub, userService: userService}
}

// WSMux multiplexes WebSocket event to corresponding handler
func (rm *roomManager) WSMux(c *SuperConnection, clientMessage Message) {
	switch clientMessage.Type {
	case "ROOM.CLIENT.SEND_MESSAGE":
	}
}
