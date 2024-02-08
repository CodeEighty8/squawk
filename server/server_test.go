package server

import (
	"chat-server/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateNewServer(t *testing.T) {
	server := CreateNewServer(&config.ServerConfig{Env: "test", Port: 9000})

	assert.Equal(t, ":9000", server.server.Addr)
}
