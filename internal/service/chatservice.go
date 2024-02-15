package service

import (
	"net/http"
	"squawk/internal/helpers"
)

type IChatService interface {
	Connect(http.ResponseWriter, *http.Request) error
}

type ChatSercice struct {
	serviceContext *helpers.DI
}

func NewChatSercice(serviceContext *helpers.DI) IChatService {
	return &ChatSercice{
		serviceContext: serviceContext,
	}
}

func (c *ChatSercice) Connect(w http.ResponseWriter, r *http.Request) error {
	if err := c.serviceContext.WebSocketAPI.HandleNewConnection(w, r); err != nil {
		return err
	}
	return nil
}
