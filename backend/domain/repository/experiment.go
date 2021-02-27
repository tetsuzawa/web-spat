package repository

import (
	"context"
	"github.com/tetsuzawa/web-spat/domain"
)

// IExperimentRepository represent repository of the experiments.
// Expect implementation by the infrastructure layer.
type IExperimentRepository interface {
	ListMDDCWActive(ctx context.Context) ([]*domain.ExperimentMDDCW, error)
	ListMDDCWInactive(ctx context.Context) ([]*domain.ExperimentMDDCW, error)
}
