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
	FindMDDById(ctx context.Context, id domain.ExperimentIdData) (*domain.ExperimentMDDData, error)
}

type experimentUseCase struct {
	r repository.IExperimentRepository
}

func NewExperimentUseCase(r repository.IExperimentRepository) IExperimentUseCase {
	return &experimentUseCase{r: r}
}

func (u *experimentUseCase) CreateMDD(ctx context.Context, e *domain.ExperimentMDDData) (*domain.ExperimentMDDData, error) {
	experimentService := domain.NewExperimentService()
	err := experimentService.ValidateExperimentMDDData(e)
	if err != nil {
		return nil, fmt.Errorf("domain error -> %w", err)
	}
	e, err = u.r.CreateMDD(ctx, e)
	if err != nil {
		return nil, fmt.Errorf("repository error -> %w", err)
	}
	return e, nil
}

func (u *experimentUseCase) ListMDDActive(ctx context.Context) ([]*domain.ExperimentMDDData, error) {
	experimentMDDs, err := u.r.ListMDDActive(ctx)
	if err != nil {
		return nil, fmt.Errorf("repositoy error -> %w", err)
	}
	return experimentMDDs, nil
}

func (u *experimentUseCase) ListMDDInactive(ctx context.Context) ([]*domain.ExperimentMDDData, error) {
	experimentMDDs, err := u.r.ListMDDInactive(ctx)
	if err != nil {
		return nil, fmt.Errorf("repositoy error -> %w", err)
	}
	return experimentMDDs, nil
}

func (u *experimentUseCase) FindMDDById(ctx context.Context, id domain.ExperimentIdData) (*domain.ExperimentMDDData, error) {
	experimentMDD, err := u.r.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("repositoy error -> %w", err)
	}
	return experimentMDD, nil
}
