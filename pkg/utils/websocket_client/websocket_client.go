package websocket_client

import (
	"fmt"
	"kill_board/internal/app/repositories"
	service "kill_board/internal/app/services"
	"kill_board/pkg/classes/auto_kill"
	"kill_board/pkg/utils/hub"
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

func SetupWebscoket(router *gin.Engine, server *socketio.Server) {

    router.GET("/gorutine", func(c *gin.Context) {
		randomString := c.Query("random_string")
		user, err := service.GetUserByRandomStringWithRelation(randomString)
		if err != nil {
			fmt.Println(err)
			c.JSON(400, gin.H{
				"error": "not found",
			})
		} else {
			worker := auto_kill.NewWorker(user.APIKey, user.NickName, randomString);
			go worker.Run()
		}
    })

    router.GET("/webhook/:random", func(ctx *gin.Context) {
		var tf = false;
        randomString := ctx.Param("random")
		user, err := service.GetUserByRandomStringWithRelation(randomString)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(400, gin.H{
				"error": fmt.Sprintf("%s - can't find user", randomString),
			})
		}
		randGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
		for _, v := range user.Percents() {
			r := randGenerator.Float64()
			if r < float64(v.Percent) * 0.01 {
				repositories.AddUserTargetKillByRandomString(randomString, v.Count)
				server.BroadcastToRoom("/", randomString, "pong2", "haha")
				ctx.JSON(200, gin.H{
					"message": fmt.Sprintf("%d킬", v.Count),
				})
				tf = true;
				break;
			}
		}
		if !tf {
			ctx.JSON(200, gin.H{
				"message": fmt.Sprintf("%d킬", 0),
			})
		}
    })

    server.OnConnect("/", func(s socketio.Conn) error {
        s.SetContext("")
        log.Println("connected:", s.ID())
        return nil
    })

    server.OnEvent("/", "random_string", func (s socketio.Conn, msg string) (error) {
        log.Println("hid")
        log.Println(msg)
        hub.GetHub().SetWsClient(msg, s)
        s.Join(msg)
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

