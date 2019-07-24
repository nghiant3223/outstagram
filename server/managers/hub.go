package managers

import (
	"fmt"
	"log"
)

// hub maintains the set of active Connections and broadcasts WSMessages to the Connections.
type hub struct {
	// Registered Connections.
	RoomID2Connection map[string]map[*Connection]bool

	Connections map[*Connection]bool

	// Inbound WSMessages from the Connections.
	BroadcastChannel chan ClientMessageWrapper

	// Register requests from the Connections.
	Register chan Subscription

	// Unregister requests from Connections.
	Unregister chan Subscription

	UserID2Connection map[uint][]*Connection
}

// NewHub returns new hub instance
func NewHub() *hub {
	return &hub{
		BroadcastChannel:  make(chan ClientMessageWrapper),
		Register:          make(chan Subscription),
		Unregister:        make(chan Subscription),
		RoomID2Connection: make(map[string]map[*Connection]bool),
		Connections:       make(map[*Connection]bool),
		UserID2Connection: make(map[uint][]*Connection),
	}
}

// Run starts a hub session
func (h *hub) Run(wsMuxes ...func(from *Connection, clientMessage ClientMessage)) {
	pubSub := pubSubClient.Subscribe("story")
	ch := pubSub.Channel()

	go func() {
		for msg := range ch {
			fmt.Println(">>>", msg.Channel, msg.Payload)
		}
	}()

	for {
		select {
		case s := <-h.Register:
			fmt.Println("A client connected to server")
			h.Connections[s.Conn] = true
		case s := <-h.Unregister:
			if _, ok := h.Connections[s.Conn]; ok {
				fmt.Println("A client disconnected from server")
				delete(h.Connections, s.Conn)
				delete(h.UserID2Connection, s.Conn.UserID)
				// TODO: remove s.Conn from rooms in which it joins
			}
		case m := <-h.BroadcastChannel:
			for _, wsMux := range wsMuxes {
				wsMux(m.Connection, m.ClientMessage)
			}

			serverTransferMessage := ServerMessage{Data: m.Data, Type: m.Type, ActorID: m.Connection.UserID}
			err := pubSubClient.Publish("story", &serverTransferMessage).Err()
			if err != nil {
				log.Println("Cannot publish message", err.Error())
			}

			//case m := <-ch:

		}
	}
}

// Emit emits ClientMessage `m` to all sockets
func (h *hub) Emit(transmitMessage ServerMessage) {
	connections := h.Connections

	for c := range connections {
		c.Send <- transmitMessage
	}
}

// EmitTo emits ClientMessage `m` to all sockets in room `room`
func (h *hub) EmitTo(transmitMessage ServerMessage, room string) {
	connections := h.RoomID2Connection[room]
	for c := range connections {
		c.Send <- transmitMessage
	}
}

// Broadcast broadcasts ClientMessageWrapper `m` to all sockets other than `conn` Connection
func (h *hub) Broadcast(conn *Connection, transmitMessage ServerMessage) {
	for c := range h.Connections {
		if c != conn {
			c.Send <- transmitMessage
		}
	}
}

// BroadcastTo broadcasts ClientMessageWrapper `m` to all sockets in room `room` other than `conn` Connection
func (h *hub) BroadcastTo(conn *Connection, transmitMessage ServerMessage, room string) {
	for c := range h.RoomID2Connection[room] {
		if c != conn {
			c.Send <- transmitMessage
		}
	}
}

// BroadcastSelective broadcasts to specific connections collection
func (h *hub) BroadcastSelective(conn *Connection, transmitMessage ServerMessage, connections ...*Connection) {
	for _, c := range connections {
		if c != conn {
			c.Send <- transmitMessage
		}
	}
}

// WSMessageMultiplexer multiplexes message type and its corresponding handler
func (h *hub) WSMessageMultiplexer(c *Connection, transmitData ClientMessage) {
	switch transmitData.Type {

	}
}
