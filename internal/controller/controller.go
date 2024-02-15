package controller

import (
	"net/http"
	"squawk/internal/service"
	"squawk/server"
)

type Controller struct {
	chatService service.IChatService
}

func NewController(chatService service.IChatService) *Controller {
	return &Controller{
		chatService: chatService,
	}
}

func (c *Controller) Connect(w http.ResponseWriter, r *http.Request) {
	c.chatService.Connect(w, r)
}

func (c *Controller) SetupRoutes(server *server.Server) {
	server.Mux.HandleFunc("GET /api/chat/v1/ws", c.chatService.Connect)
}
