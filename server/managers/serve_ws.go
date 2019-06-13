package managers

import (
	"log"

	"github.com/gin-gonic/gin"
)

// ServeWs handles websocket requests from the peer.
func ServeWs(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	conn := &Connection{Send: make(chan TransmitData), WS: ws}
	s := Subscription{conn}

	HubInstance.Register <- s
	go s.WritePump()
	go s.ReadPump()
}
