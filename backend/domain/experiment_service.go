package domain

import "fmt"

type IExperimentService interface {
	ValidateExperimentMDDData(v *ExperimentMDDData) error
}

type experimentService struct{}

func NewExperimentService() IExperimentService {
	return &experimentService{}
}

func (s *experimentService) ValidateExperimentMDDData(v *ExperimentMDDData) error {
	_, err := NewExperimentMDDFromData(v)
	return fmt.Errorf("failed to validate ExperimentMDDData -> %w", err)
}
