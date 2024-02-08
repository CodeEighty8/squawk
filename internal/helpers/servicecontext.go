package helpers

import (
	"chat-server/config"
)

type DI struct {
	AppConfig    *config.AppConfig
	WebSocketAPI WebSocketAPI
}
