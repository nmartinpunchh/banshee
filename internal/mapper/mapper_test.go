package mapper

import (
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/nmartinpunchh/banshee/internal/models"
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
	assert.Equal(t, got.GuestEntryLimit, int(jpb1.GuestEntryLimit))
	assert.Equal(t, got.ControlGroupSize, int(jpb1.ControlGroupSize))
}
func TestJourneyModelToPb(t *testing.T) {
	now := time.Now()
	mjb := &models.Journey{
		SegmentID:        "sdf",
		StartTime:        now,
		EndTime:          now,
		ControlGroupSize: 10,
		GuestEntryLimit:  3,
		Workflow:         &models.Workflow{},
	}

	got := JourneyModelToPb(mjb)
	// assert.Equal(t, got.StartTime, jpb1.StartTime)
	// assert.Equal(t, got.EndTime,mjb.EndTime)
	assert.Equal(t, got.SegmentId, mjb.SegmentID)
	assert.Equal(t, got.GuestEntryLimit, int64(mjb.GuestEntryLimit))
	assert.Equal(t, got.ControlGroupSize, int64(mjb.ControlGroupSize))
}
