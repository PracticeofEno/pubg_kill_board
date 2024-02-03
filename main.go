package main

// import (
// 	"fmt"
// 	"kill_board/pubg_api"
// )
import (
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

// func main() {
// 	fmt.Println("Hello, World!");
// 	apiService := pubg_api.CreateAPIService("eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiI4ZjY4NjcxMC04ZjIwLTAxM2MtYmQ4Yy0yYWNjZjk1NjI0ZjIiLCJpc3MiOiJnYW1lbG9ja2VyIiwiaWF0IjoxNzA0NTg2ODYxLCJwdWIiOiJibHVlaG9sZSIsInRpdGxlIjoicHViZyIsImFwcCI6IjkxZWU0Y2M4LWY2MTEtNDUxOS05NjQ1LTEzNWJlN2Y4NjkyMiJ9.tadtjH48XZwaRKRSh-ROjHFHuO2dpjrzYASkixz-px0");
// 	accountId, _ := apiService.GetAccountId("PracticeofEno2")
// 	fmt.Println(accountId);
// 	// r := gin.Default();
// 	// apiGroup := r.Group("/api");
// 	// router.SetupAPIRoutes(apiGroup);
// 	// r.Run("0.0.0.0:80") // listen and serve on 0.0.0.0:8080;
// }

func main() {
	dsn := "postgres://postgres:dkflfkd@localhost:5432/pubg?sslmode=disable"
	// dsn := "unix://user:pass@dbname/var/run/postgresql/.s.PGSQL.5432"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())
	err := db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}

	fmt.Println("Connected to the database!")
}