package mapper

import (
	"reflect"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
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
	journeyMapper.CustomMappers = []mapper.CustomFieldMapper{pbstatus2mStatus}
	journeyMapper.FieldNameMaps = map[string]string{
		"SegmentID": "SegmentId",
		"SegmentId": "SegmentID",
	}

	journeyMapper.IgnoreDestFields = []string{"Status", "Model", "SegmentID", "Workflow"}
	mJourney := &models.Journey{}
	ret := journeyMapper.Map(jpb, mJourney)
	if len(ret.Errors) > 0 {
		log.Println(ret.Errors)
		return nil
	}

	return mJourney

}

func pbstatus2mStatus(sourceVal reflect.Value, sourceType reflect.Type, destVal reflect.Value, destType reflect.Type) (handled bool) {
	if destType.Kind() == reflect.Struct {
		if sourceVal.Type().Name() == "Timestamp" && destVal.Type().Name() == "Time" {
			if sourceType.Kind() == reflect.Struct {
				convertToIface := sourceVal.Interface().(timestamp.Timestamp)
				val, err := ptypes.Timestamp(&convertToIface)
				if err != nil {
					return false

				}
				destVal.Set(reflect.ValueOf(val))
				return true
			}
		}
	}
	return false
}
