package main

import (
	"fmt"
	"github.com/Felix1Green/SIF/backend/AuthService/internal/components/token_storage"
	"github.com/Felix1Green/SIF/backend/AuthService/internal/components/user_storage"
	"github.com/Felix1Green/SIF/backend/AuthService/internal/interactor"
	"github.com/Felix1Green/SIF/backend/AuthService/internal/rpc"
	"github.com/gomodule/redigo/redis"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"time"

	"github.com/Felix1Green/SIF/backend/AuthService/internal/generated/service"
)

func main() {
	logger := logrus.New()
	port := 8887
	pool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		MaxConnections: 10,
		AcquireTimeout: time.Minute,
		ConnConfig: pgx.ConnConfig{
			Host:     "postgresql",
			Port:     5432,
			Database: "auth_service",
			User:     "auth",
			Password: "auth_pass",
		},
	})
	if err != nil {
		logger.Fatalf("cannot connect to database: %s", err.Error())
	}

	defer func() {
		if pool != nil {
			pool.Close()
		}
	}()

	redisPool := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", "redis:6379") },
	}

	postgresUserStorage := user_storage.NewPostgresUserStorage(pool)
	cacheStorage := user_storage.NewRedisUserStorage(redisPool, 10, 10)
	tokenStorage := token_storage.NewStorage(redisPool, 10, 10)

	authInteractor := interactor.NewInteractor(tokenStorage, postgresUserStorage, cacheStorage, logger, 32, "some_salt")

	serv := grpc.NewServer()
	grpcHandler := rpc.New(authInteractor)
	service.RegisterAuthServer(serv, grpcHandler)

	logger.Infof("starting auth service listening on port: %d", port)
	lis, err := net.Listen("tcp", fmt.Sprintf("auth:%d", port))
	if err != nil {
		logger.Fatalf("err during listening to port: %d", port)
	}

	err = serv.Serve(lis)
	if err != nil {
		logger.Fatalf("error in grpc network: %s", err.Error())
	}
}
