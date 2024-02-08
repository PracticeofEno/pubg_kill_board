package main

import (
	"kill_board/internal/app/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default();
    usersGroup := r.Group("/users");
    handlers.SetupUsersRoute(usersGroup);
    r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080;
}

// func main() {
// 	client := utils.GetClient()
// 	ctx := context.Background()
// 	client.User.FindUnique(db.User.ID.Equals(1)).Update(
// 	)
// }