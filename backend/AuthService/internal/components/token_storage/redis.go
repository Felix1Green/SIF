package token_storage

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/gomodule/redigo/redis"

	"github.com/Felix1Green/SIF/backend/AuthService/internal"
)

type RedisTokenStorage struct {
	pool            *redis.Pool
	backoffMaxValue time.Duration
	backoffMaxTries int64
}

var (
	component = "token_storage"
)

type specifications struct {
	BackoffMaxValue int64 `split_words:"true"`
	BackoffMaxTries int64 `split_words:"true"`
}

func NewStorage(pool *redis.Pool, backoffMaxValue int64, backoffMaxTries int64) *RedisTokenStorage {
	backoffDurationMaxValue := time.Duration(backoffMaxValue) * time.Second
	return &RedisTokenStorage{
		pool:            pool,
		backoffMaxValue: backoffDurationMaxValue,
		backoffMaxTries: backoffMaxTries,
	}
}

func NewStorageFromEnv(pool *redis.Pool) (*RedisTokenStorage, error) {
	options := &specifications{}
	err := internal.EnvOptions("token_storage", options)

	if err != nil {
		return nil, err
	}

	return NewStorage(pool, options.BackoffMaxValue, options.BackoffMaxTries), nil
}

func (s *RedisTokenStorage) Set(token string, value int64) error {
	conn := s.pool.Get()
	defer func() {
		_ = conn.Close()
	}()
	k := s.createKey(token)
	_, err := s.backoffDo(conn, "SET", k, strconv.Itoa(int(value)))
	return err
}

func (s *RedisTokenStorage) Get(token string) (int64, error) {
	conn := s.pool.Get()
	defer func() {
		_ = conn.Close()
	}()
	k := s.createKey(token)

	value, err := redis.Int(s.backoffDo(conn, "GET", k))

	return int64(value), err
}

func (s *RedisTokenStorage) Del(token string) error {
	conn := s.pool.Get()
	defer func() {
		_ = conn.Close()
	}()
	k := s.createKey(token)

	_, err := s.backoffDo(conn, "DEL", k)
	return err
}

func (s *RedisTokenStorage) createKey(key string) string {
	return fmt.Sprintf("%s.%s", component, key)
}

func (s *RedisTokenStorage) backoffDo(conn redis.Conn, commandName string, args ...interface{}) (reply interface{}, err error) {
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
