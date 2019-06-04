package service

import (
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/nmartinpunchh/banshee/configs"
	"github.com/nmartinpunchh/banshee/internal/handler"
	"github.com/nmartinpunchh/banshee/internal/repository"
	workflowapipb "github.com/nmartinpunchh/banshee/pb/punchh/workflowapi"
	"google.golang.org/grpc"
)

// Run ..
func Run() error {
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
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(),
		),
	)
	workflowapipb.RegisterWorkflowAPIServer(s, grpcHandler)
	return s.Serve(lis)

}
