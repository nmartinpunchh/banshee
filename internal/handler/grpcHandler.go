package handler

import (
	"context"

	"github.com/nmartinpunchh/banshee/internal/models"
	"github.com/nmartinpunchh/banshee/internal/repository"
	workflowpb "github.com/nmartinpunchh/banshee/pb/punchh/workflow"
	workflowapipb "github.com/nmartinpunchh/banshee/pb/punchh/workflowapi"
)

// GrpcHandler represents the grpc service
type GrpcHandler struct {
	Repository repository.IRepository
}

// CreateWorkflow ...
func (s *GrpcHandler) CreateWorkflow(ctx context.Context, req *workflowapipb.CreateWorkflowRequest) (*workflowapipb.CreateWorkflowResponse, error) {
	//TODO: Use automapper to map the pb to model

	w := &models.Workflow{
		Root: models.Statement{
			Activity: &models.ActivityInvocation{
				Name:      req.Workflow.Root.ActivityInvocation.Name,
				Arguments: []*models.Argument{},
				// Arguments: []*models.Argument{
				// 	Arguments: req.Workflow.Root.ActivityInvocation.Arguments[0],
				// },
				Result: req.Workflow.Root.ActivityInvocation.Result,
			},
			Sequence: &models.Sequence{},
			Parallel: &models.Parallel{},
		},
	}

	workflow, err := s.Repository.Create(w)
	//TODO temp
	_ = workflow
	if err != nil {
		return nil, err
	}

	// temp
	//TODO: Use automapper to map the model to pb
	wfr := &workflowapipb.CreateWorkflowResponse{}
	return wfr, nil

}

// ReadWorkflow ...
func (s *GrpcHandler) ReadWorkflow(ctx context.Context, req *workflowapipb.ReadWorkflowRequest) (*workflowapipb.ReadWorkflowResponse, error) {

	workflows, err := s.Repository.GetAll()
	if err != nil {
		return nil, err
	}

	// temp
	_ = workflows

	//TODO: Use automapper to map the domain back to the pb
	rfr := &workflowapipb.ReadWorkflowResponse{
		Workflow: &workflowpb.Workflow{
			Root: &workflowpb.Statement{
				ActivityInvocation: &workflowpb.ActivityInvocation{
					Name:   workflows[0].Root.Activity.Name,
					Result: workflows[0].Root.Activity.Result,
				},
			},
		},
	}
	return rfr, nil

}
