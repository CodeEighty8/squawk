package helpers

import (
	"squawk/config"
)

type DI struct {
	AppConfig    *config.AppConfig
	WebSocketAPI WebSocketAPI
}
