package managers

import "log"

// StoryManager manages WebSocket events related to Story
type StoryManager struct {
	Hub *Hub
}

// NewStoryManager returns new StoryManager
func NewStoryManager(hub *Hub) *StoryManager {
	return &StoryManager{Hub: hub}
}

// WSMux multiplexes WebSocket event to corresponding handler
func (sm *StoryManager) WSMux(from *Connection, transmitData TransmitData) {
	switch transmitData.Type {
	case "STORY.CLIENT.POST_STORY":
		sm.Hub.Emit(TransmitData{Data: "A new story", Type: "STORY.SERVER.POST_STORY"})
	default:
		log.Fatal("Event not supported")
	}
}
