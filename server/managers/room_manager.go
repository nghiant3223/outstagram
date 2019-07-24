package managers

// storyManager manages WebSocket events related to Story
type roomManager struct {
	Hub *hub
}

// NewStoryManager returns new storyManager
func NewRoomManager(hub *hub) *roomManager {
	return &roomManager{Hub: hub}
}

// WSMux multiplexes WebSocket event to corresponding handler
func (rm *roomManager) WSMux(c *SuperConnection, clientMessage Message) {
	switch clientMessage.Type {
	case "ROOM.CLIENT.SEND_MESSAGE":
	}
}
