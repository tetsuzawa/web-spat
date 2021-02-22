package domain

// ExperimentName is name of a experiment.
// The length must be in 64 characters.
type ExperimentName string

// NewExperimentName generates ExperimentName.
func NewExperimentName(v string) (ExperimentName, error) {
	if len(v) > 64 {
		return "", ErrTooLongExperimentName
	}
	return ExperimentName(v), nil
}

// ExperimentDescription is description of a experiment.
// The length must be in 1024 characters.
type ExperimentDescription string

// NewExperimentDescription generates ExperimentDescription.
func NewExperimentDescription(v string) (ExperimentDescription, error) {
	if len(v) > 1024 {
		return "", ErrTooLongExperimentDescription
	}
	return ExperimentDescription(v), nil
}
