package usecase

import (
	"context"
	"fmt"

	"github.com/tetsuzawa/web-spat/domain"
	"github.com/tetsuzawa/web-spat/domain/repository"
)

// IExperimentUseCase represent UseCase of the experiments.
type IExperimentUseCase interface {
	CreateMDD(ctx context.Context, e *domain.ExperimentMDDData) (*domain.ExperimentMDDData, error)
	ListMDDActive(ctx context.Context) ([]*domain.ExperimentMDDData, error)
	ListMDDInactive(ctx context.Context) ([]*domain.ExperimentMDDData, error)
}

type experimentUseCase struct {
	r repository.IExperimentRepository
}

func NewExperimentUseCase(r repository.IExperimentRepository) IExperimentUseCase {
	return &experimentUseCase{r: r}
}

func (u *experimentUseCase) CreateMDD(ctx context.Context, e *domain.ExperimentMDDData) (*domain.ExperimentMDDData, error) {
	return nil, nil
}

func (u *experimentUseCase) ListMDDActive(ctx context.Context) ([]*domain.ExperimentMDDData, error) {
	experimentMDDs, err := u.r.ListMDDActive(ctx)
	if err != nil {
		return nil, fmt.Errorf("repositoy error -> %w", err)
	}
	experimentMDDData := make([]*domain.ExperimentMDDData, len(experimentMDDs))
	for _, v := range experimentMDDs {
		experimentMDDData = append(experimentMDDData, domain.NewExperimentMDDData(v))
	}
	return experimentMDDData, nil
}

func (u *experimentUseCase) ListMDDInactive(ctx context.Context) ([]*domain.ExperimentMDDData, error) {
	experimentMDDs, err := u.r.ListMDDInactive(ctx)
	if err != nil {
		return nil, fmt.Errorf("repositoy error -> %w", err)
	}
	experimentMDDData := make([]*domain.ExperimentMDDData, len(experimentMDDs))
	for _, v := range experimentMDDs {
		experimentMDDData = append(experimentMDDData, domain.NewExperimentMDDData(v))
	}
	return experimentMDDData, nil
}
