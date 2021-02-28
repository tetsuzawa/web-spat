package persistence_mock

import (
	"context"

	"github.com/tetsuzawa/web-spat/domain"
	"github.com/tetsuzawa/web-spat/domain/repository"
)

// experimentRepository Implements repository.ExperimentsRepository
type experimentRepository struct{}

// NewExperimentRepository returns initialized ExperimentRepositoryImpl
func NewExperimentRepository() repository.IExperimentRepository {
	return &experimentRepository{}
}

func (r *experimentRepository) ListMDDActive(ctx context.Context) ([]*domain.ExperimentMDD, error) {
	experiments := []*domain.ExperimentMDD{
		{
			QuestPlusParameterNormCDF: domain.QuestPlusParameterNormCDF{
				StimDomain: domain.StimDomainNormCDF{
					Intensity: []float64{1, 2, 3},
				},
				ParamDomain: domain.ParamDomainNormCDF{
					Mean:           []domain.Mean{7, 8, 9},
					SD:             []domain.SD{4, 5, 6},
					LowerAsymptote: []domain.LowerAsymptote{0.5},
					LapseRate:      []domain.LapseRate{0.01, 0.02, 0.03},
				},
				OutcomeDomain: domain.OutcomeDomain{
					Response: "correct",
				},
				Prior: domain.PriorNormCDF{
					Mean:           []domain.Probability{0.32, 0.32, 0.32},
					SD:             []domain.Probability{0.32, 0.32, 0.32},
					LowerAsymptote: []domain.Probability{0.32},
					LapseRate:      []domain.Probability{0.32, 0.32, 0.32},
				},
				Func:                  domain.FuncNormCDF,
				StimScale:             domain.StimScaleLinear,
				StimSelectionMethod:   domain.StimSelectionMethodMinEntropy,
				ParamEstimationMethod: domain.ParamEstimationMethodMean,
			},
			Name:                     "移動方向弁別 角度",
			Description:              "角度",
			Azimuth:                  450,
			Altitude:                 0,
			CoordinateVariable:       "azimuth",
			MovingSoundConstant:      "velocity",
			MovingSoundConstantValue: 800,
			NumTrials:                100,
		},
	}
	return experiments, nil
}

func (r *experimentRepository) ListMDDInactive(ctx context.Context) ([]*domain.ExperimentMDD, error) {
	experiments := []*domain.ExperimentMDD{
		{
			QuestPlusParameterNormCDF: domain.QuestPlusParameterNormCDF{
				StimDomain: domain.StimDomainNormCDF{
					Intensity: []float64{1, 2, 3},
				},
				ParamDomain: domain.ParamDomainNormCDF{
					Mean:           []domain.Mean{7, 8, 9},
					SD:             []domain.SD{4, 5, 6},
					LowerAsymptote: []domain.LowerAsymptote{0.5},
					LapseRate:      []domain.LapseRate{0.01, 0.02, 0.03},
				},
				OutcomeDomain: domain.OutcomeDomain{
					Response: "correct",
				},
				Prior: domain.PriorNormCDF{
					Mean:           []domain.Probability{0.32, 0.32, 0.32},
					SD:             []domain.Probability{0.32, 0.32, 0.32},
					LowerAsymptote: []domain.Probability{0.32},
					LapseRate:      []domain.Probability{0.32, 0.32, 0.32},
				},
				Func:                  domain.FuncNormCDF,
				StimScale:             domain.StimScaleLinear,
				StimSelectionMethod:   domain.StimSelectionMethodMinEntropy,
				ParamEstimationMethod: domain.ParamEstimationMethodMean,
			},
			Name:                     "移動方向弁別 角度",
			Description:              "角度",
			Azimuth:                  450,
			Altitude:                 0,
			CoordinateVariable:       "azimuth",
			MovingSoundConstant:      "velocity",
			MovingSoundConstantValue: 800,
			NumTrials:                100,
		},
	}
	return experiments, nil
}
