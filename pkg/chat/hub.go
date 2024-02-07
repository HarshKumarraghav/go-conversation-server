package chat

import (
	"github.com/gorilla/websocket"
)

type message struct {
	data []byte
	room string
}

type connection struct {
	ws   *websocket.Conn
	send chan []byte
}

type subscription struct {
	conn *connection
	room string
}

type Hub struct {
	rooms      map[string]map[*connection]bool
	broadcast  chan message
	register   chan subscription
	unregister chan subscription
}

var H = Hub{
	rooms:      make(map[string]map[*connection]bool),
	broadcast:  make(chan message),
	register:   make(chan subscription),
	unregister: make(chan subscription),
}
