package managers

import (
	"log"
	"outstagram/server/db"
	"outstagram/server/repos/userrepo"
	"outstagram/server/services/userservice"
)

// storyManager manages WebSocket events related to Story
type storyManager struct {
	Hub         *hub
	userService *userservice.UserService
}

// NewStoryManager returns new storyManager
func NewStoryManager(hub *hub) *storyManager {
	dbConn, _ := db.New()
	userRepo := userrepo.New(dbConn)
	userService := userservice.New(userRepo)

	return &storyManager{Hub: hub, userService: userService}
}

// WSMux multiplexes WebSocket event to corresponding handler
func (sm *storyManager) WSMux(c *Connection, transmitData ClientMessage) {
	switch transmitData.Type {
	case "STORY.CLIENT.POST_STORY":
		var followerConnections []*Connection
		for _, user := range sm.userService.GetFollowers(c.UserID) {
			if connection, ok := sm.Hub.UserID2Connection[user.ID]; ok {
				followerConnections = append(followerConnections, connection)
			}
		}

		x := ServerMessage{Data: transmitData.Data, Type: "STORY.SERVER.POST_STORY", ActorID: &c.UserID}
		sm.Hub.BroadcastSelective(c, x, followerConnections)
	default:
		log.Fatal("Event not supported")
	}
}
