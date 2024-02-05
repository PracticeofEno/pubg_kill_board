package handlers

import (
	service "kill_board/internal/app/services"

	"github.com/gin-gonic/gin"
)

func SetupUsersRoute(apiGroup *gin.RouterGroup) {
    // API 라우트 설정
    apiGroup.GET("/", func (c *gin.Context)  {
        c.JSON(200, gin.H{
            "message": "Hello, World!",
        })
    })

    apiGroup.GET("/id", func (c *gin.Context) {
        apiKey := c.Query("api_key")
        if apiKey == "" {
            c.JSON(400, gin.H{
                "error": "api_key is null",
            })
        } else {
            id, err := service.CheckOrCreateUserByAPIKey(apiKey)
            if err != nil {
                c.JSON(400, gin.H{
                    "error": "invalid api key",
                })
            }
            c.JSON(200, gin.H{
                "id": id,
            })
        }
    })

    apiGroup.GET("/random_string", func(ctx *gin.Context) {
        // apiKey := c.Query("api_key")
        
    })
}