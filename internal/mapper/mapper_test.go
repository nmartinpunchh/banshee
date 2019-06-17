package mapper

import (
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	journeypb "github.com/nmartinpunchh/banshee/pb/punchh/journey"
	workflowpb "github.com/nmartinpunchh/banshee/pb/punchh/workflow"
	"github.com/stretchr/testify/assert"
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

	got := JourneyPbToModel(jpb1)
	// assert.Equal(t, got.StartTime, jpb1.StartTime)
	// assert.Equal(t, got.EndTime, jpb1.EndTime)
	assert.Equal(t, got.SegmentID, jpb1.SegmentId)
	assert.Equal(t, int64(got.GuestEntryLimit), jpb1.GuestEntryLimit)
	assert.Equal(t, int64(got.ControlGroupSize), jpb1.ControlGroupSize)
}
