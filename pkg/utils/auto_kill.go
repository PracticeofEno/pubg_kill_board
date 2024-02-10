package utils

import (
    "fmt"
)

type Worker struct {
    ApiKey string
    Nickname string

}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[string]map[*Client]bool),
		unregister: make(chan *Client),
		register:   make(chan *Client),
		broadcast:  make(chan Message),
	}
}