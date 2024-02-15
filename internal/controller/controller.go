package controller

import (
	"net/http"
	"squawk/internal/errors"
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
	if err := c.chatService.Connect(w, r); err != nil {

		switch err.(type) {
		case *errors.UUIDParseError:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		case *errors.WSUpgradeError:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (c *Controller) SetupRoutes(server *server.Server) {
	server.Mux.HandleFunc("GET /api/chat/v1/ws/{id}", c.Connect)
}
