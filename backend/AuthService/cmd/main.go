package main

import (
	"google.golang.org/grpc"

	"github.com/Felix1Green/SIF/backend/AuthService/internal/generated/service"
	"github.com/Felix1Green/SIF/backend/AuthService/internal/rpc/auth"
)

func main() {
	serv := grpc.NewServer()
	handler := auth.New()
	service.RegisterAuthServer(serv, handler)
}
