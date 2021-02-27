package domain

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

// Velocity is the Velocity of moving angle.
type Velocity uint64

// NewVelocity generates Velocity.
func NewVelocity(v uint64) Velocity {
	return Velocity(v)
}

type VelocityRangeLower uint64

// NewVelocityRangeLower generates NewVelocityRangeLower.
func NewVelocityRangeLower(v uint64) VelocityRangeLower {
	return VelocityRangeLower(v)
}

type VelocityRangeUpper uint64

// NewVelocityRangeUpper generates NewVelocityRangeUpper.
func NewVelocityRangeUpper(v uint64) VelocityRangeUpper {
	return VelocityRangeUpper(v)
}

type VelocityRangeStep uint64

// NewVelocityRangeStep generates NewVelocityRangeStep.
func NewVelocityRangeStep(v uint64) VelocityRangeStep {
	return VelocityRangeStep(v)
}

type NumTrials uint64

// NewNumTrials generates NewNumTrials.
func NewNumTrials(v uint64) NumTrials {
	return NumTrials(v)
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
}

type QuestPlusParameterNormCDF struct {
	StimDomain            StimDomainNormCDF     `json:"stim_domain"`
	ParamDomain           ParamDomainNormCDF    `json:"param_domain"`
	OutcomeDomain         OutcomeDomain         `json:"outcome_domain"`
	Prior                 PriorNormCDF          `json:"prior"`
	Func                  Func                  `json:"func"`
	StimScale             StimScale             `json:"stim_scale"`
	StimSelectionMethod   StimSelectionMethod   `json:"stim_selection_method"`
	ParamEstimationMethod ParamEstimationMethod `json:"param_estimation_method"`
}

type StimDomainNormCDF struct {
	Intensity []float64 `json:"intensity"`
}

type ParamDomainNormCDF struct {
	Mean           []Mean           `json:"mean"`
	SD             []SD             `json:"sd"`
	LowerAsymptote []LowerAsymptote `json:"lower_asymptote"`
	LapseRate      []LapseRate      `json:"lapse_rate"`
}

type Mean float64

func NewMean(v float64) Mean {
	return Mean(v)
}

type SD float64

func NewSD(v float64) (SD, error) {
	if v < 0 {
		return 0, ErrInvalidSDValue
	}
	return SD(v), nil
}

type LowerAsymptote float64

func NewLowerAsymptote(v float64) (LowerAsymptote, error) {
	if v < 0 || 1 < v {
		return 0, ErrInvalidLowerAsymptoteValue
	}
	return LowerAsymptote(v), nil
}

type LapseRate float64

func NewLapseRate(v float64) (LapseRate, error) {
	if v < 0 || 1 < v {
		return 0, ErrInvalidLapseRateValue
	}
	return LapseRate(v), nil
}

type OutcomeDomain struct {
	Response Response `json:"response"`
}

type Response string

const (
	ResponseCorrect   = Response("correct")
	ResponseIncorrect = Response("incorrect")
)

type PriorNormCDF struct {
	Mean           []Probability `json:"mean"`
	SD             []Probability `json:"sd"`
	LowerAsymptote []Probability `json:"lower_asymptote"`
	LapseRate      []Probability `json:"lapse_rate"`
}

type Probability float64

func NewProbability(v float64) (Probability, error) {
	if v < 0 || 1 > v {
		return 0, ErrInvalidProbabilityValue
	}
	return Probability(v), nil
}

type Func string

const (
	FuncNormCDF = Func("norm_cdf")
)

type StimScale string

const (
	StimScaleLinear = StimScale("linear")
	StimScaleLog10  = StimScale("log10")
)

type StimSelectionMethod string

const (
	StimSelectionMethodMinEntropy  = StimSelectionMethod("min_entropy")
	StimSelectionMethodMinNEntropy = StimSelectionMethod("min_n_entropy")
)

type ParamEstimationMethod string

const (
	ParamEstimationMethodMode = ParamEstimationMethod("mode")
	ParamEstimationMethodMean = ParamEstimationMethod("mean")
)

type QuestPlusResultNormCDF struct {
	NumTrials                 NumTrials
	Width                     Width
	Velocity                  Velocity
	Azimuth                   Azimuth
	Altitude                  Altitude
	ActualRotationDirection   RotationDirection
	AnsweredRotationDirection RotationDirection
	Response                  Response
	MeanEstimation            Mean
	SDEstimation              SD
	LowerAsymptoteEstimation  LowerAsymptote
	LapseRateEstimation       LapseRate
}

type RotationDirection string

const (
	RotationDirectionPositive = "positive"
	RotationDirectionNegative = "negative"
)

type ResultMDDCW struct {
	ExperimentMDDCW ExperimentMDDCW
	ResultDetail    []QuestPlusResultNormCDF
	Subject         Subject
	Mean            Mean
	SD              SD
	LowerAsymptote  LowerAsymptote
	LapseRate       LapseRate
}

type Subject struct {
	Sex                    Sex // ISO5218
	Age                    Age
	DeafAndHearingImpaired DeafAndHearingImpaired
}

type Sex string

const (
	SexNotKnown      = "not known"
	SexMale          = "male"
	SexFemale        = "female"
	SexNotApplicable = "not applicable"
)

type Age uint64

func NewAge(v uint64) Age {
	return Age(v)
}

type DeafAndHearingImpaired bool
