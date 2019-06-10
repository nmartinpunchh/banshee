package mapper

import (
	mapper "github.com/nicolas-martin/go-automapper"
)

// M ..
var M = mapper.Mapper{
	PanicOnIncompatibleTypes: false,
	PanicOnMissingField:      false,
}
