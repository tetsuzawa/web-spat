package domain

import "fmt"

var (
	ErrTooLongExperimentName        = fmt.Errorf("too long experiment name")
	ErrTooLongExperimentDescription = fmt.Errorf("too long experiment description")
	ErrInvalidAzimuthValue          = fmt.Errorf("azimuth must be in the range 0~3600 [10^-1 deg]")
	ErrInvalidAltitudeValue         = fmt.Errorf("altitude must be in the range -900~900 [10^-1 deg]")
	ErrInvalidSDValue               = fmt.Errorf("SD must be a positive value")
	ErrInvalidProbabilityValue      = fmt.Errorf("probability must be in the range 0~1.0")
	ErrInvalidLowerAsymptoteValue   = fmt.Errorf("lower asymptote must be in the range 0~1.0")
	ErrInvalidLapseRateValue        = fmt.Errorf("lapse rate must be in the range 0~1.0")
	ErrInvalidCoordinateVariable    = fmt.Errorf("lapse rate must be `azimuth` or `altitude`")
	ErrLengthNotMatch               = fmt.Errorf("the length of the values does not match")
	ErrInvalidFunc                  = fmt.Errorf("invalid func")
	ErrInvalidStimScale             = fmt.Errorf("invalid stim scale")
	ErrInvalidStimSelectionMethod   = fmt.Errorf("invalid stim selection method")
	ErrInvalidParamEstimationMethod = fmt.Errorf("invalid param estimation method")
	ErrInvalidMovingSoundConstant   = fmt.Errorf("invalid param moving sound constant")
)
