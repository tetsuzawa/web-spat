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
