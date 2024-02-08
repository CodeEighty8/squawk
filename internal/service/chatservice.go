package service

import (
	"chat-server/internal/helpers"
	"net/http"
)

type IChatService interface {
	Connect(http.ResponseWriter, *http.Request)
}

type ChatSercice struct {
	serviceContext *helpers.DI
}

func NewChatSercice(serviceContext *helpers.DI) IChatService {
	return &ChatSercice{
		serviceContext: serviceContext,
	}
}

func (c *ChatSercice) Connect(w http.ResponseWriter, r *http.Request) {
	c.serviceContext.WebSocketAPI.HandleNewConnection(w, r)
}
