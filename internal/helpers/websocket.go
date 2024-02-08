package helpers

import (
	"log"
	"net/http"
	"squawk/internal/models"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type WebSocketAPI interface {
	HandleNewConnection(w http.ResponseWriter, r *http.Request) error
	RegisterWS(*websocket.Conn)
}

type WebSocketImpl struct {
	serviceContext *DI
	pool           *connectionPool
}

func NewWSConnection(serviceContext *DI) WebSocketAPI {
	pool := newConnectionPool()
	go pool.Start()
	return &WebSocketImpl{
		serviceContext: serviceContext,
		pool:           pool,
	}
}

func (ws *WebSocketImpl) HandleNewConnection(w http.ResponseWriter, r *http.Request) error {
	log.Printf("creating new connection")
	wsConn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error upgrading connection: %v", err)
		return err
	}

	ws.RegisterWS(wsConn)
	return nil
}

func (ws *WebSocketImpl) RegisterWS(conn *websocket.Conn) {

	regMsg := models.Register{
		ID:     uuid.New(),
		WSConn: conn,
	}

	ws.pool.register <- regMsg
	go ws.ListenForChatMessages(regMsg.ID, conn)
}

func (ws *WebSocketImpl) ListenForChatMessages(id uuid.UUID, conn *websocket.Conn) {
	// Close connection once the Liste For Chat Messages exits.
	// This method is involked in its own long running go routine so the connection should close only if the go routien exits.
	defer func() {
		log.Printf("cleaning up connection %s", id)
		if err := recover(); err != nil {
			log.Printf("Panic occurred: %v", err)
		}
		ws.pool.unregister <- models.Register{
			ID:     id,
			WSConn: conn,
		}
		conn.Close()
	}()

	log.Printf("Listening for chat messages on connection %s", id)

	for {
		var msg models.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("error reading JSON: %v", err)
			break
		}

		msg.ConnID = id
		ws.pool.broadcast <- msg
	}
}
