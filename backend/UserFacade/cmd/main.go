package main

import (
	_ "UserFacade/docs"
	"UserFacade/internal/clients/auth_service"
	"UserFacade/internal/clients/profile_service"
	"UserFacade/internal/middleware"
	"UserFacade/internal/rpc/get_all_profiles"
	"UserFacade/internal/rpc/login"
	"UserFacade/internal/rpc/logout"
	"UserFacade/internal/rpc/register"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

// Backend doc
// @title SIF Backend API
// @version 0.5
// @description This is a backend API
// @host https://localhost:8080
// @BasePath /
func main() {
	logger := logrus.New()

	logger.Info("Initializing dependent services")
	authServiceClient, err := auth_service.NewClientFromEnv()
	if err != nil {
		logger.Errorf("cannot connect to auth service: %s", err.Error())
		return
	}
	profileServiceClient, err := profile_service.NewClientFromEnv()
	if err != nil {
		logger.Errorf("cannot connect to profile service: %s", err.Error())
		return
	}
	logger.Info("Dependent services initializing finished")

	loginHandler := login.NewHandler(authServiceClient, profileServiceClient, logger)
	logoutHandler := logout.NewHandler(authServiceClient, logger)
	registerHandler := register.NewHandler(authServiceClient, profileServiceClient, logger)
	getAllProfilesHandler := get_all_profiles.NewHandler(profileServiceClient, logger)
	//handler := http.NewServeMux()
	handler := mux.NewRouter()
	handler.HandleFunc("/login", loginHandler.Handle)
	handler.HandleFunc("/register", registerHandler.Handle)
	handler.HandleFunc("/logout", logoutHandler.Handle)
	handler.HandleFunc("/profiles", getAllProfilesHandler.Handle)

	handlers := middleware.SetupMiddleware(handler)
	handler.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	err = http.ListenAndServe(":8080", handlers)
	if err != nil {
		return
	}
}
