package service

import (
	"log"
	"net"

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

	s := grpc.NewServer()
	workflowapipb.RegisterWorkflowAPIServer(s, grpcHandler)
	return s.Serve(lis)

}
