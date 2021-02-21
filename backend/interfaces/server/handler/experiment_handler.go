package handler

import (
	"github.com/labstack/echo/v4"
)

type ExperimentHandler struct{}

func NewExperimentsHandler() *ExperimentHandler {
	return &ExperimentHandler{}
}

// ListExperiments returns a list of experiments.
// (GET /experiments)
func (h *ExperimentHandler) ListExperiments(ctx echo.Context) error {


	return nil
}

// GetExperimentById represents a experiment by specified ID
// (GET /experiments/{id})
func (h *ExperimentHandler) GetExperimentById(ctx echo.Context, id int) error {
	// TODO
	return nil
}

// RegisterResultOfExperimentId registers the result of the specified experiment ID
// (POST /experiments/{id}/results)
func (h *ExperimentHandler) RegisterResultOfExperimentId(ctx echo.Context, id int) error {
	// TODO
	return nil
}
