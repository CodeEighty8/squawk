package controller

import (
	"net/http"
	"squawk/internal/service"
	"squawk/server"

	"github.com/go-chi/chi/v5"
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
	server.Mux.Route("/api/chat/v1", func(r chi.Router) {
		r.Get("/ws", c.Connect)
	})
}
