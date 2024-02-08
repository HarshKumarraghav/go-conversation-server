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

// The code is creating a variable `H` of type `Hub` and initializing it with a struct literal. The
// struct literal sets the values of the fields `rooms`, `broadcast`, `register`, and `unregister`.
var H = Hub{
	rooms:      make(map[string]map[*connection]bool),
	broadcast:  make(chan message),
	register:   make(chan subscription),
	unregister: make(chan subscription),
}

func (h *Hub) Hub() {
	for {
		select {
		case s := <-h.register:
			connections := h.rooms[s.room]
			if connections == nil {
				connections = make(map[*connection]bool)
				h.rooms[s.room] = connections
			}
		case s := <-h.unregister:
			connections := h.rooms[s.room]
			if connections != nil {
				if _, ok := connections[s.conn]; ok {
					delete(connections, s.conn)
					close(s.conn.send)
					if len(connections) == 0 {
						delete(h.rooms, s.room)
					}
				}
			}
		case m := <-h.broadcast:
			connections := h.rooms[m.room]
			for c := range connections {
				select {
				case c.send <- m.data:
				default:
					close(c.send)
					delete(connections, c)
					if len(connections) == 0 {
						delete(h.rooms, m.room)
					}
				}
			}
		}
	}
}
