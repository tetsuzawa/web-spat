// Package Openapi provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package Openapi

// EstimatedParameters defines model for EstimatedParameters.
type EstimatedParameters struct {
	LapseRate      float64 `json:"lapse_rate"`
	LowerAsymptote float64 `json:"lower_asymptote"`
	Mean           float64 `json:"mean"`
	Sd             float64 `json:"sd"`
}

// ExperimentMovementDirectionDiscriminationConstantWidth defines model for ExperimentMovementDirectionDiscriminationConstantWidth.
type ExperimentMovementDirectionDiscriminationConstantWidth struct {

	// 0[10^-1 deg] is the front, 900[10^-1 deg] is the zenith, -900[10^-1 deg] is the nadir
	Altitude int `json:"altitude"`

	// 0[10^-1 deg] is the front, 900[10^-1 deg] is the right side
	Azimuth            int    `json:"azimuth"`
	CoordinateVariable string `json:"coordinate_variable"`
	Description        string `json:"description"`
	Id                 int64  `json:"id"`
	Name               string `json:"name"`
	NumTrials          int    `json:"num_trials"`
	Status             string `json:"status"`
	VelocityRangeLower int    `json:"velocity_range_lower"`
	VelocityRangeStep  int    `json:"velocity_range_step"`
	VelocityRangeUpper int    `json:"velocity_range_upper"`

	// width of moving sound [10^-1 deg]
	Width int `json:"width"`
}

// Experiments defines model for Experiments.
type Experiments []interface{}

// ResultMovementDirectionDiscriminationConstantWidth defines model for ResultMovementDirectionDiscriminationConstantWidth.
type ResultMovementDirectionDiscriminationConstantWidth struct {
	EstimatedParameters EstimatedParameters `json:"estimated_parameters"`
	ExperimentId        int64               `json:"experiment_id"`
	Subject             Subject             `json:"subject"`
}

// Subject defines model for Subject.
type Subject struct {
	Age                    int  `json:"age"`
	DeafAndHearingImpaired bool `json:"deaf_and_hearing_impaired"`

	// '0: not known, 1: mail, 2: female, 9: not applicable. ISO 5218'
	Sex string `json:"sex"`
}

// RegisterResultOfExperimentIdJSONBody defines parameters for RegisterResultOfExperimentId.
type RegisterResultOfExperimentIdJSONBody ResultMovementDirectionDiscriminationConstantWidth

// RegisterResultOfExperimentIdJSONRequestBody defines body for RegisterResultOfExperimentId for application/json ContentType.
type RegisterResultOfExperimentIdJSONRequestBody RegisterResultOfExperimentIdJSONBody
