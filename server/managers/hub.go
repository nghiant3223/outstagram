package managers

// Hub maintains the set of active Connections and broadcasts WSMessages to the Connections.
type Hub struct {
	// Registered Connections.
	Rooms map[string]map[*Connection]bool

	Connections map[*Connection]bool

	// Inbound WSMessages from the Connections.
	BroadcastChannel chan WSMessage

	// Register requests from the Connections.
	Register chan Subscription

	// Unregister requests from Connections.
	Unregister chan Subscription
}

// NewHub returns new Hub instance
func NewHub() *Hub {
	return &Hub{
		BroadcastChannel: make(chan WSMessage),
		Register:         make(chan Subscription),
		Unregister:       make(chan Subscription),
		Rooms:            make(map[string]map[*Connection]bool),
		Connections:      make(map[*Connection]bool),
	}

}

// Run starts a hub session
func (h *Hub) Run(wsMuxes ...func(from *Connection, transmitData TransmitData)) {
	for {
		select {
		case s := <-h.Register:
			h.Connections[s.Conn] = true
		case s := <-h.Unregister:
			if _, ok := h.Connections[s.Conn]; ok {
				delete(h.Connections, s.Conn)
				// TODO: remove s.Conn from rooms in which it joins
			}
		case m := <-h.BroadcastChannel:
			for _, wsMux := range wsMuxes {
				wsMux(m.From, m.TransmitData)
			}
		}
	}
}

// Emit emits TransmitData `m` to all sockets
func (h *Hub) Emit(transmitData TransmitData) {
	connections := h.Connections
	for c := range connections {
		c.Send <- transmitData
	}
}

// EmitTo emits TransmitData `m` to all sockets in room `room`
func (h *Hub) EmitTo(transmitData TransmitData, room string) {
	connections := h.Rooms[room]
	for c := range connections {
		c.Send <- transmitData
	}
}

// Broadcast broadcasts WSMessage `m` to all sockets other than `conn` Connection
func (h *Hub) Broadcast(conn *Connection, transmitData TransmitData) {
	for c := range h.Connections {
		if c != conn {
			c.Send <- transmitData
		}
	}
}

// BroadcastTo broadcasts WSMessage `m` to all sockets in room `room` other than `conn` Connection
func (h *Hub) BroadcastTo(conn *Connection, transmitData TransmitData, room string) {
	for c := range h.Rooms[room] {
		if c != conn {
			c.Send <- transmitData
		}
	}
}

// BroadcastSelective broadcasts to specific connections collection
func (h *Hub) BroadcastSelective(conn *Connection, trasmitData TransmitData, connections []*Connection) {
	for _, c := range connections {
		if c != conn {
			c.Send <- trasmitData
		}
	}
}

// WSMessageMultiplexer multiplexes message type and its corresponding handler
func (h *Hub) WSMessageMultiplexer(from *Connection, transmitData TransmitData) {
	switch transmitData.Type {
		
	}
}