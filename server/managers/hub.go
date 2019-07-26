package managers

import (
	"encoding/json"
	"fmt"
	"github.com/segmentio/ksuid"
	"log"
	"strings"
)

// hub maintains the set of active Connections and broadcasts WSMessages to the Connections.
type hub struct {
	// Inbound WSMessages from the Connections.
	BroadcastChannel chan ClientMessage

	// RegisterChannel requests from the Connections.
	RegisterChannel chan Subscription

	// UnregisterChannel requests from Connections.
	UnregisterChannel chan Subscription

	// All connections
	Connections map[*Connection]bool

	// A map of userID to Connections
	UserID2Connection map[uint][]*Connection

	// A map of roomID to Connections
	RoomID2Connection map[uint][]*Connection

	// Api provider
	APIProvider APIProvider

	// UID used for Pub/Sub
	UID string
}

// NewHub returns new hub instance
func NewHub() *hub {
	return &hub{
		BroadcastChannel:  make(chan ClientMessage),
		RegisterChannel:   make(chan Subscription),
		UnregisterChannel: make(chan Subscription),
		Connections:       make(map[*Connection]bool),
		UserID2Connection: make(map[uint][]*Connection),
		RoomID2Connection: make(map[uint][]*Connection),
		APIProvider:       NewLocalAPIProvider(),
		UID:               ksuid.New().String(),
	}
}

// Run starts a hub session
func (h *hub) run(wsMuxes ...func(from *SuperConnection, clientMessage Message)) {
	pubSub := pubSubClient.Subscribe("STORY", "ROOM")
	pubSubCh := pubSub.Channel()

	for {
		select {
		case s := <-h.RegisterChannel:
			h.Connections[s.SuperConn.Connection] = true
			roomIDs := h.APIProvider.GetUserRoomIDs(s.SuperConn.UserID)
			for _, roomID := range roomIDs {
				h.RoomID2Connection[roomID] = append(h.RoomID2Connection[roomID], s.SuperConn.Connection)
				log.Printf("User %v joins room %v", s.SuperConn.UserID, roomID)
			}

			log.Println("A client connected to server")

		case s := <-h.UnregisterChannel:
			if _, ok := h.Connections[s.SuperConn.Connection]; ok {
				delete(h.Connections, s.SuperConn.Connection)

				userConnections := h.UserID2Connection[s.SuperConn.UserID]
				for i, conn := range userConnections {
					if conn == s.SuperConn.Connection {
						userConnections = append(userConnections[:i], userConnections[i+1:]...)
					}
				}

				roomIDs := h.APIProvider.GetUserRoomIDs(s.SuperConn.UserID)
				for _, id := range roomIDs {
					for i, conn := range h.RoomID2Connection[id] {
						if conn == s.SuperConn.Connection {
							h.RoomID2Connection[id] = append(h.RoomID2Connection[id][:i], h.RoomID2Connection[id][i+1:]...)
							log.Printf("User %v leaves room %v\n", s.SuperConn.UserID, id)
						}
					}
				}

				log.Println("A client disconnected from server")
			}

		case cm := <-h.BroadcastChannel:
			for _, wsMux := range wsMuxes {
				log.Println("I handle message client sent to me")
				wsMux(cm.SuperConnection, cm.Message)
			}

			serverMessage := ServerMessage{Message: cm.Message, ActorID: cm.UserID, ServerUID: h.UID}
			sServerMessage, err := json.Marshal(&serverMessage)
			if err != nil {
				log.Println("Cannot marshal message: ", err.Error())
				break
			}

			channel := strings.Split(serverMessage.Type, ".")[0]
			if err := pubSubClient.Publish(channel, sServerMessage).Err(); err != nil {
				log.Println("Cannot publish message: ", err.Error())
				break
			}

		case sm := <-pubSubCh:
			var serverMessage ServerMessage
			if err := json.Unmarshal([]byte(sm.Payload), &serverMessage); err != nil {
				log.Println(err.Error())
				break
			}

			// If message comes from the same source, do nothing
			if serverMessage.ServerUID == h.UID {
				break
			}

			for _, wsMux := range wsMuxes {
				wsMux(&SuperConnection{UserID: serverMessage.ActorID}, serverMessage.Message)
			}
		}
	}
}

// Emit emits Message `m` to all sockets
func (h *hub) Emit(transmitMessage ServerMessage) {
	connections := h.Connections

	for c := range connections {
		c.Send <- transmitMessage
	}
}

// EmitTo emits Message `m` to all sockets in room `room`
func (h *hub) EmitTo(serverMessage ServerMessage, roomID uint) {
	connections := h.RoomID2Connection[roomID]
	for _, c := range connections {
		c.Send <- serverMessage
	}
}

// Broadcast broadcasts ClientMessage `m` to all sockets other than `conn` Connection
func (h *hub) Broadcast(conn *SuperConnection, serverMessage ServerMessage) {
	for c := range h.Connections {
		if c != conn.Connection {
			c.Send <- serverMessage
		}
	}
}

// BroadcastTo broadcasts ClientMessage `m` to all sockets in room `room` other than `conn` Connection
func (h *hub) BroadcastTo(conn *SuperConnection, serverMessage ServerMessage, roomID uint) {
	for _, c := range h.RoomID2Connection[roomID] {
		if c != conn.Connection {
			c.Send <- serverMessage
		}
	}
}

// BroadcastSelective broadcasts to specific connections collection
func (h *hub) BroadcastSelective(conn *SuperConnection, serverMessage ServerMessage, connections ...*Connection) {
	for i, c := range connections {
		if c != conn.Connection {
			c.Send <- serverMessage
			fmt.Println(i, "Send message", serverMessage)
		} else {
			fmt.Println("Do not send")
		}
	}
}

// WSMessageMultiplexer multiplexes message type and its corresponding handler
func (h *hub) WSMessageMultiplexer(c *Connection, serverMessage Message) {
	switch serverMessage.Type {

	}
}
