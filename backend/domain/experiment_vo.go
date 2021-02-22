package domain

type ExperimentName string

func NewExperimentName(v string) (ExperimentName, error) {
	if len(v) > 64 {
		return "", ErrTooLongExperimentName
	}
	return ExperimentName(v), nil
}
