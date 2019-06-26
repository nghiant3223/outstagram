package managers

import (
	"time"

	"github.com/gorilla/websocket"
)

// Connection is an middleman between the websocket Connection and the hub.
type Connection struct {
	// The websocket Connection.
	WS *websocket.Conn

	// Buffered channel of outbound WSMessages.
	Send chan TransmitMessageDTO

	// UserID of who trigger the event
	UserID *uint
}

// Write writes a message with the given message type and payload.
func (c *Connection) Write(mt int, payload []byte) error {
	c.WS.SetWriteDeadline(time.Now().Add(writeWait))
	return c.WS.WriteMessage(mt, payload)
}
