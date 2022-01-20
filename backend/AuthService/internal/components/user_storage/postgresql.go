package user_storage

import (
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

func (s *PostgresUserStorage) GetUser(username, password string) (*entities.User, error) {
	if s.pool == nil {
		return nil, internal.InternalServiceError
	}

	query := "SELECT user_id, username, password from User where username = $1 and password = $2"
	user := &entities.User{}
	result := s.pool.QueryRow(query, username, password)
	if result == nil {
		return nil, internal.UserNotFoundError
	}
	err := result.Scan(&user.UserID, &user.Username, &user.Password)
	if err != nil {
		return nil, internal.UserNotFoundError
	}

	return user, nil
}

func (s *PostgresUserStorage) CreateUser(username, password string, userID *int64) (*entities.User, error) {
	if s.pool == nil {
		return nil, internal.InternalServiceError
	}
	user := &entities.User{}

	query := "INSERT INTO User (username, password) VALUES ($1, $2) RETURNING user_id"
	result := s.pool.QueryRow(query, username, password)
	err := result.Scan(&user.UserID)
	if err != nil {
		return nil, internal.UserAlreadyRegistered
	}

	user.Username = &username
	user.Password = &password
	return user, nil
}
