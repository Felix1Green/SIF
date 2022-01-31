package main

import (
	_ "ChatService/docs"
	"ChatService/internal/clients/auth_service"
	"ChatService/internal/rpc"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

// Backend chat doc
// @title SIF Backend Chat API
// @version 0.5
// @description This is a chat API
// @host https://localhost:8890
// @BasePath /
func main() {
	logger := logrus.New()

	logger.Info("initializing deps clients")
	authClient, err := auth_service.NewClientFromEnv()
	if err != nil {
		logger.Errorf("failed to initialize auth client, err: %s", err.Error())
		return
	}
	logger.Info("deps clients initializing finished with success")

	// handlers
	socketHandler := rpc.NewHandler(logger, authClient)

	// configure router
	router := mux.NewRouter()
	router.HandleFunc("/chat", socketHandler.Handle)
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	err = http.ListenAndServe(":8890", router)
	if err != nil {
		logger.Errorf("exception occured during port listening: %s", err.Error())
		return
	}
}
