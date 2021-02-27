package domain

type ExperimentMDDCWData struct {
	QuestPlusParameterNormCDF QuestPlusParameterNormCDFData
	Name                      string `json:"name"`
	Description               string `json:"description"`
	Azimuth                   int    `json:"azimuth"`
	Altitude                  int    `json:"altitude"`
	CoordinateVariable        string `json:"coordinate_variable"`
	Width                     int    `json:"width"`
	VelocityRangeLower        int    `json:"velocity_range_lower"`
	VelocityRangeUpper        int    `json:"velocity_range_upper"`
	VelocityRangeStep         int    `json:"velocity_range_step"`
	NumTrials                 int    `json:"num_trials"`
}

func NewExperimentMDDCWData(v *ExperimentMDDCW) *ExperimentMDDCWData {
	return &ExperimentMDDCWData{
		QuestPlusParameterNormCDF: *NewQuestPlusParameterNormCDFData(&v.QuestPlusParameterNormCDF),
		Name:                      string(v.Name),
		Description:               string(v.Description),
		Azimuth:                   int(v.Azimuth),
		Altitude:                  int(v.Altitude),
		CoordinateVariable:        string(v.CoordinateVariable),
		Width:                     int(v.Width),
		VelocityRangeLower:        int(v.VelocityRangeLower),
		VelocityRangeUpper:        int(v.VelocityRangeUpper),
		VelocityRangeStep:         int(v.VelocityRangeStep),
		NumTrials:                 int(v.NumTrials),
	}
}

type QuestPlusParameterNormCDFData struct {
	StimDomain            StimDomainNormCDFData  `json:"stim_domain"`
	ParamDomain           ParamDomainNormCDFData `json:"param_domain"`
	OutcomeDomain         OutcomeDomainData      `json:"outcome_domain"`
	Prior                 PriorNormCDFData       `json:"prior"`
	Func                  string                 `json:"func"`
	StimScale             string                 `json:"stim_scale"`
	StimSelectionMethod   string                 `json:"stim_selection_method"`
	ParamEstimationMethod string                 `json:"param_estimation_method"`
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
	Intensity []float64 `json:"intensity"`
}

func NewStimDomainNormCDFData(v *StimDomainNormCDF) *StimDomainNormCDFData {
	return &StimDomainNormCDFData{
		Intensity: v.Intensity,
	}
}

type ParamDomainNormCDFData struct {
	Mean           []float64 `json:"mean"`
	SD             []float64 `json:"sd"`
	LowerAsymptote []float64 `json:"lower_asymptote"`
	LapseRate      []float64 `json:"lapse_rate"`
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
	Response string `json:"response"`
}

func NewOutcomeDomainData(v *OutcomeDomain) *OutcomeDomainData {
	return &OutcomeDomainData{
		Response: string(v.Response),
	}
}

type PriorNormCDFData struct {
	Mean           []float64 `json:"mean"`
	SD             []float64 `json:"sd"`
	LowerAsymptote []float64 `json:"lower_asymptote"`
	LapseRate      []float64 `json:"lapse_rate"`
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
	NumTrials                 int     `json:"num_trials"`
	Width                     int     `json:"width"`
	Velocity                  int     `json:"velocity"`
	Azimuth                   int     `json:"azimuth"`
	Altitude                  int     `json:"altitude"`
	ActualRotationDirection   string  `json:"actual_rotation_direction"`
	AnsweredRotationDirection string  `json:"answered_rotation_direction"`
	Response                  string  `json:"response"`
	MeanEstimation            float64 `json:"mean_estimation"`
	SDEstimation              float64 `json:"sd_estimation"`
	LowerAsymptoteEstimation  float64 `json:"lower_asymptote_estimation"`
	LapseRateEstimation       float64 `json:"lapse_rate_estimation"`
}

func NewQuestPlusResultNormCDFData(v *QuestPlusResultNormCDF) *QuestPlusResultNormCDFData {
	return &QuestPlusResultNormCDFData{
		NumTrials:                 int(v.NumTrials),
		Width:                     int(v.Width),
		Velocity:                  int(v.Velocity),
		Azimuth:                   int(v.Azimuth),
		Altitude:                  int(v.Altitude),
		ActualRotationDirection:   string(v.ActualRotationDirection),
		AnsweredRotationDirection: string(v.AnsweredRotationDirection),
		Response:                  string(v.Response),
		MeanEstimation:            float64(v.MeanEstimation),
		SDEstimation:              float64(v.SDEstimation),
		LowerAsymptoteEstimation:  float64(v.LowerAsymptoteEstimation),
		LapseRateEstimation:       float64(v.LapseRateEstimation),
	}
}

type ResultMDDCWData struct {
	ExperimentMDDCW ExperimentMDDCWData          `json:"experiment_mddcw"`
	ResultDetail    []QuestPlusResultNormCDFData `json:"result_detail"`
	Subject         SubjectData                  `json:"subject"`
	Mean            float64                      `json:"mean"`
	Sd              float64                      `json:"sd"`
	LowerAsymptote  float64                      `json:"lower_asymptote"`
	LapseRate       float64                      `json:"lapse_rate"`
}

func NewResultMDDCWData(v *ResultMDDCW) *ResultMDDCWData {
	resultDetailData := make([]QuestPlusResultNormCDFData, len(v.ResultDetail))
	for _, x := range v.ResultDetail {
		q := NewQuestPlusResultNormCDFData(&x)
		resultDetailData = append(resultDetailData, *q)
	}
	return &ResultMDDCWData{
		ExperimentMDDCW: *NewExperimentMDDCWData(&v.ExperimentMDDCW),
		ResultDetail:    resultDetailData,
		Subject:         *NewSubjectData(&v.Subject),
		Mean:            float64(v.Mean),
		Sd:              float64(v.SD),
		LowerAsymptote:  float64(v.LowerAsymptote),
		LapseRate:       float64(v.LapseRate),
	}
}

type SubjectData struct {
	Sex                    string `json:"sex"`
	Age                    int    `json:"age"`
	DeafAndHearingImpaired bool   `json:"deaf_and_hearing_impaired"`
}

func NewSubjectData(v *Subject) *SubjectData {
	return &SubjectData{
		Sex:                    string(v.Sex),
		Age:                    int(v.Age),
		DeafAndHearingImpaired: bool(v.DeafAndHearingImpaired),
	}
}
