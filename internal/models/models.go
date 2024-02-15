package models

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type WSConnection struct {
	ConnID     uuid.UUID
	UserName   string
	Connection *websocket.Conn
}

type Message struct {
	ConnID  uuid.UUID `json:"connID"`
	Sender  string    `json:"sender"`
	Content string    `json:"content"`
}
