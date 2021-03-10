package persistence_mock

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/tetsuzawa/web-spat/infrastructure/dbutil"
	"os"

	"github.com/tetsuzawa/web-spat/domain"
	"github.com/tetsuzawa/web-spat/domain/repository"
)

// experimentRepository Implements repository.ExperimentsRepository
type experimentRepository struct {
	db *sqlx.DB
}

// NewExperimentRepository returns initialized ExperimentRepositoryImpl
func NewExperimentRepository(db *sqlx.DB) repository.IExperimentRepository {
	return &experimentRepository{db: db}
}

func (r *experimentRepository) CreateMDD(ctx context.Context, e *domain.ExperimentMDDData) (*domain.ExperimentMDDData, error) {

	if err := dbutil.TXHandler(r.db, func(tx *sqlx.Tx) error {
		res, err := tx.ExecContext(ctx, `
INSERT INTO experiment_mdd
`)
		if err != nil {
			return fmt.Errorf("failed to insert to `experiment_mdd` table -> %w", err)
		}

		experimentID, err := res.LastInsertId()
		if err != nil {
			return fmt.Errorf("failed to get last inserted ID -> %w", err)
		}

		stmt, err := tx.PreparexContext(ctx, `
INSERT INTO experiment_mdd_active (experiment_id) 
VALUES (?)
`)
		if err != nil {
			return fmt.Errorf("failed to prepare statement for `experiment_mdd_active` table -> %w", err)
		}
		_, err = stmt.ExecContext(ctx, experimentID)
		if err != nil {
			return fmt.Errorf("failed to insert to `experiment_mdd_active` table -> %w", err)
		}

		stmt, err = tx.PreparexContext(ctx, `
INSERT INTO experiment_mdd_active (experiment_id) 
VALUES (?)
`)
		if err != nil {
			return fmt.Errorf("failed to prepare statement for `experiment_mdd_active` table -> %w", err)
		}
		_, err = stmt.ExecContext(ctx, experimentID)
		if err != nil {
			return fmt.Errorf("failed to insert to `experiment_mdd_active` table -> %w", err)
		}

		b, err := json.Marshal(e.QuestPlusParameterNormCDF)
		if err != nil {
			return fmt.Errorf("failed to marshal QuestPlusParameterNormCDF -> %w", err)
		}
		tmpDirPath := os.TempDir()
		uid, err := uuid.NewRandom()
		if err != nil {
			return fmt.Errorf("failed to generate uuid -> %w", err)
		}
		questplusParameterPath := tmpDirPath + "questplus_param_" + uid.String() + ".json"
		f, err := os.Create(questplusParameterPath)
		defer f.Close()
		if err != nil {
			return fmt.Errorf("failed to create questplus param to json file -> %w", err)
		}
		if _, err := f.Write(b); err != nil {
			return fmt.Errorf("failed to write questplus param to json file -> %w", err)
		}

		stmt, err = tx.PreparexContext(ctx, `
INSERT INTO questplus_parameter_normcdf (questplus_parameter_json_url) 
VALUES (?)
`)
		if err != nil {
			return fmt.Errorf("failed to prepare statement for `questplus_parameter_normcdf` table -> %w", err)
		}
		res, err = stmt.ExecContext(ctx, questplusParameterPath)
		if err != nil {
			return fmt.Errorf("failed to insert to `experiment_mdd_active` table -> %w", err)
		}

		questplusParameterID, err := res.LastInsertId()
		if err != nil {
			return fmt.Errorf("failed to get last inserted ID -> %w", err)
		}

		stmt, err = tx.PreparexContext(ctx, `
INSERT INTO experiment_mdd_detail (experiment_id, questplus_parameter_normcdf_id, name, description, azimuth, altitude, coordinate_variable, moving_sound_constant, moving_sound_constant_value, num_trials) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`)
		if err != nil {
			return fmt.Errorf("failed to prepare statement for `experiment_mdd_detail` table -> %w", err)
		}
		_, err = stmt.ExecContext(ctx, experimentID, questplusParameterID, e.Name, e.Description, e.Azimuth, e.Altitude, e.CoordinateVariable, e.MovingSoundConstant, e.MovingSoundConstantValue, e.NumTrials)
		if err != nil {
			return fmt.Errorf("failed to insert to `experiment_mdd_detail` table -> %w", err)
		}
		e.Id = uint64(experimentID)

		return nil

	}); err != nil {
		return nil, fmt.Errorf("failed to exec in the transaction -> %w", err)
	}
	return e, nil
}

func (r *experimentRepository) ListMDDActive(ctx context.Context) ([]*domain.ExperimentMDDData, error) {
	experiments := []*domain.ExperimentMDDData{
		{
			QuestPlusParameterNormCDF: domain.QuestPlusParameterNormCDFData{
				StimDomain: domain.StimDomainNormCDFData{
					Intensity: []float64{1, 2, 3},
				},
				ParamDomain: domain.ParamDomainNormCDFData{
					Mean:           []float64{7, 8, 9},
					SD:             []float64{4, 5, 6},
					LowerAsymptote: []float64{0.5},
					LapseRate:      []float64{0.01, 0.02, 0.03},
				},
				OutcomeDomain: domain.OutcomeDomainData{
					Response: []string{"correct", "incorrect"},
				},
				Prior: domain.PriorNormCDFData{
					Mean:           []float64{0.32, 0.32, 0.32},
					SD:             []float64{0.32, 0.32, 0.32},
					LowerAsymptote: []float64{0.32},
					LapseRate:      []float64{0.32, 0.32, 0.32},
				},
				Func:                  string(domain.FuncNormCDF),
				StimScale:             string(domain.StimScaleLinear),
				StimSelectionMethod:   string(domain.StimSelectionMethodMinEntropy),
				ParamEstimationMethod: string(domain.ParamEstimationMethodMean),
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

func (r *experimentRepository) ListMDDInactive(ctx context.Context) ([]*domain.ExperimentMDDData, error) {
	experiments := []*domain.ExperimentMDDData{
		{
			QuestPlusParameterNormCDF: domain.QuestPlusParameterNormCDFData{
				StimDomain: domain.StimDomainNormCDFData{
					Intensity: []float64{1, 2, 3},
				},
				ParamDomain: domain.ParamDomainNormCDFData{
					Mean:           []float64{7, 8, 9},
					SD:             []float64{4, 5, 6},
					LowerAsymptote: []float64{0.5},
					LapseRate:      []float64{0.01, 0.02, 0.03},
				},
				OutcomeDomain: domain.OutcomeDomainData{
					Response: []string{"correct", "incorrect"},
				},
				Prior: domain.PriorNormCDFData{
					Mean:           []float64{0.32, 0.32, 0.32},
					SD:             []float64{0.32, 0.32, 0.32},
					LowerAsymptote: []float64{0.32},
					LapseRate:      []float64{0.32, 0.32, 0.32},
				},
				Func:                  string(domain.FuncNormCDF),
				StimScale:             string(domain.StimScaleLinear),
				StimSelectionMethod:   string(domain.StimSelectionMethodMinEntropy),
				ParamEstimationMethod: string(domain.ParamEstimationMethodMean),
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
