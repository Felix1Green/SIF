package profile_service

import (
	"UserFacade/internal/generated/clients/profile"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func NewClientFromEnv() (profile.ProfileClient, error) {
	target := "profile:8889"
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

	serviceClient := profile.NewProfileClient(client)
	return serviceClient, nil
}
