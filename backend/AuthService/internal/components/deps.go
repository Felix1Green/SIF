package components

import "github.com/Felix1Green/SIF/backend/AuthService/internal/entities"

type TokenStorage interface {
	Set(token string, value int64) error
	Get(token string) (int64, error)
	Del(token string) error
}

type UserStorage interface {
	GetUser(username, password string) (*entities.User, error)
	CreateUser(username, password string, userID *int64) (*entities.User, error)
}
