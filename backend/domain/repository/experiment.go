package repository

import (
	"context"

	"github.com/tetsuzawa/web-spat/domain"
)

// IExperimentRepository represent repository of the experiments.
// Expect implementation by the infrastructure layer.
type IExperimentRepository interface {
	ListMDDActive(ctx context.Context) ([]*domain.ExperimentMDD, error)
	ListMDDInactive(ctx context.Context) ([]*domain.ExperimentMDD, error)
}
