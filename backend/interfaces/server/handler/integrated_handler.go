package handler

// IntegratedHandler implements ServerInterface generated by OpenAPI.
type IntegratedHandler struct {
	ExperimentsHandler
	UtilHandler
}

func NewIntegratedHandler(e ExperimentsHandler, u UtilHandler) *IntegratedHandler {
	return &IntegratedHandler{ExperimentsHandler: e, UtilHandler: u}
}