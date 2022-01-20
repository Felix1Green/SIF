package rpc

import (
	"context"

	"github.com/Felix1Green/SIF/backend/AuthService/internal"
	"github.com/Felix1Green/SIF/backend/AuthService/internal/entities"
	"github.com/Felix1Green/SIF/backend/AuthService/internal/generated/service"
)

type handler struct {
	authInteractor internal.Interactor
	service.UnimplementedAuthServer
}

func New(authInteractor internal.Interactor) *handler {
	return &handler{
		authInteractor: authInteractor,
	}
}

func (s *handler) Auth(ctx context.Context, in *service.AuthIn) (*service.AuthOut, error) {
	var (
		userId    *int64  = nil
		userToken *string = nil
		success           = true
		outErr    service.Errors
	)

	user, err := s.authInteractor.Auth(&entities.User{
		Username:  in.Username,
		Password:  in.Password,
		AuthToken: in.AuthToken,
	})
	if err != nil {
		success = false
		switch err{
		case internal.UserNotFoundError:
			outErr = service.Errors_IncorrectUser
		case internal.NoAuthenticationDataProvidedError:
			outErr = service.Errors_NoAuthDataProvided
		default:
			outErr = service.Errors_InternalServiceError
		}

		return &service.AuthOut{
			Success: success,
			Error: &outErr,
		}, nil
	} else if user != nil {
		userId = user.UserID
		userToken = user.AuthToken
	}
	return &service.AuthOut{
		UserId:    userId,
		Success:   success,
		Error:     &outErr,
		UserToken: userToken,
	}, nil
}

func (s *handler) LogOut(ctx context.Context, in *service.LogoutIn) (*service.LogoutOut, error) {
	if in.AuthToken == ""{
		outErr := service.Errors_NoAuthDataProvided
		return &service.LogoutOut{
			Success: false,
			Error: &outErr,
		},nil
	}

	err := s.authInteractor.Logout(in.AuthToken)
	if err != nil{
		outErr := service.Errors_InternalServiceError
		return &service.LogoutOut{
			Success: false,
			Error: &outErr,
		}, nil
	}

	return &service.LogoutOut{
		Success: true,
	}, nil
}

func (s *handler) Register(ctx context.Context, in *service.RegisterIn) (*service.RegisterOut, error) {
	var (
		userId    *int64  = nil
		userToken *string = nil
		success           = true
		outErr    service.Errors
	)

	user, err := s.authInteractor.Register(&entities.User{
		Username:  &in.UserName,
		Password:  &in.Password,
	})
	if err != nil {
		success = false
		switch err{
		case internal.UserAlreadyRegistered:
			outErr = service.Errors_IncorrectUser
		case internal.NoAuthenticationDataProvidedError:
			outErr = service.Errors_NoAuthDataProvided
		default:
			outErr = service.Errors_InternalServiceError
		}

		return &service.RegisterOut{
			Success: success,
			Error: &outErr,
		}, nil
	} else if user != nil {
		userId = user.UserID
		userToken = user.AuthToken
	}

	return &service.RegisterOut{
		UserId:    userId,
		Success:   success,
		Error:     &outErr,
		UserToken: userToken,
	}, nil
}
