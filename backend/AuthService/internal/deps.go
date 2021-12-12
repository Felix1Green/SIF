package internal

import "github.com/Felix1Green/SIF/backend/AuthService/internal/entities"

type Interactor interface {
	Auth(user *entities.User) (*int64, error)
}
