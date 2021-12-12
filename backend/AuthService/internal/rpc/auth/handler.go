package auth

import (
	"context"

	"github.com/Felix1Green/SIF/backend/AuthService/internal"
	"github.com/Felix1Green/SIF/backend/AuthService/internal/entities"
	"github.com/Felix1Green/SIF/backend/AuthService/internal/generated/service"
)

type Handler struct {
	authInteractor internal.Interactor
	service.UnimplementedAuthServer
}

func New() *Handler {
	return &Handler{}
}

func (s *Handler) Auth(ctx context.Context, in *service.AuthIn) (*service.AuthOut, error) {
	var (
		success         = true
		outErr  *string = nil
	)

	userID, err := s.authInteractor.Auth(&entities.User{
		Username:  in.Username,
		Password:  in.Password,
		AuthToken: in.AuthToken,
	})
	if err != nil {
		success = false
		castedErr := err.Error()
		outErr = &castedErr
	}
	return &service.AuthOut{
		UserId:  userID,
		Success: success,
		Error:   outErr,
	}, nil
}

func (s *Handler) LogOut(ctx context.Context, in *service.LogoutIn) (*service.LogoutOut, error) {
	return nil, nil
}

func (s *Handler) Register(ctx context.Context, in *service.RegisterIn) (*service.RegisterOut, error) {
	return nil, nil
}
