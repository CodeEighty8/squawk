package helpers

import (
	"log"
	"squawk/internal/models"

	"github.com/google/uuid"
)

type connectionPool struct {
	connections map[uuid.UUID]*models.WSConnection
	broadcast   chan models.Message
	unregister  chan uuid.UUID
}

func newConnectionPool() *connectionPool {
	return &connectionPool{
		connections: map[uuid.UUID]*models.WSConnection{},
		broadcast:   make(chan models.Message),
		unregister:  make(chan uuid.UUID),
	}
}

func (pool *connectionPool) Start() {

	log.Println("starting connection pool")

	for {
		select {

		case id := <-pool.unregister:
			log.Printf("unregistering connection %s", id)
			if conn, ok := pool.connections[id]; ok {
				delete(pool.connections, id)
				conn.Connection.Close()
			} else {
				log.Printf("unable to find connection in pool for %s", id)
			}

		case msg := <-pool.broadcast:
			log.Printf("broadcasting message: %s", msg)
			for id, conn := range pool.connections {
				err := conn.Connection.WriteJSON(msg)
				if err != nil {
					log.Printf("error writing json to connection %s: %v", id, err)
					delete(pool.connections, id)
					conn.Connection.Close()
				}
			}
		}
	}
}

func (pool *connectionPool) SaveConn(conn *models.WSConnection) {
	pool.connections[conn.ConnID] = conn
}
