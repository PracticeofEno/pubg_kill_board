package utils

import (
	"log"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

type Hub struct {
    Clients map[string]socketio.Conn
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

var hub = &Hub{
    Clients: make(map[string]socketio.Conn),
}

func SetupWebscoket(router *gin.Engine, server *socketio.Server) {
    server.OnConnect("/", func(s socketio.Conn) error {
        s.SetContext("")
        log.Println("connected:", s.ID())
        return nil
    })

    server.OnEvent("/", "random_string", func (s socketio.Conn, msg string) (error) {
        log.Println("hid")
        log.Println(msg)
        hub.SetWsClient(msg, s)
        // s.Emit("pong2","pong")
        return nil
    })

    server.OnError("/", func(s socketio.Conn, e error) {
        log.Println("meet error:", e)
    })

    server.OnDisconnect("/", func(s socketio.Conn, reason string) {
        log.Println("closed", reason)
    })

    router.GET("/socket.io/*any", gin.WrapH(server))
    router.POST("/socket.io/*any", gin.WrapH(server))

    go func() {
        if err := server.Serve(); err != nil {
            log.Fatalf("socketio listen error: %s\n", err)
        }
    }()
}

