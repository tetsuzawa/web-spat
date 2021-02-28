package handler

// IntegratedHandler implements ServerInterface generated by OpenAPI.
type IntegratedHandler struct {
	ExperimentHandler
	UtilHandler
}

func NewIntegratedHandler(e ExperimentHandler, u UtilHandler) *IntegratedHandler {
	return &IntegratedHandler{ExperimentHandler: e, UtilHandler: u}
}
