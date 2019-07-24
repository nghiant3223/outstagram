package managers

import (
	"github.com/go-redis/redis"
	"net/http"
	"outstagram/server/db"
	"time"

	"github.com/gorilla/websocket"
)

// Message is data transmitted between peers
type Message struct {
	Data interface{} `json:"data"`
	Type string      `json:"type"`
}

// ClientMessage is wrapper for data transmitted from client to server
type ClientMessage struct {
	Message
	*SuperConnection
}

// TransmitDataDTO is transmitted from server to client
type ServerMessage struct {
	Message
	ActorID uint `json:"actorID,omitempty"`
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
var RoomManager *roomManager
var pubSubClient *redis.Client

func Init() {
	Hub = NewHub()
	StoryManager = NewStoryManager(Hub)
	RoomManager = NewRoomManager(Hub)
	pubSubClient = db.NewRedisClient()
}

func (h *hub) Run() {
	h.run(StoryManager.WSMux, RoomManager.WSMux)
}
