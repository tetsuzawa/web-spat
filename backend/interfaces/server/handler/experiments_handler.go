package handler

import "github.com/labstack/echo/v4"

type ExperimentsHandler struct{}

func NewExperimentsHandler() *ExperimentsHandler {
	return &ExperimentsHandler{}
}

// ListExperiments returns a list of experiments.
// (GET /experiments)
func (h *ExperimentsHandler) ListExperiments(ctx echo.Context) error {
	// TODO
	return nil
}

// GetExperimentById represents a experiment by specified ID
// (GET /experiments/{id})
func (h *ExperimentsHandler) GetExperimentById(ctx echo.Context, id int) error {
	// TODO
	return nil
}

// RegisterResultOfExperimentId registers the result of the specified experiment ID
// (POST /experiments/{id}/results)
func (h *ExperimentsHandler) RegisterResultOfExperimentId(ctx echo.Context, id int) error {
	// TODO
	return nil
}
