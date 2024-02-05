package utils

import "kill_board/db"

var client *db.PrismaClient

func GetClient() *db.PrismaClient {
	if client != nil {
		return client
	} else {
		return db.NewClient()
	}
}