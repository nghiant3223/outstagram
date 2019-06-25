package managers

import "log"

// storyManager manages WebSocket events related to Story
type storyManager struct {
	Hub *hub
}

// NewStoryManager returns new storyManager
func NewStoryManager(hub *hub) *storyManager {
	return &storyManager{Hub: hub}
}

// WSMux multiplexes WebSocket event to corresponding handler
func (sm *storyManager) WSMux(from *Connection, transmitData TransmitData) {
	switch transmitData.Type {
	case "STORY.CLIENT.POST_STORY":
		sm.Hub.Emit(TransmitData{Data: "A new story", Type: "STORY.SERVER.POST_STORY"})
	default:
		log.Fatal("Event not supported")
	}
}
