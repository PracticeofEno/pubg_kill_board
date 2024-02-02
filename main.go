package main

import (
	"fmt"
	"kill_board/pubg_api"
	"kill_board/router"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!");
	pubg_api.CreateAPIService("test").Tmp();
	r := gin.Default();
	apiGroup := r.Group("/api");
	router.SetupAPIRoutes(apiGroup);
	r.Run("0.0.0.0:80") // listen and serve on 0.0.0.0:8080;
}

// func main() {
// 	var apiService = pubg_api.CreateAPIService("eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiI4ZjY4NjcxMC04ZjIwLTAxM2MtYmQ4Yy0yYWNjZjk1NjI0ZjIiLCJpc3MiOiJnYW1lbG9ja2VyIiwiaWF0IjoxNzA0NTg2ODYxLCJwdWIiOiJibHVlaG9sZSIsInRpdGxlIjoicHViZyIsImFwcCI6IjkxZWU0Y2M4LWY2MTEtNDUxOS05NjQ1LTEzNWJlN2Y4NjkyMiJ9.tadtjH48XZwaRKRSh-ROjHFHuO2dpjrzYASkixz-px0");
// 	apiService.Tmp();
// 	apiService.GetPlayer();

// 	var apiService2 = pubg_api.CreateAPIService("eyJ0aeXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiI4ZjY4NjcxMC04ZjIwLTAxM2MtYmQ4Yy0yYWNjZjk1NjI0ZjIiLCJpc3MiOiJnYW1lbG9ja2VyIiwiaWF0IjoxNzA0NTg2ODYxLCJwdWIiOiJibHVlaG9sZSIsInRpdGxlIjoicHViZyIsImFwcCI6IjkxZWU0Y2M4LWY2MTEtNDUxOS05NjQ1LTEzNWJlN2Y4NjkyMiJ9.tadtjH48XZwaRKRSh-ROjHFHuO2dpjrzYASkixz-px0");
// 	apiService2.Tmp();
// 	apiService2.GetPlayer();
// }