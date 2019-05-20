package managers

import "time"

// Write writes a message with the given message type and payload.
func (c *Connection) Write(mt int, payload []byte) error {
	c.WS.SetWriteDeadline(time.Now().Add(writeWait))
	return c.WS.WriteMessage(mt, payload)
}