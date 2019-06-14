package mapper

import (
	"reflect"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/nmartinpunchh/banshee/internal/models"
	journeypb "github.com/nmartinpunchh/banshee/pb/punchh/journey"
)

func TestJourneyPbToModel(t *testing.T) {
	now := time.Now()
	pbStartTime, _ := ptypes.TimestampProto(now)
	jpb1 := &journeypb.Journey{
		SegmentId: "sdf",
		StartTime: pbStartTime,
		EndTime:   pbStartTime,
	}

	jm := &models.Journey{
		SegmentID: "sdf",
		StartTime: now,
		EndTime:   now,
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
				t.Errorf("JourneyPbToModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
