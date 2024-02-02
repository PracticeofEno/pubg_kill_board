package router

import "github.com/gin-gonic/gin"
func SetupAPIRoutes(apiGroup *gin.RouterGroup) {
    // API 라우트 설정
	apiGroup.GET("/", func (c *gin.Context)  {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
}