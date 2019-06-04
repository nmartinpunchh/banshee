package models

import "github.com/jinzhu/gorm"

// Workflow is the type used to express the workflow definition. Variables are a map of valuables. Variables can be
// used as input to Activity.
type Workflow struct {
	//	Variables map[string]string
	gorm.Model
	Root *Statement
}

// Statement is the building block of dsl workflow. A Statement can be a simple ActivityInvocation or it
// could be a Sequence or Parallel.
type Statement struct {
	gorm.Model
	WorkflowID         int
	ActivityInvocation *ActivityInvocation
	Sequence           *Sequence
	Parallel           *Parallel
}

// Sequence consist of a collection of Statements that runs in sequential.
type Sequence struct {
	gorm.Model
	StatementID int
	Elements    []*Statement `gorm:"polymorphic:Statement;"`
}

// Parallel can be a collection of Statements that runs in parallel.
type Parallel struct {
	gorm.Model
	StatementID int
	Branches    []*Statement `gorm:"polymorphic:Statement;"`
}

// ActivityInvocation is used to express invoking an Activity. The Arguments defined expected arguments as input to
// the Activity, the result specify the name of variable that it will store the result as which can then be used as
// arguments to subsequent ActivityInvocation.
type ActivityInvocation struct {
	gorm.Model
	StatementID int
	Name        string
	Arguments   []*Argument
	Result      string
}

// Argument ..
type Argument struct {
	gorm.Model
	ActivityInvocationID int
	Argument             string
}
