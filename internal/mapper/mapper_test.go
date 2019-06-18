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
		StartTime:        pbStartTime,
		EndTime:          pbStartTime,
		ControlGroupSize: int64(10),
		GuestEntryLimit:  int64(3),
		Workflow: &workflowpb.Workflow{
			Root: &workflowpb.Statement{
				ActivityInvocation: &workflowpb.ActivityInvocation{

					Name:      "map test",
					Result:    "123",
					Arguments: []string{"1a", "2a", "3a"},
				},
			},
		},
	}

	got := JourneyPbToModel(jpb1)
	assert.NotNil(t, got)
	assert.Equal(t, got.GuestEntryLimit, int(jpb1.GuestEntryLimit))
	assert.Equal(t, got.ControlGroupSize, int(jpb1.ControlGroupSize))
	assert.Equal(t, got.Workflow.Root.ActivityInvocation.Name, jpb1.Workflow.Root.ActivityInvocation.Name)
	assert.Equal(t, got.Workflow.Root.ActivityInvocation.Result, jpb1.Workflow.Root.ActivityInvocation.Result)
	assert.Equal(t, len(got.Workflow.Root.ActivityInvocation.Arguments), len(jpb1.Workflow.Root.ActivityInvocation.Arguments))
}
func TestJourneyModelToPb(t *testing.T) {
	now := time.Now()
	mjb := &models.Journey{
		StartTime:        now,
		EndTime:          now,
		ControlGroupSize: 10,
		GuestEntryLimit:  3,
		Workflow:         &models.Workflow{},
	}

	got := JourneyModelToPb(mjb)
	// assert.Equal(t, got.StartTime, jpb1.StartTime)
	// assert.Equal(t, got.EndTime,mjb.EndTime)
	assert.Equal(t, got.GuestEntryLimit, int64(mjb.GuestEntryLimit))
	assert.Equal(t, got.ControlGroupSize, int64(mjb.ControlGroupSize))
}
