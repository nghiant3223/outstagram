package managers

import (
	"github.com/go-redis/redis"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// ClientMessage is data transmitted between peers
type ClientMessage struct {
	Data interface{} `json:"data"`
	Type string      `json:"type"`
}

// ClientMessageWrapper is wrapper for data transmitted from client to server
type ClientMessageWrapper struct {
	TransmitData ClientMessage
	Connection   *Connection
}

// TransmitDataDTO is transmitted from server to client
type ServerMessage struct {
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
	maxMessageSize = 512 * 1024
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024 * 1024,
	WriteBufferSize: 1024 * 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var Hub *hub
var StoryManager *storyManager
var pubSubClient *redis.Client

func Init() {
	Hub = NewHub()
	StoryManager = NewStoryManager(Hub)

	pubSubClient = redis.NewClient(&redis.Options{
		Addr:     "cache:6379",
		Password: "",
		DB:       0,
	})
}
