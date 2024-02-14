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


// func main() {
//     apiService := utils.CreateAPIService("eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiI4ZjY4NjcxMC04ZjIwLTAxM2MtYmQ4Yy0yYWNjZjk1NjI0ZjIiLCJpc3MiOiJnYW1lbG9ja2VyIiwiaWF0IjoxNzA0NTg2ODYxLCJwdWIiOiJibHVlaG9sZSIsInRpdGxlIjoicHViZyIsImFwcCI6IjkxZWU0Y2M4LWY2MTEtNDUxOS05NjQ1LTEzNWJlN2Y4NjkyMiJ9.tadtjH48XZwaRKRSh-ROjHFHuO2dpjrzYASkixz-px0")
//     accountId, err:= apiService.GetAccountId("PracticeofEno2")
//     fmt.Println(accountId)
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//     repositories.AddUserCurrentKillByApiKey(apiService.ApiKey, 5)
//     // ids, err := apiService.GetLastMatchId()
//     // if err != nil {
//     //     fmt.Println(err)
//     //     return
//     // }
//     // fmt.Printf("last match id : %s\n", ids)

//     // userData, err := apiService.GetMatchData(ids)
//     // if err != nil {
//     //     fmt.Println(err)
//     //     return
//     // }
//     // fmt.Println(userData)
    
//     // fmt.Println(matchData)
    
// 	// client := utils.GetClient()
// 	// ctx := context.Background()
    
// }