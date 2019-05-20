package managers

import (
	"time"
	"log"
	"encoding/json"
	
	"github.com/gorilla/websocket"
)	

// Subscription is the wrapper for socket subscription
type Subscription struct {
	Conn *Connection
}

// ReadPump pumps messages from the websocket Connection to the hub.
func (s Subscription) ReadPump() {
	c := s.Conn
	defer func() {
		HubInstance.Unregister <- s
		c.WS.Close()
	}()
	c.WS.SetReadLimit(maxMessageSize)
	c.WS.SetReadDeadline(time.Now().Add(pongWait))
	c.WS.SetPongHandler(func(string) error { c.WS.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		var transmitData TransmitData
		err := c.WS.ReadJSON(&transmitData)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("Error: %v", err)
			}
			break
		}
		m := WSMessage{TransmitData: transmitData, From: c}
		HubInstance.BroadcastChannel <- m
	}
}

// WritePump pumps messages from the hub to the websocket Connection.
func (s *Subscription) WritePump() {
	c := s.Conn
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.WS.Close()
	}()
	for {
		select {
		case transmitData, ok := <-c.Send:
			if !ok {
				c.Write(websocket.CloseMessage, []byte{})
				return
			}
			if transmitDataJSON, errJSON := json.Marshal(transmitData); errJSON != nil {
				return
			} else if err := c.Write(websocket.TextMessage, transmitDataJSON); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.Write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}