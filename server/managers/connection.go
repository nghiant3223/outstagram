package managers

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

// Connection is an middleman between the websocket Connection and the hub.
type Connection struct {
	// The websocket Connection.
	WS *websocket.Conn

	// Buffered channel of outbound WSMessages.
	Send chan ServerMessage
}

// Write writes a message with the given message type and payload.
func (c *Connection) Write(mt int, payload []byte) error {
	err := c.WS.SetWriteDeadline(time.Now().Add(writeWait))
	if err != nil {
		fmt.Println("4.", err.Error())
	}
	return c.WS.WriteMessage(mt, payload)
}

// SuperConnection is wrapper struct of Connection
// It's used to add UserID to the Connection
// If SuperConnection's Connection is nil, the Connection is between servers
// Otherwise, the Connection is between client and server
type SuperConnection struct {
	*Connection
	UserID uint
}
