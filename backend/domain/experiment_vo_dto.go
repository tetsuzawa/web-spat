package domain

import "fmt"

type ExperimentIdData uint64

func NewExperimentIDData(v ExperimentId) ExperimentIdData {
	return ExperimentIdData(v)
}

type ExperimentMDDData struct {
	Id                        uint64                        `db:"id"`
	QuestPlusParameterNormCDF QuestPlusParameterNormCDFData `db:"quest_plus_parameter_norm_cdf"`
	Name                      string                        `db:"name"`
	Description               string                        `db:"description"`
	Azimuth                   uint64                        `db:"azimuth"`
	Altitude                  int64                         `db:"altitude"`
	CoordinateVariable        string                        `db:"coordinate_variable"`
	MovingSoundConstant       string                        `db:"moving_sound_constant"`
	MovingSoundConstantValue  uint64                        `db:"moving_sound_constant_value"`
	NumTrials                 uint64                        `db:"num_trials"`
}

func NewExperimentMDDData(v *ExperimentMDD) *ExperimentMDDData {
	return &ExperimentMDDData{
		Id:                        uint64(v.Id),
		QuestPlusParameterNormCDF: *NewQuestPlusParameterNormCDFData(&v.QuestPlusParameterNormCDF),
		Name:                      string(v.Name),
		Description:               string(v.Description),
		Azimuth:                   uint64(v.Azimuth),
		Altitude:                  int64(v.Altitude),
		CoordinateVariable:        string(v.CoordinateVariable),
		MovingSoundConstant:       string(v.MovingSoundConstant),
		MovingSoundConstantValue:  uint64(v.MovingSoundConstantValue),
		NumTrials:                 uint64(v.NumTrials),
	}
}

func NewExperimentMDDFromData(v *ExperimentMDDData) (*ExperimentMDD, error) {
	questPlusParameterNormCDFVO, err := NewQuestPlusParameterNormCDFFromData(&v.QuestPlusParameterNormCDF)
	if err != nil {
		return nil, fmt.Errorf("failed to create ExperimentMDD -> %w", err)
	}
	experimentNameDO, err := NewExperimentName(v.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to create ExperimentMDD -> %w", err)
	}
	experimentDescriptionDO, err := NewExperimentDescription(v.Description)
	if err != nil {
		return nil, fmt.Errorf("failed to create ExperimentMDD -> %w", err)
	}
	experimentAzimuthDO, err := NewAzimuth(v.Azimuth)
	if err != nil {
		return nil, fmt.Errorf("failed to create ExperimentMDD -> %w", err)
	}
	experimentAltitudeDO, err := NewAltitude(v.Altitude)
	if err != nil {
		return nil, fmt.Errorf("failed to create ExperimentMDD -> %w", err)
	}
	coordinateVariableDO, err := NewCoordinateVariable(v.CoordinateVariable)
	if err != nil {
		return nil, fmt.Errorf("failed to create ExperimentMDD -> %w", err)
	}
	movingSoundConstantDO, err := NewMovingSoundConstant(v.CoordinateVariable)
	if err != nil {
		return nil, fmt.Errorf("failed to create ExperimentMDD -> %w", err)
	}

	return &ExperimentMDD{
		Id:                        NewExperimentId(v.Id),
		QuestPlusParameterNormCDF: *questPlusParameterNormCDFVO,
		Name:                      experimentNameDO,
		Description:               experimentDescriptionDO,
		Azimuth:                   experimentAzimuthDO,
		Altitude:                  experimentAltitudeDO,
		CoordinateVariable:        coordinateVariableDO,
		MovingSoundConstant:       movingSoundConstantDO,
		MovingSoundConstantValue:  NewMovingSoundConstantValue(v.MovingSoundConstantValue),
		NumTrials:                 NewNumTrials(v.NumTrials),
	}, nil
}

type QuestPlusParameterNormCDFData struct {
	StimDomain            StimDomainNormCDFData  `json:"stim_domain" db:"stim_domain"`
	ParamDomain           ParamDomainNormCDFData `json:"param_domain" db:"param_domain"`
	OutcomeDomain         OutcomeDomainData      `json:"outcome_domain" db:"outcome_domain"`
	Prior                 PriorNormCDFData       `json:"prior" db:"prior"`
	Func                  string                 `json:"func" db:"func"`
	StimScale             string                 `json:"stim_scale" db:"stim_scale"`
	StimSelectionMethod   string                 `json:"stim_selection_method" db:"stim_selection_method"`
	ParamEstimationMethod string                 `json:"param_estimation_method" db:"param_estimation_method"`
}

func NewQuestPlusParameterNormCDFData(v *QuestPlusParameterNormCDF) *QuestPlusParameterNormCDFData {
	return &QuestPlusParameterNormCDFData{
		StimDomain:            *NewStimDomainNormCDFData(&v.StimDomain),
		ParamDomain:           *NewParamDomainNormCDFData(&v.ParamDomain),
		OutcomeDomain:         *NewOutcomeDomainData(&v.OutcomeDomain),
		Prior:                 *NewPriorNormCDFData(&v.Prior),
		Func:                  string(v.Func),
		StimScale:             string(v.StimScale),
		StimSelectionMethod:   string(v.StimSelectionMethod),
		ParamEstimationMethod: string(v.ParamEstimationMethod),
	}
}

func NewQuestPlusParameterNormCDFFromData(v *QuestPlusParameterNormCDFData) (*QuestPlusParameterNormCDF, error) {
	paramDomainDO, err := NewParamDomainNormCDFFromData(&v.ParamDomain)
	if err != nil {
		return nil, fmt.Errorf("failed to create QuestPlusParameterNormCDFData -> %w", err)
	}
	priorNormCDFDO, err := NewPriorNormCDFFromData(&v.Prior)
	if err != nil {
		return nil, fmt.Errorf("failed to create QuestPlusParameterNormCDFData -> %w", err)
	}
	funcDO, err := NewFunc(v.Func)
	if err != nil {
		return nil, fmt.Errorf("failed to create QuestPlusParameterNormCDFData -> %w", err)
	}
	stimScaleDO, err := NewStimScale(v.Func)
	if err != nil {
		return nil, fmt.Errorf("failed to create QuestPlusParameterNormCDFData -> %w", err)
	}
	stimSelectionMethodDO, err := NewStimSelectionMethod(v.Func)
	if err != nil {
		return nil, fmt.Errorf("failed to create QuestPlusParameterNormCDFData -> %w", err)
	}
	paramEstimationMethodDO, err := NewParamEstimationMethod(v.Func)
	if err != nil {
		return nil, fmt.Errorf("failed to create QuestPlusParameterNormCDFData -> %w", err)
	}
	return &QuestPlusParameterNormCDF{
		StimDomain:            *NewStimDomainNormCDF(v.StimDomain.Intensity),
		ParamDomain:           *paramDomainDO,
		OutcomeDomain:         *NewOutcomeDomain(),
		Prior:                 *priorNormCDFDO,
		Func:                  funcDO,
		StimScale:             stimScaleDO,
		StimSelectionMethod:   stimSelectionMethodDO,
		ParamEstimationMethod: paramEstimationMethodDO,
	}, nil
}

type StimDomainNormCDFData struct {
	Intensity []float64 `json:"intensity" db:"intensity"`
}

func NewStimDomainNormCDFData(v *StimDomainNormCDF) *StimDomainNormCDFData {
	return &StimDomainNormCDFData{
		Intensity: v.Intensity,
	}
}

type ParamDomainNormCDFData struct {
	Mean           []float64 `json:"mean" db:"mean"`
	SD             []float64 `json:"sd" db:"sd"`
	LowerAsymptote []float64 `json:"lower_asymptote" db:"lower_asymptote"`
	LapseRate      []float64 `json:"lapse_rate" db:"lapse_rate"`
}

func NewParamDomainNormCDFData(v *ParamDomainNormCDF) *ParamDomainNormCDFData {
	mean := make([]float64, len(v.Mean))
	for i, x := range v.Mean {
		mean[i] = float64(x)
	}
	sd := make([]float64, len(v.SD))
	for i, x := range v.SD {
		sd[i] = float64(x)
	}
	lowerAsymptote := make([]float64, len(v.LowerAsymptote))
	for i, x := range v.LowerAsymptote {
		lowerAsymptote[i] = float64(x)
	}
	lapseRate := make([]float64, len(v.LapseRate))
	for i, x := range v.LapseRate {
		lapseRate[i] = float64(x)
	}

	return &ParamDomainNormCDFData{
		Mean:           mean,
		SD:             sd,
		LowerAsymptote: lowerAsymptote,
		LapseRate:      lapseRate,
	}
}

func NewParamDomainNormCDFFromData(v *ParamDomainNormCDFData) (*ParamDomainNormCDF, error) {
	return NewParamDomainNormCDF(v.Mean, v.SD, v.LowerAsymptote, v.LapseRate)
}

type OutcomeDomainData struct {
	Response []string `json:"response" db:"response"`
}

func NewOutcomeDomainData(v *OutcomeDomain) *OutcomeDomainData {
	response := make([]string, len(v.Response))
	for i, x := range v.Response {
		response[i] = string(x)
	}
	return &OutcomeDomainData{
		Response: response,
	}
}

type PriorNormCDFData struct {
	Mean           []float64 `json:"mean" db:"mean"`
	SD             []float64 `json:"sd" db:"sd"`
	LowerAsymptote []float64 `json:"lower_asymptote" db:"lower_asymptote"`
	LapseRate      []float64 `json:"lapse_rate" db:"lapse_rate"`
}

func NewPriorNormCDFData(v *PriorNormCDF) *PriorNormCDFData {
	mean := make([]float64, len(v.Mean))
	for i, x := range v.Mean {
		mean[i] = float64(x)
	}
	sd := make([]float64, len(v.SD))
	for i, x := range v.SD {
		sd[i] = float64(x)
	}
	lowerAsymptote := make([]float64, len(v.LowerAsymptote))
	for i, x := range v.LowerAsymptote {
		lowerAsymptote[i] = float64(x)
	}
	lapseRate := make([]float64, len(v.LapseRate))
	for i, x := range v.LapseRate {
		lapseRate[i] = float64(x)
	}
	return &PriorNormCDFData{
		Mean:           mean,
		SD:             sd,
		LowerAsymptote: lowerAsymptote,
		LapseRate:      lapseRate,
	}
}

func NewPriorNormCDFFromData(v *PriorNormCDFData) (*PriorNormCDF, error) {
	return NewPriorNormCDF(v.Mean, v.SD, v.LowerAsymptote, v.LapseRate)
}

type QuestPlusResultNormCDFData struct {
	NumTrials                 uint64  `json:"num_trials" db:"num_trials"`
	Width                     uint64  `json:"width" db:"width"`
	Velocity                  uint64  `json:"velocity" db:"velocity"`
	Azimuth                   uint64  `json:"azimuth" db:"azimuth"`
	Altitude                  int64   `json:"altitude" db:"altitude"`
	ActualRotationDirection   string  `json:"actual_rotation_direction" db:"actual_rotation_direction"`
	AnsweredRotationDirection string  `json:"answered_rotation_direction" db:"answered_rotation_direction"`
	Response                  string  `json:"response" db:"response"`
	MeanEstimation            float64 `json:"mean_estimation" db:"mean_estimation"`
	SDEstimation              float64 `json:"sd_estimation" db:"sd_estimation"`
	LowerAsymptoteEstimation  float64 `json:"lower_asymptote_estimation" db:"lower_asymptote_estimation"`
	LapseRateEstimation       float64 `json:"lapse_rate_estimation" db:"lapse_rate_estimation"`
}

func NewQuestPlusResultNormCDFData(v *QuestPlusResultNormCDF) *QuestPlusResultNormCDFData {
	return &QuestPlusResultNormCDFData{
		NumTrials:                 uint64(v.NumTrials),
		Width:                     uint64(v.Width),
		Velocity:                  uint64(v.Velocity),
		Azimuth:                   uint64(v.Azimuth),
		Altitude:                  int64(v.Altitude),
		ActualRotationDirection:   string(v.ActualRotationDirection),
		AnsweredRotationDirection: string(v.AnsweredRotationDirection),
		Response:                  string(v.Response),
		MeanEstimation:            float64(v.MeanEstimation),
		SDEstimation:              float64(v.SDEstimation),
		LowerAsymptoteEstimation:  float64(v.LowerAsymptoteEstimation),
		LapseRateEstimation:       float64(v.LapseRateEstimation),
	}
}

func NewQuestPlusResultNormCDFFromData(v *QuestPlusResultNormCDFData) (*QuestPlusResultNormCDF, error) {
	return NewQuestPlusResultNormCDF(
		v.NumTrials,
		v.Width,
		v.Velocity,
		v.Azimuth,
		v.Altitude,
		v.ActualRotationDirection,
		v.AnsweredRotationDirection,
		v.Response,
		v.MeanEstimation,
		v.SDEstimation,
		v.LowerAsymptoteEstimation,
		v.LapseRateEstimation,
	)
}

type ResultMDDData struct {
	ExperimentMDD  ExperimentMDDData            `json:"experiment_mdd" db:"experiment_mdd"`
	ResultDetail   []QuestPlusResultNormCDFData `json:"result_detail" db:"result_detail"`
	Subject        SubjectData                  `json:"subject" db:"subject"`
	Mean           float64                      `json:"mean" db:"mean"`
	Sd             float64                      `json:"sd" db:"sd"`
	LowerAsymptote float64                      `json:"lower_asymptote" db:"lower_asymptote"`
	LapseRate      float64                      `json:"lapse_rate" db:"lapse_rate"`
}

func NewResultMDDData(v *ResultMDD) *ResultMDDData {
	resultDetailData := make([]QuestPlusResultNormCDFData, len(v.ResultDetail))
	for i, x := range v.ResultDetail {
		q := NewQuestPlusResultNormCDFData(&x)
		resultDetailData[i] = *q
	}
	return &ResultMDDData{
		ExperimentMDD:  *NewExperimentMDDData(&v.ExperimentMDD),
		ResultDetail:   resultDetailData,
		Subject:        *NewSubjectData(&v.Subject),
		Mean:           float64(v.Mean),
		Sd:             float64(v.SD),
		LowerAsymptote: float64(v.LowerAsymptote),
		LapseRate:      float64(v.LapseRate),
	}
}

func NewResultMDDFromData(v *ResultMDDData) (*ResultMDD, error) {
	experimentMDDDO, err := NewExperimentMDDFromData(&v.ExperimentMDD)
	if err != nil {
		return nil, fmt.Errorf("failed to create ResultMDD -> %w", err)
	}

	resultDetailDO := make([]QuestPlusResultNormCDF, len(v.ResultDetail))
	for i, x := range v.ResultDetail {
		q, err := NewQuestPlusResultNormCDFFromData(&x)
		if err != nil {
			return nil, fmt.Errorf("failed to create ResultMDD -> %w", err)
		}
		resultDetailDO[i] = *q
	}
	subjectDO, err := NewSubjectFromData(&v.Subject)
	if err != nil {
		return nil, fmt.Errorf("failed to create ResultMDD -> %w", err)
	}
	meanDO := NewMean(v.Mean)
	sdDO, err := NewSD(v.Sd)
	if err != nil {
		return nil, fmt.Errorf("failed to create property of ResultMDD: SD-> %w", err)
	}
	lowerAsymptoteDO, err := NewLowerAsymptote(v.LowerAsymptote)
	if err != nil {
		return nil, fmt.Errorf("failed to create property of ResultMDD: LowerAsymptote -> %w", err)
	}
	lapseRateDO, err := NewLapseRate(v.LapseRate)
	if err != nil {
		return nil, fmt.Errorf("failed to create property of ResultMDD: LapseRate -> %w", err)
	}
	return &ResultMDD{
		ExperimentMDD:  *experimentMDDDO,
		ResultDetail:   resultDetailDO,
		Subject:        *subjectDO,
		Mean:           meanDO,
		SD:             sdDO,
		LowerAsymptote: lowerAsymptoteDO,
		LapseRate:      lapseRateDO,
	}, nil
}

type SubjectData struct {
	Sex                    string `db:"sex"`
	Age                    uint64 `db:"age"`
	DeafAndHearingImpaired bool   `db:"deaf_and_hearing_impaired"`
}

func NewSubjectData(v *Subject) *SubjectData {
	return &SubjectData{
		Sex:                    string(v.Sex),
		Age:                    uint64(v.Age),
		DeafAndHearingImpaired: bool(v.DeafAndHearingImpaired),
	}
}

func NewSubjectFromData(v *SubjectData) (*Subject, error) {
	sexDO, err := NewSex(v.Sex)
	if err != nil {
		return nil, fmt.Errorf("failed to create Subect -> %w", err)
	}
	return &Subject{
		Sex:                    sexDO,
		Age:                    NewAge(v.Age),
		DeafAndHearingImpaired: NewDeafAndHearingImpaired(v.DeafAndHearingImpaired),
	}, nil
}
