package main

import "context"

// GrpcService represents the grpc service
type GrpcService struct {
	Repository Repository.IRepository
}

// CreateWorkflow ...
func (s *GrpcService) CreateWorkflow(ctx context.Context, req *CreateWorkflowRequest) (*CreateWorkflowResponse, error) {
	workflows, err := s.Repository.GetAll()
	if err != nil {
		return nil, err
	}

	//TODO: Use automapper to map the domain back to the dto
	return

}
