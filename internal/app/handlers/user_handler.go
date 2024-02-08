package handlers

import (
	"fmt"
	service "kill_board/internal/app/services"
	"kill_board/pkg/utils"

	"github.com/gin-gonic/gin"
)

func SetupUsersRoute(apiGroup *gin.RouterGroup) {
    // API 라우트 설정
    apiGroup.GET("/exist", func(ctx *gin.Context) {
        randomString := ctx.Query("random_string")
        exist := service.ExistRandomString(randomString)
        ctx.JSON(200, gin.H{
            "exist": exist,
        })
    })

    apiGroup.GET("/user", func (c *gin.Context)  {
        random_string := c.Query("random_string")
        user, err := service.GetUserByRandomStringWithRelation(random_string)
        if err != nil {
            fmt.Println(err)
            c.JSON(400, gin.H{
                "error": "not found",
            })
        } else {
            c.JSON(200, gin.H{
                "user": user,
            })
        }
    })

    apiGroup.POST("/user", func (c *gin.Context) {
        objTmp := utils.UpdateUserData{}
        errA := c.ShouldBind(&objTmp);
        if errA != nil {
            c.JSON(400, gin.H{
                "error": "bad request - invalid json format",
            })
            return
        }
        service.UpdateUserData(objTmp);
    })

    apiGroup.GET("/random_string", func(c *gin.Context) {
        apiKey := c.Query("api_key")
        randomString, err := service.CheckOrCreateUserByAPIKey(apiKey)
        if err != nil {
            fmt.Println(err)
            c.JSON(400, gin.H{
                "error": "not found",
            })
        } else {
            c.JSON(200, gin.H{
                "random_string": randomString,
            })
        }
    })

    apiGroup.POST("/percent", func(c *gin.Context) {
        obj := utils.PercentReqeust{}
        errA := c.ShouldBind(&obj);
        // Body형태가 안맞으면 거부 
        if errA != nil {
            c.JSON(400, gin.H{
                "error": "bad request - invalid json format",
            })
            return
        }
        fmt.Println(obj);

        // 존재하지 않는 랜덤스트링이면 거부 
        exist := service.ExistRandomString(obj.RandomString)
        if !exist {
            c.JSON(400, gin.H{
                "error": "random string is not found",
            })
            return
        }
        service.RecreatePercentData(obj);
    })
}