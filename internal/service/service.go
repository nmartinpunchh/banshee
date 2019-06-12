package service

import (
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/nmartinpunchh/banshee/configs"
	"github.com/nmartinpunchh/banshee/internal/handler"
	"github.com/nmartinpunchh/banshee/internal/repository"
	workflowapipb "github.com/nmartinpunchh/banshee/pb/punchh/journeyapi"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Run ..
func Run() error {
	logger := log.New()
	logger.SetLevel(log.DebugLevel)
	logEntry := log.NewEntry(logger)
	env := configs.Load()
	repo := repository.Init(env)
	grpcHandler := &handler.GrpcHandler{
		Repository: repo,
	}

	lis, err := net.Listen("tcp", env.Address)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Starting gRPC service on %s", env.Address)

	s := grpc.NewServer(
		// https://github.com/grpc-ecosystem/go-grpc-middleware
		// for existing middleware
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(),
			grpc_logrus.UnaryServerInterceptor(logEntry),
		),
	)
	workflowapipb.RegisterJourneyAPIServer(s, grpcHandler)
	return s.Serve(lis)

}
