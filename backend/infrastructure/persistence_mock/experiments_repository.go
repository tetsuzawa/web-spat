package persistence_mock

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/tetsuzawa/web-spat/domain"
	"github.com/tetsuzawa/web-spat/domain/repository"
	"github.com/tetsuzawa/web-spat/infrastructure/dbutil"
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
	rows, err := r.db.QueryxContext(ctx, `
WITH ex AS (
    SELECT * from experiment_mdd_detail AS ex_detail
    INNER JOIN experiment_mdd_active AS ex_active ON ex_detail.experiment_id = ex_active.experiment_id
    LEFT JOIN questplus_parameter_normcdf AS qp ON ex_detail.questplus_parameter_normcdf_id = qp.id
)

SELECT experiment_id, name, description, azimuth, altitude, coordinate_variable, moving_sound_constant, moving_sound_constant_value, num_trials, questplus_parameter_json_url
FROM ex
`)
	if err != nil {
		return nil, fmt.Errorf("failed to exec query -> %w", err)
	}
	defer rows.Close()

	var experiments []*domain.ExperimentMDDData
	for rows.Next() {
		e := &domain.ExperimentMDDData{}
		var qpParamURL string
		err := rows.Scan(e.Id, e.Name, e.Description, e.Azimuth, e.Altitude, e.CoordinateVariable, e.MovingSoundConstant, e.MovingSoundConstantValue, e.NumTrials, qpParamURL)
		if err != nil {
			return nil, fmt.Errorf("failed to scan ExperimentMDDData -> %w", err)
		}
		f, err := os.Open(qpParamURL)
		if err != nil {
			return nil, fmt.Errorf("failed to open QuestPlusParameter file -> %w", err)
		}
		b, err := io.ReadAll(f)
		if err != nil {
			return nil, fmt.Errorf("failed to read QuestPlusParameter file -> %w", err)
		}
		if err := json.Unmarshal(b, &e.QuestPlusParameterNormCDF); err != nil {
			return nil, fmt.Errorf("failed to unmarshal QuestPlusParameter -> %w", err)
		}
		experiments = append(experiments, e)
	}
	return experiments, nil
}

func (r *experimentRepository) ListMDDInactive(ctx context.Context) ([]*domain.ExperimentMDDData, error) {
	rows, err := r.db.QueryxContext(ctx, `
WITH ex AS (
    SELECT * from experiment_mdd_detail AS ex_detail
    INNER JOIN experiment_mdd_inactive AS ex_inactive ON ex_detail.experiment_id = ex_inactive.experiment_id
    LEFT JOIN questplus_parameter_normcdf AS qp ON ex_detail.questplus_parameter_normcdf_id = qp.id
)

SELECT experiment_id, name, description, azimuth, altitude, coordinate_variable, moving_sound_constant, moving_sound_constant_value, num_trials, questplus_parameter_json_url
FROM ex
`)
	if err != nil {
		return nil, fmt.Errorf("failed to exec query -> %w", err)
	}
	defer rows.Close()

	var experiments []*domain.ExperimentMDDData
	for rows.Next() {
		e := &domain.ExperimentMDDData{}
		var qpParamURL string
		err := rows.Scan(e.Id, e.Name, e.Description, e.Azimuth, e.Altitude, e.CoordinateVariable, e.MovingSoundConstant, e.MovingSoundConstantValue, e.NumTrials, qpParamURL)
		if err != nil {
			return nil, fmt.Errorf("failed to scan ExperimentMDDData -> %w", err)
		}
		f, err := os.Open(qpParamURL)
		if err != nil {
			return nil, fmt.Errorf("failed to open QuestPlusParameter file -> %w", err)
		}
		b, err := io.ReadAll(f)
		if err != nil {
			return nil, fmt.Errorf("failed to read QuestPlusParameter file -> %w", err)
		}
		if err := json.Unmarshal(b, &e.QuestPlusParameterNormCDF); err != nil {
			return nil, fmt.Errorf("failed to unmarshal QuestPlusParameter -> %w", err)
		}
		experiments = append(experiments, e)
	}
	return experiments, nil
}


func (r *experimentRepository)FindByID(ctx context.Context, id domain.ExperimentIdData) (*domain.ExperimentMDDData, error){
	row := r.db.QueryRowxContext(ctx, `
WITH ex AS (
    SELECT * from experiment_mdd_detail AS ex_detail
    INNER JOIN experiment_mdd_inactive AS ex_inactive ON ex_detail.experiment_id = ex_inactive.experiment_id
    LEFT JOIN questplus_parameter_normcdf AS qp ON ex_detail.questplus_parameter_normcdf_id = qp.id
    WHERE ex_detail.experiment_id = ?
)

SELECT experiment_id, name, description, azimuth, altitude, coordinate_variable, moving_sound_constant, moving_sound_constant_value, num_trials, questplus_parameter_json_url
FROM ex
`,id)
	if row.Err()==sql.ErrNoRows{
		return nil,nil
	}

	e := &domain.ExperimentMDDData{}
	var qpParamURL string
	err := row.Scan(e.Id, e.Name, e.Description, e.Azimuth, e.Altitude, e.CoordinateVariable, e.MovingSoundConstant, e.MovingSoundConstantValue, e.NumTrials, qpParamURL)
	if err != nil {
		return nil, fmt.Errorf("failed to scan ExperimentMDDData -> %w", err)
	}
	f, err := os.Open(qpParamURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open QuestPlusParameter file -> %w", err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read QuestPlusParameter file -> %w", err)
	}
	if err := json.Unmarshal(b, &e.QuestPlusParameterNormCDF); err != nil {
		return nil, fmt.Errorf("failed to unmarshal QuestPlusParameter -> %w", err)
	}
	return e, nil

}
