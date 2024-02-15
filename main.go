package main

import (
	"kill_board/internal/app/handlers"
	"kill_board/pkg/utils/websocket_client"
	"net/http"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

func GinMiddleware(allowOrigin string) gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

        if c.Request.Method == http.MethodOptions {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }

        c.Request.Header.Del("Origin")

        c.Next()
    }
}


func main() {
    r := gin.Default();
    usersGroup := r.Group("/users");
    handlers.SetupUsersRoute(usersGroup);
    r.Use(GinMiddleware("*"))

    server := socketio.NewServer(nil)
    websocket_client.SetupWebscoket(r, server)
    defer server.Close()
    
    r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080;
}