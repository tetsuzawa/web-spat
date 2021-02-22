package domain

import "time"

type ExperimentMDDCW struct {
	ID        int64
	CreatedAt time.Time
}

type ExperimentMDDCWActive struct {
	ID           int64
	ExperimentID int64
	CreatedAt    time.Time
}

type ExperimentMDDCWInactive struct {
	ID           int64
	ExperimentID int64
	CreatedAt    time.Time
}

type ExperimentMDDCWDetail struct {
	ID                          int64
	ExperimentID                int64
	QuestplusParameterNormcdfID int64
	Name                        ExperimentName
	Description                 ExperimentDescription
	Azimuth                     Azimuth // 0[10^-1 deg] is the front, 900[10^-1 deg] is the right side
	Altitude                    uint64 // 0[10^-1 deg] is the front, 900[10^-1 deg] is the zenith, -900[10^-1 deg] is the nadir
	CoordinateVariableID        int64
	Width                       uint64 // [10^-1 deg]
	VelocityRangeLower          uint64 // [10^-1 deg/sec]
	VelocityRangeUpper          uint64 // [10^-1 deg/sec]
	VelocityRangeStep           uint64 // [10^-1 deg/sec]
	NumTrials                   uint64
	CreatedAt                   time.Time
}

type MCoordinateVariable struct {
	ID        int64
	Type      string // azimuth or altitude
	CreatedAt time.Time
}

type QuestplusParameterNormcdf struct {
	ID                        int64
	QuestplusParameterJSONURL string
	CreatedAt                 time.Time
}

type ResultMDDCW struct {
	ID             int64
	ExperimentID   int64
	ResultURL      string
	SubjectID      int64
	Mean           float64
	Sd             float64
	LowerAsymptote float64
	LapseRate      float64
	CreatedAt      time.Time
}

type Subject struct {
	ID                     int64
	Sex                    string // ISO5218
	Age                    uint64
	DeafAndHearingImpaired int64
	CreatedAt              time.Time
}
