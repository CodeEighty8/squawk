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
	Sender   uuid.UUID `json:"sender"`
	Receiver uuid.UUID `json:"receiver"`
	Content  string    `json:"content"`
}
