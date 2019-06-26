package managers

import (
	"log"
)

// storyManager manages WebSocket events related to Story
type storyManager struct {
	Hub *hub
}

// NewStoryManager returns new storyManager
func NewStoryManager(hub *hub) *storyManager {
	return &storyManager{Hub: hub}
}

// WSMux multiplexes WebSocket event to corresponding handler
func (sm *storyManager) WSMux(c *Connection, transmitData TransmitData) {
	switch transmitData.Type {
	case "STORY.CLIENT.POST_STORY":
		sm.Hub.Broadcast(c, TransmitMessageDTO{Data:transmitData.Data, Type:"STORY.SERVER.POST_STORY",ActorID:c.UserID})
	default:
		log.Fatal("Event not supported")
	}
}
