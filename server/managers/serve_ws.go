package managers

import (
	"log"
	"outstagram/server/utils"

	"github.com/gin-gonic/gin"
)

// ServeWs handles websocket requests from the peer.
func ServeWs(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	var superConnection *SuperConnection
	if userID, err := utils.StringToUint(c.Query("userID")); err != nil {
		connection := &Connection{Send: make(chan ServerMessage), WS: ws}
		superConnection = &SuperConnection{Connection: connection, UserID: 0}
	} else {
		connection := &Connection{Send: make(chan ServerMessage), WS: ws}
		superConnection = &SuperConnection{Connection: connection, UserID: userID}
		Hub.UserID2Connection[userID] = append(Hub.UserID2Connection[userID], connection)
	}

	subscription := Subscription{SuperConn: superConnection}
	Hub.RegisterChannel <- subscription

	go subscription.WritePump()
	go subscription.ReadPump()
}
