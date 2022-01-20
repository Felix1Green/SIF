package internal

import (
	"context"
	"fmt"
	"github.com/Felix1Green/SIF/backend/AuthService/internal/entities"
)

type Interactor interface {
	Auth(user *entities.User) (*entities.User, error)
	Logout(token string) error
	Register(user *entities.User) (*entities.User, error)
}

var(
	NoAuthenticationDataProvidedError = fmt.Errorf("no authentication data provided")
	InternalServiceError              = fmt.Errorf("internal service error: some dependencies are not available")
	UserNotFoundError    = fmt.Errorf("user not found")
	UserAlreadyRegistered = fmt.Errorf("user already registered")
)

// Logger определяет интерфейс логгера, используемый пакетом internal.
type Logger interface {
	Info(ctx context.Context, args ...interface{})
	Debug(ctx context.Context, args ...interface{})
	Error(ctx context.Context, args ...interface{})
	Warning(ctx context.Context, args ...interface{})
}
