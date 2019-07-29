package managers

import (
	"fmt"
)

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
	fmt.Println(clientMessage)
	switch clientMessage.Type {
	case "ROOM.CLIENT.SEND_MESSAGE":
		targetRoomID := uint(clientMessage.Data.(map[string]interface{})["targetRoomID"].(float64))
		message := Message{Data: clientMessage.Data, Type: "ROOM.SERVER.SEND_MESSAGE"}
		serverMessage := ServerMessage{Message: message, ActorID: c.UserID}
		roomConnections := rm.Hub.RoomID2Connection[targetRoomID]
		rm.Hub.BroadcastSelective(c, serverMessage, roomConnections...)

	case "ROOM.CLIENT.TYPING":
		fmt.Println("typing on server")
		targetRoomID := uint(clientMessage.Data.(map[string]interface{})["targetRoomID"].(float64))
		message := Message{Data: clientMessage.Data, Type: "ROOM.SERVER.TYPING"}
		serverMessage := ServerMessage{Message: message, ActorID: c.UserID}
		roomConnections := rm.Hub.RoomID2Connection[targetRoomID]
		rm.Hub.BroadcastSelective(c, serverMessage, roomConnections...)

	case "ROOM.CLIENT.STOP_TYPING":
		targetRoomID := uint(clientMessage.Data.(map[string]interface{})["targetRoomID"].(float64))
		message := Message{Data: clientMessage.Data, Type: "ROOM.SERVER.STOP_TYPING"}
		serverMessage := ServerMessage{Message: message, ActorID: c.UserID}
		roomConnections := rm.Hub.RoomID2Connection[targetRoomID]
		rm.Hub.BroadcastSelective(c, serverMessage, roomConnections...)

	case "ROOM.CLIENT.CREATE_ROOM":
		targetRoomID := uint(clientMessage.Data.(map[string]interface{})["targetRoomID"].(float64))
		anotherMemberID := uint(clientMessage.Data.(map[string]interface{})["anotherMemberID"].(float64))

		for _, conn := range rm.Hub.UserID2Connection[c.UserID] {
			rm.Hub.RoomID2Connection[targetRoomID] = append(rm.Hub.RoomID2Connection[targetRoomID], conn)
		}
		for _, conn := range rm.Hub.UserID2Connection[anotherMemberID] {
			rm.Hub.RoomID2Connection[targetRoomID] = append(rm.Hub.RoomID2Connection[targetRoomID], conn)
		}

		message := Message{Data: clientMessage.Data, Type: "ROOM.SERVER.CREATE_ROOM"}
		serverMessage := ServerMessage{Message: message, ActorID: c.UserID}
		rm.Hub.BroadcastSelective(c, serverMessage, rm.Hub.RoomID2Connection[targetRoomID]...)
	}
}
