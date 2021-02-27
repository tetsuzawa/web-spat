package domain

import "time"

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
