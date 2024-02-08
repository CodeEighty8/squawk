package main

import (
	"chat-server/config"
	"chat-server/internal/controller"
	"chat-server/internal/helpers"
	"chat-server/internal/service"
	"chat-server/server"
	"log"
)

func main() {

	appConfig := &config.AppConfig{}
	config, err := config.GetConfig(appConfig)
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	serviceContext := &helpers.DI{
		AppConfig: appConfig,
	}

	wsConnection := helpers.NewWSConnection(serviceContext)
	serviceContext.WebSocketAPI = wsConnection

	chatService := service.NewChatSercice(serviceContext)
	chatController := controller.NewController(chatService)

	server := server.CreateNewServer(&config.Server)
	chatController.SetupRoutes(server)

	server.Start()
}
