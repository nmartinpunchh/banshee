package handler

import (
	"context"
	"log"

	"github.com/gogo/protobuf/jsonpb"
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

	m := jsonpb.Marshaler{}
	pstr, e := m.MarshalToString(req)
	if e != nil {
		log.Println(e)
	}
	log.Println(pstr)

	// w := &models.Workflow{
	// 	Root: &models.Statement{
	// 		ActivityInvocation: &models.ActivityInvocation{
	// 			Arguments: []*models.Argument{&models.Argument{Argument: "kdfjdskfjdk"}},
	// 			Name:      req.Workflow.Root.ActivityInvocation.Name,
	// 			Result:    req.Workflow.Root.ActivityInvocation.Result,
	// 		},
	// 	},
	// }
	// _ = w

	//
	//
	// ----
	//
	//

	i := int64(8)
	resp := &workflowapipb.CreateWorkflowResponse{Id: i}
	return resp, nil

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
					// Name:   workflows[0].Root.Activity.Name,
					// Result: workflows[0].Root.Activity.Result,
					Arguments: []string{"aa", "bb"},
					Name:      "testname",
					Result:    "testResult",
				},
			},
		},
	}
	return rfr, nil

}
