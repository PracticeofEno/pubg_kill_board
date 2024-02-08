package utils

import (
	"fmt"
	"kill_board/db"
)

var client *db.PrismaClient

func GetClient() *db.PrismaClient {
    if client != nil {
		fmt.Println("client is not null")
        return client
    } else {
        client = db.NewClient()
        if err := client.Prisma.Connect(); err != nil {
            client = nil
            return nil
        }
        return client
    }
}