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
func (sm *storyManager) WSMux(c *SuperConnection, clientMessage Message) {
	switch clientMessage.Type {
	case "STORY.CLIENT.POST_STORY":
		var followerConnections []*Connection
		for _, user := range sm.userService.GetFollowers(c.UserID) {
			if connections, ok := sm.Hub.UserID2Connection[user.ID]; ok {
				followerConnections = append(followerConnections, connections...)
			}
		}

		message := Message{Data: clientMessage.Data, Type: "STORY.SERVER.POST_STORY"}
		serverMessage := ServerMessage{Message: message, ActorID: c.UserID}
		sm.Hub.BroadcastSelective(c, serverMessage, followerConnections...)

	case "STORY.CLIENT.REACT_STORY":
		targetUserID := uint(clientMessage.Data.(map[string]interface{})["targetUserID"].(float64))
		message := Message{Data: clientMessage.Data, Type: "STORY.SERVER.REACT_STORY"}
		serverMessage := ServerMessage{Message: message, ActorID: c.UserID}
		sm.Hub.BroadcastSelective(c, serverMessage, sm.Hub.UserID2Connection[targetUserID]...)

	case "STORY.CLIENT.UNREACT_STORY":
		targetUserID := uint(clientMessage.Data.(map[string]interface{})["targetUserID"].(float64))
		message := Message{Data: clientMessage.Data, Type: "STORY.SERVER.UNREACT_STORY"}
		serverMessage := ServerMessage{Message: message, ActorID: c.UserID}
		sm.Hub.BroadcastSelective(c, serverMessage, sm.Hub.UserID2Connection[targetUserID]...)

	default:
		log.Println("Event not supported")
	}
}
