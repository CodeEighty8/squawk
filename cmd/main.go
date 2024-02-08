package main

import (
	"log"
	"squawk/config"
	"squawk/internal/controller"
	"squawk/internal/helpers"
	"squawk/internal/service"
	"squawk/server"
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
