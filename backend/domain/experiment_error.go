package domain

import "fmt"

var (
	ErrTooLongExperimentName = fmt.Errorf("too long experiment name")
	ErrTooLongExperimentDescription = fmt.Errorf("too long experiment description")
)
