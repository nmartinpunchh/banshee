package handler

import (
	"context"

	workflowpb "github.com/nmartinpunchh/banshee/pb/punchh/workflow"
	workflowapipb "github.com/nmartinpunchh/banshee/pb/punchh/workflowapi"
)

// GrpcHandler represents the grpc service
type GrpcHandler struct {
	Repository Repository.IRepository
}

// CreateWorkflow ...
func (s *GrpcHandler) CreateWorkflow(ctx context.Context, req *workflowapipb.CreateWorkflowRequest) (*workflowapipb.CreateWorkflowResponse, error) {
	workflows, err := s.Repository.GetAll()
	if err != nil {
		return nil, err
	}

	//TODO: Use automapper to map the domain back to the dto
	wfr := &workflowapipb.CreateWorkflowResponse{
		Workflow: &workflowpb.Workflow{
			Root: &workflowpb.Statement{},
		},
	}
	return wfr, nil

}
