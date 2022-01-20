package auth_service

import (
	"github.com/Felix1Green/SIF/backend/UserFacade/internal/generated/clients/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func NewClientFromEnv() (auth.AuthClient, error) {
	target := "auth:8887"
	client, err := grpc.Dial(target, grpc.WithConnectParams(grpc.ConnectParams{
		Backoff: backoff.Config{
			BaseDelay:  10 * time.Second,
			Multiplier: 1.5,
			Jitter:     1,
			MaxDelay:   5 * time.Minute,
		},
		MinConnectTimeout: 10 * time.Second,
	}), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	serviceClient := auth.NewAuthClient(client)
	return serviceClient, nil
}
