package user_storage

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/gomodule/redigo/redis"

	"github.com/Felix1Green/SIF/backend/AuthService/internal"
	"github.com/Felix1Green/SIF/backend/AuthService/internal/entities"
)

type RedisCacheUserStorage struct {
	pool            *redis.Pool
	backoffMaxValue time.Duration
	backoffMaxTries int64
}

var (
	component              = "user_storage"
	UserIdNotProvidedError = fmt.Errorf("user id not provided error")
)

type specifications struct {
	BackoffMaxValue int64 `split_words:"true"`
	BackoffMaxTries int64 `split_words:"true"`
}

func NewRedisUserStorage(pool *redis.Pool, backoffMaxValue int64, backoffMaxTries int64) *RedisCacheUserStorage {
	backoffDurationMaxValue := time.Duration(backoffMaxValue) * time.Second
	return &RedisCacheUserStorage{
		pool:            pool,
		backoffMaxValue: backoffDurationMaxValue,
		backoffMaxTries: backoffMaxTries,
	}
}

func NewRedisUserStorageFromEnv(pool *redis.Pool) (*RedisCacheUserStorage, error) {
	options := &specifications{}
	err := internal.EnvOptions("user_storage", options)

	if err != nil {
		return nil, err
	}

	return NewRedisUserStorage(pool, options.BackoffMaxValue, options.BackoffMaxTries), nil
}

func (s *RedisCacheUserStorage) GetUser(username, password string) (*entities.User, error) {
	conn := s.pool.Get()
	defer func() {
		_ = conn.Close()
	}()
	k := s.createKey(username, password)
	value, err := redis.Int64(s.backoffDo(conn, "GET", k))
	if err != nil {
		return nil, internal.UserNotFoundError
	}
	return &entities.User{
		UserID:   &value,
		Username: &username,
		Password: &password,
	}, nil
}

func (s *RedisCacheUserStorage) CreateUser(username, password string, userID *int64) (*entities.User, error) {
	if userID == nil {
		return nil, UserIdNotProvidedError
	}
	conn := s.pool.Get()
	defer func() {
		_ = conn.Close()
	}()
	k := s.createKey(username, password)
	_, err := s.backoffDo(conn, "SET", k, strconv.Itoa(int(*userID)))
	if err != nil {
		return nil, internal.InternalServiceError
	}

	return &entities.User{
		UserID:   userID,
		Username: &username,
		Password: &password,
	}, nil
}

func (s *RedisCacheUserStorage) createKey(username, password string) string {
	return fmt.Sprintf("%s.%s.%s", component, username, password)
}

func (s *RedisCacheUserStorage) backoffDo(conn redis.Conn, commandName string, args ...interface{}) (reply interface{}, err error) {
	backoffCfg := backoff.NewExponentialBackOff()
	backoffCfg.MaxInterval = s.backoffMaxValue
	retryCount := int64(0)

	_ = backoff.Retry(func() error {
		if retryCount > s.backoffMaxTries {
			return nil
		}

		reply, err = conn.Do(commandName, args...)
		retryCount++

		return err
	}, backoffCfg)

	return reply, err
}
