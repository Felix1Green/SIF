package main

import (
	"ProfileService/internal/components/profile_storage"
	profileGRPC "ProfileService/internal/generated/service/profile"
	"ProfileService/internal/interactor/profile"
	"ProfileService/internal/rpc"
	"fmt"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"time"
)

func main() {
	logger := logrus.New()
	port := 8889
	pool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		MaxConnections: 10,
		AcquireTimeout: time.Minute,
		ConnConfig: pgx.ConnConfig{
			Host:     "postgresql",
			Port:     5432,
			Database: "profile_service",
			User:     "profile",
			Password: "profile_pass",
		},
	})
	defer func() {
		if pool != nil {
			pool.Close()
		}
	}()

	if err != nil {
		logger.Fatalf("cannot connect to database: %s", err.Error())
	}

	profilesStorage := profile_storage.NewPostgresProfileStorage(pool, logger)

	profileInteractor := profile.NewInteractor(profilesStorage, logger)

	serv := grpc.NewServer()
	grpcHandler := rpc.NewHandler(profileInteractor, logger)
	profileGRPC.RegisterProfileServer(serv, grpcHandler)

	logger.Infof("starting profile service listening on port: %d", port)
	lis, err := net.Listen("tcp", fmt.Sprintf("profile:%d", port))
	if err != nil {
		logger.Fatalf("err during listening to port: %d", port)
	}

	err = serv.Serve(lis)
	if err != nil {
		logger.Fatalf("error in grpc network: %s", err.Error())
	}
}
