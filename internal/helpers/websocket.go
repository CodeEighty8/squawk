package helpers

import (
	"log"
	"net/http"
	"squawk/internal/errors"
	"squawk/internal/models"

	"github.com/google/uuid"
)

type WebSocketAPI interface {
	HandleNewConnection(w http.ResponseWriter, r *http.Request) error
	// RegisterWS()
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
	connId, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		log.Printf("error parsing uuid: %v", err)
		return errors.NewUUIDParseError(err)
	}

	log.Printf("creating new connection for id %s", connId)
	wsConn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error upgrading connection: %v", err)
		return errors.NewWSUpgradeError(err)
	}

	conn := &models.WSConnection{
		ConnID:     connId,
		Connection: wsConn,
	}

	ws.pool.saveConnection <- conn
	go ws.ListenForChatMessages(connId, conn)
	return nil
}

func (ws *WebSocketImpl) ListenForChatMessages(id uuid.UUID, conn *models.WSConnection) {
	// Close connection once the Liste For Chat Messages exits.
	// This method is involked in its own long running go routine so the connection should close only if the go routien exits.
	defer func() {
		log.Printf("cleaning up connection %s", id)
		if err := recover(); err != nil {
			log.Printf("Panic occurred: %v", err)
		}
		ws.pool.unregister <- id
	}()

	log.Printf("Listening for chat messages on connection %s", id)

	for {
		var msg models.Message
		err := conn.Connection.ReadJSON(&msg)
		if err != nil {
			log.Printf("error reading JSON: %v", err)
			break
		}

		msg.ConnID = id
		ws.pool.broadcast <- msg
	}
}
