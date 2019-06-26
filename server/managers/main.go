package managers

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// TransmitData is data transmitted between peers
type TransmitData struct {
	Data interface{} `json:"data"`
	Type string      `json:"type"`
}

// TransmitMessage is wrapper for data transmitted from client to server
type TransmitMessage struct {
	TransmitData TransmitData
	Connection   *Connection
}

// TransmitDataDTO is transmitted from server to client
type TransmitMessageDTO struct {
	Data    interface{} `json:"data"`
	Type    string      `json:"type"`
	ActorID *uint       `json:"actorID"`
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var Hub = NewHub()
var StoryManager = NewStoryManager(Hub)
