package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Workflow is the type used to express the workflow definition. Variables are a map of valuables. Variables can be
// used as input to Activity.
type Workflow struct {
	gorm.Model
	Root      *Statement
	JourneyID int
}

// Statement is the building block of workflow. A Statement can be a simple ActivityInvocation or it
// could be a Sequence or Parallel.
type Statement struct {
	StatementID        int
	WorkflowID         int
	ActivityInvocation *ActivityInvocation
	Sequence           *Sequence
	Parallel           *Parallel
}

// Sequence consist of a collection of Statements that runs in sequential.
type Sequence struct {
	SequenceID  int
	StatementID int
	Elements    []*Statement `gorm:"polymorphic:Statement;"`
}

// Parallel can be a collection of Statements that runs in parallel.
type Parallel struct {
	ParallelID  int
	StatementID int
	Branches    []*Statement `gorm:"polymorphic:Statement;"`
}

// ActivityInvocation is used to express invoking an Activity. The Arguments defined expected arguments as input to
// the Activity, the result specify the name of variable that it will store the result as which can then be used as
// arguments to subsequent ActivityInvocation.
type ActivityInvocation struct {
	ActivityInvocationID int
	StatementID          int
	Name                 string
	Arguments            []*Argument
	Result               string
}

// Argument ..
type Argument struct {
	ArgumentID           int
	ActivityInvocationID int
	Argument             string
}

// Journey ..
type Journey struct {
	gorm.Model
	StartTime        time.Time
	EndTime          time.Time
	Status           Status
	SegmentID        string
	ControlGroupSize int
	GuestEntryLimit  int
	Workflow         Workflow
}

// Status ..
type Status int32

const (
	//StatusInvalid default value not set
	StatusInvalid Status = 0
	//StatusDraft ..
	StatusDraft Status = 1
	//StatusLive ..
	StatusLive Status = 2
	//StatusDisabled ..
	StatusDisabled Status = 3
)

// StatusName converts a status int value to it's corresponding string
var StatusName = map[int32]string{
	0: "STATUS_INVALID",
	1: "STATUS_DRAFT",
	2: "STATUS_LIVE",
	3: "STATUS_DISABLED",
}

// StatusValue converts the string representation of a status to an int
var StatusValue = map[string]int32{
	"STATUS_INVALID":  0,
	"STATUS_DRAFT":    1,
	"STATUS_LIVE":     2,
	"STATUS_DISABLED": 3,
}
