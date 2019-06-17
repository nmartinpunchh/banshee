package mapper

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"
	"github.com/nmartinpunchh/banshee/internal/models"
	journeypb "github.com/nmartinpunchh/banshee/pb/punchh/journey"
	workflowpb "github.com/nmartinpunchh/banshee/pb/punchh/workflow"
	log "github.com/sirupsen/logrus"
)

func TestJourneyPbToModel(t *testing.T) {
	now := time.Now()
	pbStartTime, _ := ptypes.TimestampProto(now)
	jpb1 := &journeypb.Journey{
		SegmentId:        "sdf",
		StartTime:        pbStartTime,
		EndTime:          pbStartTime,
		ControlGroupSize: int64(10),
		GuestEntryLimit:  int64(3),
		Workflow:         &workflowpb.Workflow{},
	}

	jm := &models.Journey{
		SegmentID:        "sdf",
		StartTime:        now,
		EndTime:          now,
		ControlGroupSize: 10,
		GuestEntryLimit:  3,
		Workflow:         &models.Workflow{},
	}

	type args struct {
		jpb *journeypb.Journey
	}
	tests := []struct {
		name string
		args args
		want *models.Journey
	}{
		{
			name: "test1",
			args: args{jpb: jpb1},
			want: jm,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JourneyPbToModel(tt.args.jpb); !reflect.DeepEqual(got, tt.want) {
				r, _ := json.Marshal(got)
				marshaller1 := jsonpb.Marshaler{}

				r2, _ := marshaller1.MarshalToString(tt.args.jpb)

				log.Println(fmt.Sprintf("%s", r))
				log.Println(r2)

				t.Errorf("JourneyPbToModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
