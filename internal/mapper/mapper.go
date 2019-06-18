package mapper

import (
	"reflect"
	"time"

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
	journeyMapper.CustomMappers = []mapper.CustomFieldMapper{TimestampToTime}

	//TODO: Create custom mapper for Status
	// Must ignore the gorm foreign key as they are not part of the proto definition.
	journeyMapper.IgnoreDestFields = []string{"Status", "Model", "ArgumentID", "SegmentID", "ActivityInvocationID", "StatementID", "WorkflowID", "JourneyID"}
	mJourney := &models.Journey{}
	ret := journeyMapper.Map(jpb, mJourney)
	if len(ret.Errors) > 0 {
		log.Println(ret.Errors)
		return nil
	}

	return mJourney

}

// TimestampToTime Converts a timestamp to a time
func TimestampToTime(sourceVal reflect.Value, sourceType reflect.Type, destVal reflect.Value, destType reflect.Type) (handled bool) {
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

// JourneyModelToPb maps journeypb to a model
func JourneyModelToPb(mJourney *models.Journey) *journeypb.Journey {
	journeyMapper.CustomMappers = []mapper.CustomFieldMapper{TimeToTimestamp}
	journeyMapper.FieldNameMaps = map[string]string{
		"SegmentID": "SegmentId",
		"SegmentId": "SegmentID",
	}

	//TODO: Create custom mapper for Status
	journeyMapper.IgnoreDestFields = []string{"Model", "Status", "XXX_NoUnkeyedLiteral", "XXX_unrecognized", "XXX_sizecache"}
	pbJourney := &journeypb.Journey{}
	ret := journeyMapper.Map(mJourney, pbJourney)
	if len(ret.Errors) > 0 {
		log.Println(ret.Errors)
		return nil
	}

	return pbJourney

}

// TimeToTimestamp convert a time to timestamp
func TimeToTimestamp(sourceVal reflect.Value, sourceType reflect.Type, destVal reflect.Value, destType reflect.Type) (handled bool) {
	if destType.Kind() == reflect.Struct {
		if sourceVal.Type().Name() == "Time" && destVal.Type().Name() == "Timestamp" {
			if sourceType.Kind() == reflect.Struct {
				convertToIface := sourceVal.Interface().(time.Time)
				val, err := ptypes.TimestampProto(convertToIface)
				if err != nil {
					return false

				}
				destVal.Set(reflect.ValueOf(*val))
				return true
			}
		}
	}
	return false
}
