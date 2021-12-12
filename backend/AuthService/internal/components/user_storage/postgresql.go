package user_storage

import (
	"fmt"

	"github.com/jackc/pgx"

	"github.com/Felix1Green/SIF/backend/AuthService/internal/entities"
)

type PostgresUserStorage struct {
	pool *pgx.ConnPool
}

var (
	InternalServiceError = fmt.Errorf("internal service error: some dependencies are not available")
	UserNotFoundError    = fmt.Errorf("user not found")
)

func NewPostgresUserStorage(pool *pgx.ConnPool) *PostgresUserStorage {
	return &PostgresUserStorage{
		pool: pool,
	}
}

func (s *PostgresUserStorage) GetUser(username, password string) (*entities.User, error) {
	if s.pool == nil {
		return nil, InternalServiceError
	}

	query := "SELECT user_id, username, password from User where username = $1 and password = $2"
	user := &entities.User{}
	result := s.pool.QueryRow(query, username, password)
	if result == nil {
		return nil, UserNotFoundError
	}
	err := result.Scan(&user.UserID, &user.Username, &user.Password)
	if err != nil {
		return nil, UserNotFoundError
	}

	return user, nil
}

func (s *PostgresUserStorage) CreateUser(username, password string, userID *int64) (*entities.User, error) {
	if s.pool == nil {
		return nil, InternalServiceError
	}
	user := &entities.User{}

	query := "INSERT INTO User (username, password) VALUES ($1, $2) RETURNING user_id"
	result := s.pool.QueryRow(query, username, password)
	err := result.Scan(&user.UserID)
	if err != nil {
		return nil, err
	}

	user.Username = &username
	user.Password = &password
	return user, nil
}
