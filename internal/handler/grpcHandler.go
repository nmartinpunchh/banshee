package handler

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/nmartinpunchh/banshee/internal/models"
	"github.com/nmartinpunchh/banshee/internal/repository"
	workflowapipb "github.com/nmartinpunchh/banshee/pb/punchh/journeyapi"
	workflowpb "github.com/nmartinpunchh/banshee/pb/punchh/workflow"
)

// GrpcHandler represents the grpc service
type GrpcHandler struct {
	Repository repository.IRepository
}

// CreateWorkflow ...
func (s *GrpcHandler) CreateWorkflow(ctx context.Context, req *workflowapipb.CreateWorkflowRequest) (*workflowapipb.CreateWorkflowResponse, error) {

	m := jsonpb.Marshaler{}
	pstr, e := m.MarshalToString(req)
	if e != nil {
		log.Println(e)
	}
	log.Println(pstr)
	w := &models.Workflow{}

	// w := &models.Workflow{
	// 	Root: &models.Statement{
	// 		ActivityInvocation: &models.ActivityInvocation{
	// 			Arguments: []*models.Argument{&models.Argument{Argument: "kdfjdskfjdk"}},
	// 			Name:      req.Workflow.Root.ActivityInvocation.Name,
	// 			Result:    req.Workflow.Root.ActivityInvocation.Result,
	// 		},
	// 	},
	// }

	returnedW, err := s.Repository.Create(w)
	if err != nil {
		return nil, err
	}

	resp := &workflowapipb.CreateWorkflowResponse{Id: int64(returnedW.ID)}
	return resp, nil

}

// ReadWorkflow ...
func (s *GrpcHandler) ReadWorkflow(ctx context.Context, req *workflowapipb.ReadWorkflowRequest) (*workflowapipb.ReadWorkflowResponse, error) {

	workflow, err := s.Repository.GetByID(int(req.Id))
	if err != nil {
		return nil, err
	}
	log.Println(workflow)

	//TODO: Use automapper to map the domain back to the pb
	rfr := &workflowapipb.ReadWorkflowResponse{
		Workflow: &workflowpb.Workflow{
			Root: &workflowpb.Statement{
				ActivityInvocation: &workflowpb.ActivityInvocation{
					Arguments: []string{workflow.Root.ActivityInvocation.Arguments[0].Argument},
					Name:      workflow.Root.ActivityInvocation.Name,
					Result:    workflow.Root.ActivityInvocation.Result,
				},
			},
		},
	}
	return rfr, nil

}

// DeleteWorkflow ...
func (s *GrpcHandler) DeleteWorkflow(ctx context.Context, req *workflowapipb.DeleteWorkflowRequest) (*workflowapipb.DeleteWorkflowResponse, error) {
	reqID := int(req.Id)
	respID, err := s.Repository.Delete(reqID)
	if err != nil {
		return nil, err
	}

	dwr := &workflowapipb.DeleteWorkflowResponse{
		Id: int64(respID),
	}
	return dwr, nil

}

// CreateJourney ..
func (s *GrpcHandler) CreateJourney(ctx context.Context, req *workflowapipb.CreateJourneyRequest) (*workflowapipb.CreateWorkflowResponse, error) {
	return nil, errors.New("Not implemented")
}

// ReadJourney ..
func (s *GrpcHandler) ReadJourney(ctx context.Context, req *workflowapipb.ReadJourneyRequest) (*workflowapipb.ReadJourneyResponse, error) {
	return nil, errors.New("Not implemented")
}

// DeleteJourney ..
func (s *GrpcHandler) DeleteJourney(ctx context.Context, req *workflowapipb.DeleteJourneyRequest) (*workflowapipb.DeleteJourneyResponse, error) {

	return nil, errors.New("Not implemented")
}
