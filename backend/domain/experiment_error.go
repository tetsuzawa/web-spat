package domain

import "fmt"

var (
	ErrTooLongExperimentName = fmt.Errorf("too long experiment name")
	ErrTooLongExperimentDescription = fmt.Errorf("too long experiment description")
	ErrInvalidAzimuthValue = fmt.Errorf("azimuth must be in the range 0~3600 [10^-1 deg]")
	ErrInvalidAltitudeValue = fmt.Errorf("altitude must be in the range -900~900 [10^-1 deg]")
)
