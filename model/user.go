package main

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:myusers,alias:u"`
	ID          int        `json:"id"`
}