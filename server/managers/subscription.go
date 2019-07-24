package managers

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

// Subscription is the wrapper for socket subscription
type Subscription struct {
	SuperConn *SuperConnection
}

// ReadPump pumps messages from the websocket Connection to the hub.
func (s Subscription) ReadPump() {
	c := s.SuperConn
	defer func() {
		Hub.UnregisterChannel <- s
		c.WS.Close()
	}()

	c.WS.SetReadLimit(maxMessageSize)
	_ = c.WS.SetReadDeadline(time.Now().Add(pongWait))
	c.WS.SetPongHandler(func(string) error { _ = c.WS.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		var clientMessage Message

		err := c.WS.ReadJSON(&clientMessage)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("Error: %v", err)
			}
			break
		}

		m := ClientMessage{Message: clientMessage, SuperConnection: c}
		Hub.BroadcastChannel <- m
	}
}

// WritePump pumps messages from the hub to the websocket Connection.
func (s *Subscription) WritePump() {
	c := s.SuperConn
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		_ = c.WS.Close()
	}()

	for {
		select {
		case transmitMessageDTO, ok := <-c.Send:
			if !ok {
				err := c.Write(websocket.CloseMessage, []byte("Connection closed"))
				if err != nil {
					log.Println(err.Error())
				}
				return
			}

			if transmitMessageJSON, err := json.Marshal(transmitMessageDTO); err != nil {
				log.Println(err.Error())
				return
			} else if err := c.Write(websocket.TextMessage, transmitMessageJSON); err != nil {
				log.Println(err.Error())
				return
			}

		case <-ticker.C:
			if err := c.Write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}
