package domain

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

type QuestPlusParameterNormCDFData struct {
	StimDomain            StimDomainNormCDFData  `db:"stim_domain"`
	ParamDomain           ParamDomainNormCDFData `db:"param_domain"`
	OutcomeDomain         OutcomeDomainData      `db:"outcome_domain"`
	Prior                 PriorNormCDFData       `db:"prior"`
	Func                  string                 `db:"func"`
	StimScale             string                 `db:"stim_scale"`
	StimSelectionMethod   string                 `db:"stim_selection_method"`
	ParamEstimationMethod string                 `db:"param_estimation_method"`
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

type StimDomainNormCDFData struct {
	Intensity []float64 `db:"intensity"`
}

func NewStimDomainNormCDFData(v *StimDomainNormCDF) *StimDomainNormCDFData {
	return &StimDomainNormCDFData{
		Intensity: v.Intensity,
	}
}

type ParamDomainNormCDFData struct {
	Mean           []float64 `db:"mean"`
	SD             []float64 `db:"sd"`
	LowerAsymptote []float64 `db:"lower_asymptote"`
	LapseRate      []float64 `db:"lapse_rate"`
}

func NewParamDomainNormCDFData(v *ParamDomainNormCDF) *ParamDomainNormCDFData {
	mean := make([]float64, len(v.Mean))
	for _, x := range v.Mean {
		mean = append(mean, float64(x))
	}
	sd := make([]float64, len(v.SD))
	for _, x := range v.SD {
		sd = append(sd, float64(x))
	}
	lowerAsymptote := make([]float64, len(v.LowerAsymptote))
	for _, x := range v.LowerAsymptote {
		lowerAsymptote = append(lowerAsymptote, float64(x))
	}
	lapseRate := make([]float64, len(v.LapseRate))
	for _, x := range v.LapseRate {
		lapseRate = append(lapseRate, float64(x))
	}

	return &ParamDomainNormCDFData{
		Mean:           mean,
		SD:             sd,
		LowerAsymptote: lowerAsymptote,
		LapseRate:      lapseRate,
	}
}

type OutcomeDomainData struct {
	Response string `db:"response"`
}

func NewOutcomeDomainData(v *OutcomeDomain) *OutcomeDomainData {
	return &OutcomeDomainData{
		Response: string(v.Response),
	}
}

type PriorNormCDFData struct {
	Mean           []float64 `db:"mean"`
	SD             []float64 `db:"sd"`
	LowerAsymptote []float64 `db:"lower_asymptote"`
	LapseRate      []float64 `db:"lapse_rate"`
}

func NewPriorNormCDFData(v *PriorNormCDF) *PriorNormCDFData {
	mean := make([]float64, len(v.Mean))
	for _, x := range v.Mean {
		mean = append(mean, float64(x))
	}
	sd := make([]float64, len(v.SD))
	for _, x := range v.SD {
		sd = append(sd, float64(x))
	}
	lowerAsymptote := make([]float64, len(v.LowerAsymptote))
	for _, x := range v.LowerAsymptote {
		lowerAsymptote = append(lowerAsymptote, float64(x))
	}
	lapseRate := make([]float64, len(v.LapseRate))
	for _, x := range v.LapseRate {
		lapseRate = append(lapseRate, float64(x))
	}

	return &PriorNormCDFData{
		Mean:           mean,
		SD:             sd,
		LowerAsymptote: lowerAsymptote,
		LapseRate:      lapseRate,
	}
}

type QuestPlusResultNormCDFData struct {
	NumTrials                 uint64  `db:"num_trials"`
	Width                     uint64  `db:"width"`
	Velocity                  uint64  `db:"velocity"`
	Azimuth                   uint64  `db:"azimuth"`
	Altitude                  int64   `db:"altitude"`
	ActualRotationDirection   string  `db:"actual_rotation_direction"`
	AnsweredRotationDirection string  `db:"answered_rotation_direction"`
	Response                  string  `db:"response"`
	MeanEstimation            float64 `db:"mean_estimation"`
	SDEstimation              float64 `db:"sd_estimation"`
	LowerAsymptoteEstimation  float64 `db:"lower_asymptote_estimation"`
	LapseRateEstimation       float64 `db:"lapse_rate_estimation"`
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

type ResultMDDData struct {
	ExperimentMDD  ExperimentMDDData            `db:"experiment_MDD"`
	ResultDetail   []QuestPlusResultNormCDFData `db:"result_detail"`
	Subject        SubjectData                  `db:"subject"`
	Mean           float64                      `db:"mean"`
	Sd             float64                      `db:"sd"`
	LowerAsymptote float64                      `db:"lower_asymptote"`
	LapseRate      float64                      `db:"lapse_rate"`
}

func NewResultMDDData(v *ResultMDD) *ResultMDDData {
	resultDetailData := make([]QuestPlusResultNormCDFData, len(v.ResultDetail))
	for _, x := range v.ResultDetail {
		q := NewQuestPlusResultNormCDFData(&x)
		resultDetailData = append(resultDetailData, *q)
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
