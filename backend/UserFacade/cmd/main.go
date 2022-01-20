package main

import (
	"github.com/Felix1Green/SIF/backend/UserFacade/internal/clients/auth_service"
	"github.com/Felix1Green/SIF/backend/UserFacade/internal/middleware"
	"github.com/Felix1Green/SIF/backend/UserFacade/internal/rpc/login"
	"github.com/Felix1Green/SIF/backend/UserFacade/internal/rpc/logout"
	"github.com/Felix1Green/SIF/backend/UserFacade/internal/rpc/register"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	logger := logrus.New()

	logger.Info("Initializing dependent services")
	authServiceClient, err := auth_service.NewClientFromEnv()
	if err != nil {
		logger.Errorf("cannot connect to auth service: %s", err.Error())
		return
	}
	logger.Info("Dependent services initializing finished")

	loginHandler := login.NewLoginHandler(authServiceClient, logger)
	logoutHandler := logout.NewHandler(authServiceClient, logger)
	registerHandler := register.NewHandler(authServiceClient, logger)

	handler := http.NewServeMux()
	handler.HandleFunc("/login", loginHandler.Handle)
	handler.HandleFunc("/register", registerHandler.Handle)
	handler.HandleFunc("/logout", logoutHandler.Handle)

	handlers := middleware.SetupMiddleware(handler)
	err = http.ListenAndServe(":8080", handlers)
	if err != nil {
		return
	}
}
