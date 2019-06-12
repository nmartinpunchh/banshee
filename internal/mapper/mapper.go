package mapper

import (
	"github.com/nmartinpunchh/banshee/internal/models"
	journeypb "github.com/nmartinpunchh/banshee/pb/punchh/journey"
	mapper "github.com/nmartinpunchh/go-automapper"
	log "github.com/sirupsen/logrus"
)

var journeyMapper = mapper.Mapper{
	PanicOnIncompatibleTypes: false,
	PanicOnMissingField:      false,
}

// JourneyPbToModel maps journeypb to a model
func JourneyPbToModel(jpb *journeypb.Journey) *models.Journey {
	mJourney := &models.Journey{}
	ret := journeyMapper.Map(jpb, mJourney)
	if len(ret.Errors) > 0 {
		log.Println(ret.Errors)
		return nil
	}

	return mJourney

}
