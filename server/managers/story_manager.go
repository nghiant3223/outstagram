package managers

// storyManager manages WebSocket events related to Story
type storyManager struct {
	Hub *hub
}

// NewStoryManager returns new storyManager
func NewStoryManager(hub *hub) *storyManager {
	return &storyManager{Hub: hub}
}

// WSMux multiplexes WebSocket event to corresponding handler
func (sm *storyManager) WSMux(c *SuperConnection, clientMessage Message) {
	switch clientMessage.Type {
	case "STORY.CLIENT.POST_STORY":
		var followerConnections []*Connection
		for _, follower := range sm.Hub.APIProvider.GetUserFollowers(c.UserID) {
			if connections, ok := sm.Hub.UserID2Connection[follower.ID]; ok {
				followerConnections = append(followerConnections, connections...)
			}
		}

		userConnections := sm.Hub.UserID2Connection[c.UserID]
		message := Message{Data: clientMessage.Data, Type: "STORY.SERVER.POST_STORY"}
		serverMessage := ServerMessage{Message: message, ActorID: c.UserID}
		sm.Hub.BroadcastSelective(c, serverMessage, append(userConnections, followerConnections...)...)

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
	}
}
