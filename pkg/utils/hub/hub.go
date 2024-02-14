package hub

import (
	socketio "github.com/googollee/go-socket.io"
)

type Hub struct {
    Clients map[string]socketio.Conn
}

var hub = &Hub{
    Clients: make(map[string]socketio.Conn),
}

func GetHub() *Hub {
    return hub
}

func (h *Hub) GetWsClient(key string) socketio.Conn {
    return h.Clients[key]
}

func (h *Hub) SetWsClient(key string, value socketio.Conn) {
	h.Clients[key] = value
}

