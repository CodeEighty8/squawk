package helpers

import (
	"log"
	"squawk/internal/models"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type connectionPool struct {
	connections map[uuid.UUID]*websocket.Conn
	broadcast   chan models.Message
	register    chan models.Register
	unregister  chan models.Register
}

func newConnectionPool() *connectionPool {
	return &connectionPool{
		connections: map[uuid.UUID]*websocket.Conn{},
		broadcast:   make(chan models.Message),
		register:    make(chan models.Register),
		unregister:  make(chan models.Register),
	}
}

func (pool *connectionPool) Start() {

	log.Println("starting connection pool")

	for {
		select {

		case regMsg := <-pool.register:
			log.Printf("registering new connection %s", regMsg.ID)
			pool.connections[regMsg.ID] = regMsg.WSConn

		case regMsg := <-pool.unregister:
			log.Printf("unregistering connection %s", regMsg.ID)
			if conn, ok := pool.connections[regMsg.ID]; ok {
				delete(pool.connections, regMsg.ID)
				conn.Close()
			}

		case msg := <-pool.broadcast:
			log.Printf("broadcasting message: %s", msg)
			for id, conn := range pool.connections {
				err := conn.WriteJSON(msg)
				if err != nil {
					log.Printf("error writing json to connection %s: %v", id, err)
					delete(pool.connections, id)
					conn.Close()
				}
			}
		}
	}
}
