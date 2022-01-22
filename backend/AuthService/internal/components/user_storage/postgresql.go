package user_storage

import (
	"fmt"
	"github.com/Felix1Green/SIF/backend/AuthService/internal"

	"github.com/jackc/pgx"

	"github.com/Felix1Green/SIF/backend/AuthService/internal/entities"
)

type PostgresUserStorage struct {
	pool *pgx.ConnPool
}

func NewPostgresUserStorage(pool *pgx.ConnPool) *PostgresUserStorage {
	return &PostgresUserStorage{
		pool: pool,
	}
}

func (s *PostgresUserStorage) GetUser(username string) (*entities.User, error) {
	if s.pool == nil {
		return nil, internal.InternalServiceError
	}

	query := "SELECT user_id, username, password from Users where username = $1"
	user := &entities.User{}
	result := s.pool.QueryRow(query, username)
	if result == nil {
		fmt.Println("result is nil")
		return nil, internal.UserNotFoundError
	}
	err := result.Scan(&user.UserID, &user.Username, &user.Password)
	if err != nil {
		fmt.Println(err.Error())
		return nil, internal.UserNotFoundError
	}

	return user, nil
}

func (s *PostgresUserStorage) CreateUser(username, password string, userID *int64) (*entities.User, error) {
	if s.pool == nil {
		return nil, internal.InternalServiceError
	}
	user := &entities.User{}

	query := "INSERT INTO Users (username, password) VALUES ($1, $2) RETURNING user_id"
	result := s.pool.QueryRow(query, username, password)
	err := result.Scan(&user.UserID)
	if err != nil {
		return nil, internal.UserAlreadyRegistered
	}

	user.Username = &username
	user.Password = &password
	return user, nil
}
