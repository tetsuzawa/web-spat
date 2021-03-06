// Package openapi provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package openapi

// Age defines model for Age.
type Age int64

// Altitude defines model for Altitude.
type Altitude int64

// Azimuth defines model for Azimuth.
type Azimuth int64

// CoordinateVariable defines model for CoordinateVariable.
type CoordinateVariable string

// List of CoordinateVariable
const (
	CoordinateVariable_altitude CoordinateVariable = "altitude"
	CoordinateVariable_azimuth  CoordinateVariable = "azimuth"
)

// DeafAndHearingImpaired defines model for DeafAndHearingImpaired.
type DeafAndHearingImpaired bool

// ExperimentDescription defines model for ExperimentDescription.
type ExperimentDescription string

// ExperimentId defines model for ExperimentId.
type ExperimentId int64

// ExperimentMDD defines model for ExperimentMDD.
type ExperimentMDD struct {

	// 0[10^-1 deg] is the front, 900[10^-1 deg] is the zenith, -900[10^-1 deg] is the nadir
	Altitude Altitude `json:"altitude"`

	// 0[10^-1 deg] is the front, 900[10^-1 deg] is the right side
	Azimuth             Azimuth               `json:"azimuth"`
	CoordinateVariable  CoordinateVariable    `json:"coordinate_variable"`
	Description         ExperimentDescription `json:"description"`
	Id                  ExperimentId          `json:"id"`
	MovingSoundConstant MovingSoundConstant   `json:"moving_sound_constant"`

	// width or velocity
	MovingSoundConstantValue  int64                     `json:"moving_sound_constant_value"`
	Name                      ExperimentName            `json:"name"`
	NumTrials                 NumTrials                 `json:"num_trials"`
	QuestPlusParameterNormCdf QuestPlusParameterNormCDF `json:"quest_plus_parameter_norm_cdf"`
}

// ExperimentName defines model for ExperimentName.
type ExperimentName string

// ExperimentsMDD defines model for ExperimentsMDD.
type ExperimentsMDD []ExperimentMDD

// LapseRate defines model for LapseRate.
type LapseRate float64

// LowerAsymptote defines model for LowerAsymptote.
type LowerAsymptote float64

// Mean defines model for Mean.
type Mean float64

// MovingSoundConstant defines model for MovingSoundConstant.
type MovingSoundConstant string

// List of MovingSoundConstant
const (
	MovingSoundConstant_velocity MovingSoundConstant = "velocity"
	MovingSoundConstant_width    MovingSoundConstant = "width"
)

// NumTrials defines model for NumTrials.
type NumTrials int64

// OutcomeDomain defines model for OutcomeDomain.
type OutcomeDomain struct {
	Response string `json:"response"`
}

// ParamDomainNormCDF defines model for ParamDomainNormCDF.
type ParamDomainNormCDF struct {
	LapseRate      []LapseRate      `json:"lapse_rate"`
	LowerAsymptote []LowerAsymptote `json:"lower_asymptote"`
	Mean           []Mean           `json:"mean"`
	Sd             []SD             `json:"sd"`
}

// PriorNormCDF defines model for PriorNormCDF.
type PriorNormCDF struct {
	LapseRate      []Probability `json:"lapse_rate"`
	LowerAsymptote []Probability `json:"lower_asymptote"`
	Mean           []Probability `json:"mean"`
	Sd             []Probability `json:"sd"`
}

// Probability defines model for Probability.
type Probability float64

// QuestPlusFunc defines model for QuestPlusFunc.
type QuestPlusFunc string

// List of QuestPlusFunc
const (
	QuestPlusFunc_logistic QuestPlusFunc = "logistic"
	QuestPlusFunc_norm_cdf QuestPlusFunc = "norm_cdf"
	QuestPlusFunc_weibull  QuestPlusFunc = "weibull"
)

// QuestPlusParamEstimationMethod defines model for QuestPlusParamEstimationMethod.
type QuestPlusParamEstimationMethod string

// List of QuestPlusParamEstimationMethod
const (
	QuestPlusParamEstimationMethod_mean QuestPlusParamEstimationMethod = "mean"
	QuestPlusParamEstimationMethod_mode QuestPlusParamEstimationMethod = "mode"
)

// QuestPlusParameterNormCDF defines model for QuestPlusParameterNormCDF.
type QuestPlusParameterNormCDF struct {
	Func                  QuestPlusFunc                  `json:"func"`
	OutcomeDomain         OutcomeDomain                  `json:"outcome_domain"`
	ParamDomain           ParamDomainNormCDF             `json:"param_domain"`
	ParamEstimationMethod QuestPlusParamEstimationMethod `json:"param_estimation_method"`
	Prior                 PriorNormCDF                   `json:"prior"`
	StimDomain            StimDomainNormCDF              `json:"stim_domain"`
	StimScale             QuestPlusStimScale             `json:"stim_scale"`
	StimSelectionMethod   QuestPlusStimSelectionMethod   `json:"stim_selection_method"`
}

// QuestPlusResultNormCDF defines model for QuestPlusResultNormCDF.
type QuestPlusResultNormCDF struct {
	ActualRotationDirection RotationDirection `json:"actual_rotation_direction"`

	// 0[10^-1 deg] is the front, 900[10^-1 deg] is the zenith, -900[10^-1 deg] is the nadir
	Altitude                  Altitude          `json:"altitude"`
	AnsweredRotationDirection RotationDirection `json:"answered_rotation_direction"`

	// 0[10^-1 deg] is the front, 900[10^-1 deg] is the right side
	Azimuth                  Azimuth        `json:"azimuth"`
	LapseRateEstimation      LapseRate      `json:"lapse_rate_estimation"`
	LowerAsymptoteEstimation LowerAsymptote `json:"lower_asymptote_estimation"`
	MeanEstimation           Mean           `json:"mean_estimation"`
	NumTrials                NumTrials      `json:"num_trials"`
	Response                 Response       `json:"response"`
	SdEstimation             SD             `json:"sd_estimation"`
	Velocity                 Velocity       `json:"velocity"`
	Width                    Width          `json:"width"`
}

// QuestPlusStimScale defines model for QuestPlusStimScale.
type QuestPlusStimScale string

// List of QuestPlusStimScale
const (
	QuestPlusStimScale_linear QuestPlusStimScale = "linear"
	QuestPlusStimScale_log10  QuestPlusStimScale = "log10"
)

// QuestPlusStimSelectionMethod defines model for QuestPlusStimSelectionMethod.
type QuestPlusStimSelectionMethod string

// List of QuestPlusStimSelectionMethod
const (
	QuestPlusStimSelectionMethod_min_entropy   QuestPlusStimSelectionMethod = "min_entropy"
	QuestPlusStimSelectionMethod_min_n_entropy QuestPlusStimSelectionMethod = "min_n_entropy"
)

// Response defines model for Response.
type Response string

// List of Response
const (
	Response_Correct   Response = "Correct"
	Response_Incorrect Response = "Incorrect"
)

// ResultMDD defines model for ResultMDD.
type ResultMDD struct {

	// Experiment for Movement Direction Discrimination
	ExperimentMdd  ExperimentMDD            `json:"experiment_mdd"`
	LapseRate      LapseRate                `json:"lapse_rate"`
	LowerAsymptote LowerAsymptote           `json:"lower_asymptote"`
	Mean           Mean                     `json:"mean"`
	ResultDetail   []QuestPlusResultNormCDF `json:"result_detail"`
	Sd             SD                       `json:"sd"`
	Subject        Subject                  `json:"subject"`
}

// RotationDirection defines model for RotationDirection.
type RotationDirection string

// List of RotationDirection
const (
	RotationDirection_negative RotationDirection = "negative"
	RotationDirection_positive RotationDirection = "positive"
)

// SD defines model for SD.
type SD float64

// Sex defines model for Sex.
type Sex string

// List of Sex
const (
	Sex_female         Sex = "female"
	Sex_male           Sex = "male"
	Sex_not_applicable Sex = "not applicable"
	Sex_not_known      Sex = "not known"
)

// StimDomainNormCDF defines model for StimDomainNormCDF.
type StimDomainNormCDF struct {
	Intensity []float64 `json:"intensity"`
}

// Subject defines model for Subject.
type Subject struct {
	Age                    Age                    `json:"age"`
	DeafAndHearingImpaired DeafAndHearingImpaired `json:"deaf_and_hearing_impaired"`

	// '0: not known, 1: male, 2: female, 9: not applicable. ISO 5218'
	Sex Sex `json:"sex"`
}

// Velocity defines model for Velocity.
type Velocity int64

// Width defines model for Width.
type Width int64

// CreateExperimentMDDJSONBody defines parameters for CreateExperimentMDD.
type CreateExperimentMDDJSONBody ExperimentMDD

// RegisterResultOfExperimentMDDByIdJSONBody defines parameters for RegisterResultOfExperimentMDDById.
type RegisterResultOfExperimentMDDByIdJSONBody ResultMDD

// CreateExperimentMDDJSONRequestBody defines body for CreateExperimentMDD for application/json ContentType.
type CreateExperimentMDDJSONRequestBody CreateExperimentMDDJSONBody

// RegisterResultOfExperimentMDDByIdJSONRequestBody defines body for RegisterResultOfExperimentMDDById for application/json ContentType.
type RegisterResultOfExperimentMDDByIdJSONRequestBody RegisterResultOfExperimentMDDByIdJSONBody

