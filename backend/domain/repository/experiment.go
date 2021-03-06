package repository

import (
	"context"

	"github.com/tetsuzawa/web-spat/domain"
)

// IExperimentRepository represent repository of the experiments.
// Expect implementation by the infrastructure layer.
type IExperimentRepository interface {
	CreateMDD(ctx context.Context, e *domain.ExperimentMDDData) (*domain.ExperimentMDDData, error)
	ListMDDActive(ctx context.Context) ([]*domain.ExperimentMDDData, error)
	ListMDDInactive(ctx context.Context) ([]*domain.ExperimentMDDData, error)
	FindByID(ctx context.Context, id domain.ExperimentIdData) (*domain.ExperimentMDDData, error)
	CreateMDDResult(ctx context.Context, res *domain.ResultMDDData) (*domain.ResultMDDData, error)
}
