package models

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Register struct {
	ID     uuid.UUID
	WSConn *websocket.Conn
}

type Message struct {
	ConnID  uuid.UUID `json:"connID"`
	Sender  string    `json:"sender"`
	Content string    `json:"content"`
}
