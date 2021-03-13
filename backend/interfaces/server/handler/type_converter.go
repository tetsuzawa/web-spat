package handler

import (
	"github.com/tetsuzawa/web-spat/domain"
	"github.com/tetsuzawa/web-spat/interfaces/server/openapi"
)

func NewExperimentIdData(v openapi.ExperimentId) domain.ExperimentIdData {
	return domain.ExperimentIdData(v)
}

func NewExperimentMDDData(v *openapi.ExperimentMDD) *domain.ExperimentMDDData {
	return &domain.ExperimentMDDData{
		Id:                        uint64(v.Id),
		QuestPlusParameterNormCDF: *NewQuestPlusParameterNormCDFData(&v.QuestPlusParameterNormCdf),
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

func NewExperimentMDDOAPI(v *domain.ExperimentMDDData) *openapi.ExperimentMDD {
	return &openapi.ExperimentMDD{
		Id:                        openapi.ExperimentId(v.Id),
		QuestPlusParameterNormCdf: *NewQuestPlusParameterNormCDFOAPI(&v.QuestPlusParameterNormCDF),
		Name:                      openapi.ExperimentName(v.Name),
		Description:               openapi.ExperimentDescription(v.Description),
		Azimuth:                   openapi.Azimuth(v.Azimuth),
		Altitude:                  openapi.Altitude(v.Altitude),
		CoordinateVariable:        openapi.CoordinateVariable(v.CoordinateVariable),
		MovingSoundConstant:       openapi.MovingSoundConstant(v.MovingSoundConstant),
		MovingSoundConstantValue:  int64(v.MovingSoundConstantValue),
		NumTrials:                 openapi.NumTrials(v.NumTrials),
	}
}

func NewQuestPlusParameterNormCDFData(v *openapi.QuestPlusParameterNormCDF) *domain.QuestPlusParameterNormCDFData {
	return &domain.QuestPlusParameterNormCDFData{
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

func NewQuestPlusParameterNormCDFOAPI(v *domain.QuestPlusParameterNormCDFData) *openapi.QuestPlusParameterNormCDF {
	return &openapi.QuestPlusParameterNormCDF{
		StimDomain:            *NewStimDomainNormCDFOAPI(&v.StimDomain),
		ParamDomain:           *NewParamDomainNormCDFOAPI(&v.ParamDomain),
		OutcomeDomain:         *NewOutcomeDomainOAPI(&v.OutcomeDomain),
		Prior:                 *NewPriorNormCDFOAPI(&v.Prior),
		Func:                  openapi.QuestPlusFunc(v.Func),
		StimScale:             openapi.QuestPlusStimScale(v.StimScale),
		StimSelectionMethod:   openapi.QuestPlusStimSelectionMethod(v.StimSelectionMethod),
		ParamEstimationMethod: openapi.QuestPlusParamEstimationMethod(v.ParamEstimationMethod),
	}
}

func NewStimDomainNormCDFData(v *openapi.StimDomainNormCDF) *domain.StimDomainNormCDFData {
	return &domain.StimDomainNormCDFData{
		Intensity: v.Intensity,
	}
}

func NewStimDomainNormCDFOAPI(v *domain.StimDomainNormCDFData) *openapi.StimDomainNormCDF {
	return &openapi.StimDomainNormCDF{
		Intensity: v.Intensity,
	}
}
func NewParamDomainNormCDFData(v *openapi.ParamDomainNormCDF) *domain.ParamDomainNormCDFData {
	mean := make([]float64, len(v.Mean))
	for _, x := range v.Mean {
		mean = append(mean, float64(x))
	}
	sd := make([]float64, len(v.Sd))
	for _, x := range v.Sd {
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

	return &domain.ParamDomainNormCDFData{
		Mean:           mean,
		SD:             sd,
		LowerAsymptote: lowerAsymptote,
		LapseRate:      lapseRate,
	}
}

func NewParamDomainNormCDFOAPI(v *domain.ParamDomainNormCDFData) *openapi.ParamDomainNormCDF {
	mean := make([]openapi.Mean, len(v.Mean))
	for _, x := range v.Mean {
		mean = append(mean, openapi.Mean(x))
	}
	sd := make([]openapi.SD, len(v.SD))
	for _, x := range v.SD {
		sd = append(sd, openapi.SD(x))
	}
	lowerAsymptote := make([]openapi.LowerAsymptote, len(v.LowerAsymptote))
	for _, x := range v.LowerAsymptote {
		lowerAsymptote = append(lowerAsymptote, openapi.LowerAsymptote(x))
	}
	lapseRate := make([]openapi.LapseRate, len(v.LapseRate))
	for _, x := range v.LapseRate {
		lapseRate = append(lapseRate, openapi.LapseRate(x))
	}

	return &openapi.ParamDomainNormCDF{
		Mean:           mean,
		Sd:             sd,
		LowerAsymptote: lowerAsymptote,
		LapseRate:      lapseRate,
	}
}

func NewOutcomeDomainData(v *openapi.OutcomeDomain) *domain.OutcomeDomainData {
	resp := make([]string, len(v.Response))
	for i, x := range v.Response {
		resp[i] = string(x)
	}
	return &domain.OutcomeDomainData{
		Response: resp,
	}
}

func NewOutcomeDomainOAPI(v *domain.OutcomeDomainData) *openapi.OutcomeDomain {
	resp := make([]openapi.Response, len(v.Response))
	for i, x := range v.Response {
		resp[i] = openapi.Response(x)
	}
	return &openapi.OutcomeDomain{
		Response: resp,
	}
}

func NewPriorNormCDFData(v *openapi.PriorNormCDF) *domain.PriorNormCDFData {
	mean := make([]float64, len(v.Mean))
	for _, x := range v.Mean {
		mean = append(mean, float64(x))
	}
	sd := make([]float64, len(v.Sd))
	for _, x := range v.Sd {
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

	return &domain.PriorNormCDFData{
		Mean:           mean,
		SD:             sd,
		LowerAsymptote: lowerAsymptote,
		LapseRate:      lapseRate,
	}
}

func NewPriorNormCDFOAPI(v *domain.PriorNormCDFData) *openapi.PriorNormCDF {
	mean := make([]openapi.Probability, len(v.Mean))
	for _, x := range v.Mean {
		mean = append(mean, openapi.Probability(x))
	}
	sd := make([]openapi.Probability, len(v.SD))
	for _, x := range v.SD {
		sd = append(sd, openapi.Probability(x))
	}
	lowerAsymptote := make([]openapi.Probability, len(v.LowerAsymptote))
	for _, x := range v.LowerAsymptote {
		lowerAsymptote = append(lowerAsymptote, openapi.Probability(x))
	}
	lapseRate := make([]openapi.Probability, len(v.LapseRate))
	for _, x := range v.LapseRate {
		lapseRate = append(lapseRate, openapi.Probability(x))
	}

	return &openapi.PriorNormCDF{
		Mean:           mean,
		Sd:             sd,
		LowerAsymptote: lowerAsymptote,
		LapseRate:      lapseRate,
	}
}

func NewQuestPlusResultNormCDFData(v *openapi.QuestPlusResultNormCDF) *domain.QuestPlusResultNormCDFData {
	return &domain.QuestPlusResultNormCDFData{
		NumTrials:                 uint64(v.NumTrials),
		Width:                     uint64(v.Width),
		Velocity:                  uint64(v.Velocity),
		Azimuth:                   uint64(v.Azimuth),
		Altitude:                  int64(v.Altitude),
		ActualRotationDirection:   string(v.ActualRotationDirection),
		AnsweredRotationDirection: string(v.AnsweredRotationDirection),
		Response:                  string(v.Response),
		MeanEstimation:            float64(v.MeanEstimation),
		SDEstimation:              float64(v.SdEstimation),
		LowerAsymptoteEstimation:  float64(v.LowerAsymptoteEstimation),
		LapseRateEstimation:       float64(v.LapseRateEstimation),
	}
}

func NewQuestPlusResultNormCDFOAPI(v *domain.QuestPlusResultNormCDFData) *openapi.QuestPlusResultNormCDF {
	return &openapi.QuestPlusResultNormCDF{
		NumTrials:                 openapi.NumTrials(v.NumTrials),
		Width:                     openapi.Width(v.Width),
		Velocity:                  openapi.Velocity(v.Velocity),
		Azimuth:                   openapi.Azimuth(v.Azimuth),
		Altitude:                  openapi.Altitude(v.Altitude),
		ActualRotationDirection:   openapi.RotationDirection(v.ActualRotationDirection),
		AnsweredRotationDirection: openapi.RotationDirection(v.AnsweredRotationDirection),
		Response:                  openapi.Response(v.Response),
		MeanEstimation:            openapi.Mean(v.MeanEstimation),
		SdEstimation:              openapi.SD(v.SDEstimation),
		LowerAsymptoteEstimation:  openapi.LowerAsymptote(v.LowerAsymptoteEstimation),
		LapseRateEstimation:       openapi.LapseRate(v.LapseRateEstimation),
	}
}

func NewResultMDDData(v *openapi.ResultMDD) *domain.ResultMDDData {
	resultDetail := make([]domain.QuestPlusResultNormCDFData, len(v.ResultDetail))
	for _, x := range v.ResultDetail {
		q := NewQuestPlusResultNormCDFData(&x)
		resultDetail = append(resultDetail, *q)
	}
	return &domain.ResultMDDData{
		ExperimentMDD:  *NewExperimentMDDData(&v.ExperimentMdd),
		ResultDetail:   resultDetail,
		Subject:        *NewSubjectData(&v.Subject),
		Mean:           float64(v.Mean),
		Sd:             float64(v.Sd),
		LowerAsymptote: float64(v.LowerAsymptote),
		LapseRate:      float64(v.LapseRate),
	}
}

func NewResultMDDOAPI(v *domain.ResultMDDData) *openapi.ResultMDD {
	resultDetail := make([]openapi.QuestPlusResultNormCDF, len(v.ResultDetail))
	for _, x := range v.ResultDetail {
		q := NewQuestPlusResultNormCDFOAPI(&x)
		resultDetail = append(resultDetail, *q)
	}
	return &openapi.ResultMDD{
		ExperimentMdd:  *NewExperimentMDDOAPI(&v.ExperimentMDD),
		ResultDetail:   resultDetail,
		Subject:        *NewSubjectOAPI(&v.Subject),
		Mean:           openapi.Mean(v.Mean),
		Sd:             openapi.SD(v.Sd),
		LowerAsymptote: openapi.LowerAsymptote(v.LowerAsymptote),
		LapseRate:      openapi.LapseRate(v.LapseRate),
	}
}

func NewSubjectData(v *openapi.Subject) *domain.SubjectData {
	return &domain.SubjectData{
		Sex:                    string(v.Sex),
		Age:                    uint64(v.Age),
		DeafAndHearingImpaired: bool(v.DeafAndHearingImpaired),
	}
}

func NewSubjectOAPI(v *domain.SubjectData) *openapi.Subject {
	return &openapi.Subject{
		Sex:                    openapi.Sex(v.Sex),
		Age:                    openapi.Age(v.Age),
		DeafAndHearingImpaired: openapi.DeafAndHearingImpaired(v.DeafAndHearingImpaired),
	}
}
