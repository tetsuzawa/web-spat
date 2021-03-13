package domain

import "fmt"

type ExperimentId uint64

// NewExperimentId generates ExperimentId.
func NewExperimentId(v uint64) ExperimentId {
	return ExperimentId(v)
}

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
func NewAltitude(v int64) (Altitude, error) {
	if v < -900 || 900 < v {
		return 0, ErrInvalidAltitudeValue
	}
	return Altitude(v), nil
}

// CoordinateVariable is the variable
type CoordinateVariable string

func NewCoordinateVariable(v string) (CoordinateVariable, error) {
	if v == string(CoordinateVariableAzimuth) {
		return CoordinateVariableAzimuth, nil
	} else if v == string(CoordinateVariableAltitude) {
		return CoordinateVariableAltitude, nil
	} else {
		return "", ErrInvalidCoordinateVariable
	}
}

const (
	CoordinateVariableAzimuth  = CoordinateVariable("azimuth")
	CoordinateVariableAltitude = CoordinateVariable("altitude")
)

// MovingSoundConstant is the variable
type MovingSoundConstant string

// NewNewMovingSoundConstant generates MovingSoundConstant.
func NewMovingSoundConstant(v string) (MovingSoundConstant, error) {
	if v == string(MovingSoundConstantWidth) {
		return MovingSoundConstantWidth, nil
	} else if v == string(MovingSoundConstantVelocity) {
		return MovingSoundConstantVelocity, nil
	} else {
		return "", ErrInvalidMovingSoundConstant
	}
}

const (
	MovingSoundConstantWidth    = MovingSoundConstant("width")
	MovingSoundConstantVelocity = MovingSoundConstant("velocity")
)

// MovingSoundConstantValue is the value of MovingSoundConstant
type MovingSoundConstantValue uint64

// NewNewMovingSoundConstantValue generates MovingSoundConstantValue.
func NewMovingSoundConstantValue(v uint64) MovingSoundConstantValue {
	return MovingSoundConstantValue(v)
}

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

type NumTrials uint64

// NewNumTrials generates NewNumTrials.
func NewNumTrials(v uint64) NumTrials {
	return NumTrials(v)
}

type ExperimentMDD struct {
	Id                        ExperimentId
	QuestPlusParameterNormCDF QuestPlusParameterNormCDF
	Name                      ExperimentName
	Description               ExperimentDescription
	Azimuth                   Azimuth  // 0[10^-1 deg] is the front, 900[10^-1 deg] is the right side
	Altitude                  Altitude // 0[10^-1 deg] is the front, 900[10^-1 deg] is the zenith, -900[10^-1 deg] is the nadir
	CoordinateVariable        CoordinateVariable
	MovingSoundConstant       MovingSoundConstant
	MovingSoundConstantValue  MovingSoundConstantValue
	NumTrials                 NumTrials
}

func NewExperimentMDD(
	id uint64,
	qpParam QuestPlusParameterNormCDF,
	name string,
	description string,
	azimuth uint64,
	altitude int64,
	coordinateVariable string,
	movingSoundConstant string,
	movingSoundConstantValue,
	numTrials uint64,
) (*ExperimentMDD, error) {
	nameDO, err := NewExperimentName(name)
	if err != nil {
		return nil, fmt.Errorf("failed to create property of ExperimentMDD: Name -> %w", err)
	}
	descriptionDO, err := NewExperimentDescription(description)
	if err != nil {
		return nil, fmt.Errorf("failed to create property of ExperimentMDD: Description -> %w", err)
	}
	azimuthDO, err := NewAzimuth(azimuth)
	if err != nil {
		return nil, fmt.Errorf("failed to create property of ExperimentMDD: Azimuth -> %w", err)
	}
	altitudeDO, err := NewAltitude(altitude)
	if err != nil {
		return nil, fmt.Errorf("failed to create property of ExperimentMDD: Altitude -> %w", err)
	}
	coordinateVariableDO, err := NewCoordinateVariable(coordinateVariable)
	if err != nil {
		return nil, fmt.Errorf("failed to create property of ExperimentMDD: CoordinateVariable -> %w", err)
	}
	movingSoundConstantDO, err := NewMovingSoundConstant(movingSoundConstant)
	if err != nil {
		return nil, fmt.Errorf("failed to create property of ExperimentMDD: MovingSoundConstant -> %w", err)
	}

	return &ExperimentMDD{
		Id:                        NewExperimentId(id),
		QuestPlusParameterNormCDF: qpParam,
		Name:                      nameDO,
		Description:               descriptionDO,
		Azimuth:                   azimuthDO,
		Altitude:                  altitudeDO,
		CoordinateVariable:        coordinateVariableDO,
		MovingSoundConstant:       movingSoundConstantDO,
		MovingSoundConstantValue:  NewMovingSoundConstantValue(movingSoundConstantValue),
		NumTrials:                 NewNumTrials(numTrials),
	}, nil
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

func NewQuestPlusParameterNormCDF(
	stimDomain StimDomainNormCDF,
	paramDomain ParamDomainNormCDF,
	outcomeDomain OutcomeDomain,
	prior PriorNormCDF,
	fn Func,
	stimScale StimScale,
	stimSelectionMethod StimSelectionMethod,
	paramEstimationMethod ParamEstimationMethod,
) (*QuestPlusParameterNormCDF, error) {
	if len(paramDomain.Mean) != len(prior.Mean) {
		return nil, fmt.Errorf("the Mean of param domain or prior is invalid -> %w", ErrLengthNotMatch)
	}
	if len(paramDomain.SD) != len(prior.SD) {
		return nil, fmt.Errorf("the SD of param domain or prior is invalid -> %w", ErrLengthNotMatch)
	}
	if len(paramDomain.LowerAsymptote) != len(prior.LowerAsymptote) {
		return nil, fmt.Errorf("the LowerAsymptote of param domain or prior is invalid -> %w", ErrLengthNotMatch)
	}
	if len(paramDomain.LapseRate) != len(prior.LapseRate) {
		return nil, fmt.Errorf("the LapseRate of param domain or prior is invalid -> %w", ErrLengthNotMatch)
	}
	return &QuestPlusParameterNormCDF{
		StimDomain:            stimDomain,
		ParamDomain:           paramDomain,
		OutcomeDomain:         outcomeDomain,
		Prior:                 prior,
		Func:                  fn,
		StimScale:             stimScale,
		StimSelectionMethod:   stimSelectionMethod,
		ParamEstimationMethod: paramEstimationMethod,
	}, nil
}

type StimDomainNormCDF struct {
	Intensity []float64 `json:"intensity"`
}

func NewStimDomainNormCDF(intensity []float64) *StimDomainNormCDF {
	return &StimDomainNormCDF{Intensity: intensity}
}

type ParamDomainNormCDF struct {
	Mean           []Mean           `json:"mean"`
	SD             []SD             `json:"sd"`
	LowerAsymptote []LowerAsymptote `json:"lower_asymptote"`
	LapseRate      []LapseRate      `json:"lapse_rate"`
}

func NewParamDomainNormCDF(mean, sd, lowerAsymptote, lapseRate []float64) (*ParamDomainNormCDF, error) {
	var err error
	meanDO := make([]Mean, len(mean))
	for i, v := range mean {
		meanDO[i] = NewMean(v)
	}
	sdDO := make([]SD, len(sd))
	for i, v := range sd {
		sdDO[i], err = NewSD(v)
		if err != nil {
			return nil, fmt.Errorf("failed to create ParamDomainNormCDF -> %w", err)
		}
	}
	lowerAsymptoteDO := make([]LowerAsymptote, len(lowerAsymptote))
	for i, v := range lowerAsymptote {
		lowerAsymptoteDO[i], err = NewLowerAsymptote(v)
		if err != nil {
			return nil, fmt.Errorf("failed to create ParamDomainNormCDF -> %w", err)
		}
	}
	LapseRateDO := make([]LapseRate, len(lapseRate))
	for i, v := range lapseRate {
		LapseRateDO[i], err = NewLapseRate(v)
		if err != nil {
			return nil, fmt.Errorf("failed to create ParamDomainNormCDF -> %w", err)
		}
	}
	return &ParamDomainNormCDF{
		Mean:           meanDO,
		SD:             sdDO,
		LowerAsymptote: lowerAsymptoteDO,
		LapseRate:      LapseRateDO,
	}, nil
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
	Response []Response `json:"response"`
}

func NewOutcomeDomain() *OutcomeDomain {
	return &OutcomeDomain{Response: []Response{ResponseCorrect, ResponseIncorrect}}
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

func NewPriorNormCDF(mean, sd, lowerAsymptote, lapseRate []float64) (*PriorNormCDF, error) {
	var err error
	meanDO := make([]Probability, len(mean))
	for i, v := range mean {
		meanDO[i], err = NewProbability(v)
		if err != nil {
			return nil, fmt.Errorf("failed to create PriorNormCDF -> %w", err)
		}
	}
	sdDO := make([]Probability, len(sd))
	for i, v := range sd {
		sdDO[i], err = NewProbability(v)
		if err != nil {
			return nil, fmt.Errorf("failed to create PriorNormCDF -> %w", err)
		}
	}
	lowerAsymptoteDO := make([]Probability, len(lowerAsymptote))
	for i, v := range lowerAsymptote {
		lowerAsymptoteDO[i], err = NewProbability(v)
		if err != nil {
			return nil, fmt.Errorf("failed to create PriorNormCDF -> %w", err)
		}
	}
	lapseRateDO := make([]Probability, len(lapseRate))
	for i, v := range lapseRate {
		lapseRateDO[i], err = NewProbability(v)
		if err != nil {
			return nil, fmt.Errorf("failed to create PriorNormCDF -> %w", err)
		}
	}
	return &PriorNormCDF{
		Mean:           meanDO,
		SD:             sdDO,
		LowerAsymptote: lowerAsymptoteDO,
		LapseRate:      lapseRateDO,
	}, nil
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

func NewFunc(v string) (Func, error) {
	if v == string(FuncNormCDF) {
		return FuncNormCDF, nil
	}
	return "", ErrInvalidFunc
}

type StimScale string

func NewStimScale(v string) (StimScale, error) {
	if v == string(StimScaleLinear) {
		return StimScaleLinear, nil
	} else if v == string(StimScaleLog10) {
		return StimScaleLog10, nil
	}
	return "", ErrInvalidStimScale
}

const (
	StimScaleLinear = StimScale("linear")
	StimScaleLog10  = StimScale("log10")
)

type StimSelectionMethod string

func NewStimSelectionMethod(v string) (StimSelectionMethod, error) {
	if v == string(StimSelectionMethodMinEntropy) {
		return StimSelectionMethodMinEntropy, nil
	} else if v == string(StimSelectionMethodMinNEntropy) {
		return StimSelectionMethodMinNEntropy, nil
	}
	return "", ErrInvalidStimSelectionMethod
}

const (
	StimSelectionMethodMinEntropy  = StimSelectionMethod("min_entropy")
	StimSelectionMethodMinNEntropy = StimSelectionMethod("min_n_entropy")
)

type ParamEstimationMethod string

func NewParamEstimationMethod(v string) (ParamEstimationMethod, error) {
	if v == string(ParamEstimationMethodMode) {
		return ParamEstimationMethodMode, nil
	} else if v == string(ParamEstimationMethodMean) {
		return ParamEstimationMethodMean, nil
	}
	return "", ErrInvalidParamEstimationMethod
}

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

type ResultMDD struct {
	ExperimentMDD  ExperimentMDD
	ResultDetail   []QuestPlusResultNormCDF
	Subject        Subject
	Mean           Mean
	SD             SD
	LowerAsymptote LowerAsymptote
	LapseRate      LapseRate
}

type Subject struct {
	Sex                    Sex // ISO5218
	Age                    Age
	DeafAndHearingImpaired DeafAndHearingImpaired
}

//'0: not known, 1: mail, 2: female, 9: not applicable. ISO 5218'
type Sex string

const (
	SexNotKnown      = "0"
	SexMale          = "1"
	SexFemale        = "2"
	SexNotApplicable = "9"
)

type Age uint64

func NewAge(v uint64) Age {
	return Age(v)
}

type DeafAndHearingImpaired bool
