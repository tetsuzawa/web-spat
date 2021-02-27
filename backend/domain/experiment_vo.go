package domain

import "time"

// ExperimentName is name of a experiment.
// The length must be in 64 characters.
type ExperimentName string

// NewExperimentName generates ExperimentName.
// If the length of the argument is too long, it returns ErrTooLongExperimentName.
func NewExperimentName(v string) (ExperimentName, error) {
	if len(v) > 64 {
		return "", ErrTooLongExperimentName
	}
	return ExperimentName(v), nil
}

func (v *ExperimentName) ToString() string {
	return string(*v)
}

// ExperimentDescription is description of a experiment.
// The length must be in 1024 characters.
type ExperimentDescription string

// NewExperimentDescription generates ExperimentDescription.
// If the length of the argument is too long, it returns ErrTooLongExperimentDescription.
func NewExperimentDescription(v string) (ExperimentDescription, error) {
	if len(v) > 1024 {
		return "", ErrTooLongExperimentDescription
	}
	return ExperimentDescription(v), nil
}

func (v *ExperimentDescription) ToString() string {
	return string(*v)
}

// Azimuth is the horizontal angle.
// 0[10^-1 deg] is the front, 900[10^-1 deg] is the right side,
// 1800[10^-1 deg] is the behind and 2700[10^-1 deg] is the left side.
type Azimuth uint64

// NewAzimuth generates Azimuth.
// If the value of argument is invalid, it returns ErrInvalidAzimuthValue.
func NewAzimuth(v uint64) (Azimuth, error) {
	if v >= 3600 {
		return 0, ErrInvalidAzimuthValue
	}
	return Azimuth(v), nil
}

func (v *Azimuth) ToInt() int {
	return int(*v)
}

// Altitude is the vertical angle.
// 0[10^-1 deg] is the front, 900[10^-1 deg] is the zenith, -900[10^-1 deg] is the nadir.
type Altitude int64

// NewAltitude generates Altitude.
// If the value of argument is invalid, it returns ErrInvalidAltitudeValue.
func NewAltitude(v uint64) (Altitude, error) {
	if v < -900 || 900 < v {
		return 0, ErrInvalidAltitudeValue
	}
	return Altitude(v), nil
}

func (v *Altitude) ToInt() int {
	return int(*v)
}

// CoordinateVariable is the variable
type CoordinateVariable string

const (
	CoordinateVariableAzimuth  = CoordinateVariable("azimuth")
	CoordinateVariableAltitude = CoordinateVariable("altitude")
)

// Width is the width of moving angle.
type Width uint64

// NewWidth generates Width.
func NewWidth(v uint64) Width {
	return Width(v)
}

func (v *Width) ToInt() int {
	return int(*v)
}

type VelocityRangeLower uint64

// NewVelocityRangeLower generates NewVelocityRangeLower.
func NewVelocityRangeLower(v uint64) VelocityRangeLower {
	return VelocityRangeLower(v)
}

func (v *VelocityRangeLower) ToInt() int {
	return int(*v)
}

type VelocityRangeUpper uint64

// NewVelocityRangeUpper generates NewVelocityRangeUpper.
func NewVelocityRangeUpper(v uint64) VelocityRangeUpper {
	return VelocityRangeUpper(v)
}

func (v *VelocityRangeUpper) ToInt() int {
	return int(*v)
}

type VelocityRangeStep uint64

// NewVelocityRangeStep generates NewVelocityRangeStep.
func NewVelocityRangeStep(v uint64) VelocityRangeStep {
	return VelocityRangeStep(v)
}

func (v *VelocityRangeStep) ToInt() int {
	return int(*v)
}

type NumTrials uint64

// NewNumTrials generates NewNumTrials.
func NewNumTrials(v uint64) NumTrials {
	return NumTrials(v)
}

func (v *NumTrials) ToInt() int {
	return int(*v)
}

type ExperimentMDDCW struct {
	QuestPlusParameterNormCDF QuestPlusParameterNormCDF
	Name                      ExperimentName
	Description               ExperimentDescription
	Azimuth                   Azimuth  // 0[10^-1 deg] is the front, 900[10^-1 deg] is the right side
	Altitude                  Altitude // 0[10^-1 deg] is the front, 900[10^-1 deg] is the zenith, -900[10^-1 deg] is the nadir
	CoordinateVariable        CoordinateVariable
	Width                     Width              // [10^-1 deg]
	VelocityRangeLower        VelocityRangeLower // [10^-1 deg/sec]
	VelocityRangeUpper        VelocityRangeUpper // [10^-1 deg/sec]
	VelocityRangeStep         VelocityRangeStep  // [10^-1 deg/sec]
	NumTrials                 NumTrials
	CreatedAt                 time.Time
}

type QuestPlusParameterNormCDF struct {
	StimDomain            StimDomainNormCDF     `json:"stim_domain"`
	ParamDomain           ParamDomainNormCDF    `json:"param_domain"`
	OutcomeDomain         OutcomeDomain         `json:"outcome_domain"`
	Prior                 PriorNormCDF          `json:"prior"`
	Func                  Func                  `json:"func"`
	StimScale             StimScase             `json:"stim_scale"`
	StimSelectionMethod   StimSelectionMethod   `json:"stim_selection_method"`
	ParamEstimationMethod ParamEstimationMethod `json:"param_estimation_method"`
}

type StimDomainNormCDF struct {
	Intensity []float64 `json:"intensity"`
}

type ParamDomainNormCDF struct {
	Mean           []float64 `json:"mean"`
	SD             []float64 `json:"sd"`
	LowerAsymptote []float64 `json:"lower_asymptote"`
	LapseRate      []float64 `json:"lapse_rate"`
}

type OutcomeDomain struct {
	Response struct{} `json:"response"`
}

type PriorNormCDF struct {
	Mean           []float64 `json:"mean"`
	SD             []float64 `json:"sd"`
	LowerAsymptote []float64 `json:"lower_asymptote"`
	LapseRate      []float64 `json:"lapse_rate"`
}

type Func string

const (
	NormCDF = Func("norm_cdf")
)

type StimScase string

const (
	Linear = StimScase("linear")
	Log10  = StimScase("log10")
)

type StimSelectionMethod string

const (
	MinEntropy  = StimSelectionMethod("min_entropy")
	MinNEntropy = StimSelectionMethod("min_n_entropy")
)

type ParamEstimationMethod string

const (
	Mode = ParamEstimationMethod("mode")
	Mean = ParamEstimationMethod("mean")
)
