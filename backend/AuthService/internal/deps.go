package internal

import (
	"context"
	"github.com/Felix1Green/SIF/backend/AuthService/internal/entities"
)

type Interactor interface {
	Auth(user *entities.User) (*entities.User, error)
}

// Logger определяет интерфейс логгера, используемый пакетом internal.
type Logger interface {
	Info(ctx context.Context, args ...interface{})
	Debug(ctx context.Context, args ...interface{})
	Error(ctx context.Context, args ...interface{})
	Warning(ctx context.Context, args ...interface{})
}
