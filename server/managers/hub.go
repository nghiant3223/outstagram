package managers

import (
	"encoding/json"
	"fmt"
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
	}
}

// Run starts a hub session
func (h *hub) run(wsMuxes ...func(from *SuperConnection, clientMessage Message)) {
	pubSub := pubSubClient.Subscribe("STORY", "MESSAGE")
	pubSubCh := pubSub.Channel()

	for {
		select {
		case s := <-h.RegisterChannel:
			h.Connections[s.SuperConn.Connection] = true
			fmt.Println("A client connected to server")

		case us := <-h.UnregisterChannel:
			if _, ok := h.Connections[us.SuperConn.Connection]; ok {
				delete(h.Connections, us.SuperConn.Connection)

				userConnections := h.UserID2Connection[us.SuperConn.UserID]
				for i, conn := range userConnections {
					if conn == us.SuperConn.Connection {
						userConnections = append(userConnections[:i], userConnections[i+1])
					}
				}
				// TODO: remove us.SuperConn from rooms in which it joins
				fmt.Println("A client disconnected from server")
			}

		case cm := <-h.BroadcastChannel:
			for _, wsMux := range wsMuxes {
				fmt.Println("I handle message client sent to me")
				wsMux(cm.SuperConnection, cm.Message)
			}

			serverMessage := ServerMessage{Message: cm.Message, ActorID: cm.UserID}
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
			fmt.Println("I published message from PubSubChannel")

		case sm := <-pubSubCh:
			var serverMessage ServerMessage
			if err := json.Unmarshal([]byte(sm.Payload), &serverMessage); err != nil {
				log.Println(err.Error())
				break
			}
			fmt.Println("I received message from PubSubChannel")
			for _, wsMux := range wsMuxes {
				fmt.Println("I handle message received from PubSubChannel")
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
	for _, c := range connections {
		if c != conn.Connection {
			c.Send <- serverMessage
		} else {
			fmt.Println("Duplicate")
		}
	}
}

// WSMessageMultiplexer multiplexes message type and its corresponding handler
func (h *hub) WSMessageMultiplexer(c *Connection, serverMessage Message) {
	switch serverMessage.Type {

	}
}
