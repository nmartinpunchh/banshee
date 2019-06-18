package handler

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/nmartinpunchh/banshee/internal/mapper"
	"github.com/nmartinpunchh/banshee/internal/repository"
	journeyapipb "github.com/nmartinpunchh/banshee/pb/punchh/journeyapi"
)

// GrpcHandler represents the grpc service
type GrpcHandler struct {
	Repository repository.IRepository
}

// CreateJourney ..
func (s *GrpcHandler) CreateJourney(ctx context.Context, req *journeyapipb.CreateJourneyRequest) (*journeyapipb.CreateJourneyResponse, error) {
	m := jsonpb.Marshaler{}
	pstr, e := m.MarshalToString(req)
	if e != nil {
		log.Println(e)
	}
	log.Println(pstr)
	mJourney := mapper.JourneyPbToModel(req.GetJourney())

	returnedW, err := s.Repository.Create(mJourney)
	if err != nil {
		return nil, err
	}

	resp := &journeyapipb.CreateJourneyResponse{Id: int64(returnedW.ID)}
	return resp, nil
}

// ReadJourney ..
func (s *GrpcHandler) ReadJourney(ctx context.Context, req *journeyapipb.ReadJourneyRequest) (*journeyapipb.ReadJourneyResponse, error) {
	journey, err := s.Repository.GetByID(int(req.Id))
	if err != nil {
		return nil, err
	}
	log.Println(journey)

	pbJourney := mapper.JourneyModelToPb(journey)
	rwr := &journeyapipb.ReadJourneyResponse{Journey: pbJourney}
	return rwr, nil

}

// DeleteJourney ..
func (s *GrpcHandler) DeleteJourney(ctx context.Context, req *journeyapipb.DeleteJourneyRequest) (*journeyapipb.DeleteJourneyResponse, error) {
	reqID := int(req.Id)
	respID, err := s.Repository.Delete(reqID)
	if err != nil {
		return nil, err
	}

	dwr := &journeyapipb.DeleteJourneyResponse{
		Id: int64(respID),
	}
	return dwr, nil

}
