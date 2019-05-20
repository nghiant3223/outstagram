package managers

import (
	"log"
	"net/http"
)

// ServeWs handles websocket requests from the peer.
func ServeWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	c := &Connection{Send: make(chan TransmitData), WS: ws}
	s := Subscription{c}

	HubInstance.Register <- s
	go s.WritePump()
	s.ReadPump()
}