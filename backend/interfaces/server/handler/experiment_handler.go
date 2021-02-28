package handler

import (
	"net/http"

	"github.com/tetsuzawa/web-spat/interfaces/server/openapi"

	"github.com/labstack/echo/v4"
	"github.com/tetsuzawa/web-spat/usecase"
)

type ExperimentHandler struct {
	u usecase.IExperimentUseCase
}

func NewExperimentsHandler(u usecase.IExperimentUseCase) *ExperimentHandler {
	return &ExperimentHandler{u: u}
}

// Returns a list of active experiments.
// (GET /experiment/mdd/active)
func (h *ExperimentHandler) ListExperimentsMDDActive(ctx echo.Context) error {
	experimentMDDs, err := h.u.ListMDDActive(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, experimentMDDs)
}

// Returns a list of inactive experiments.
// (GET /experiment/mdd/inactive)
func (h *ExperimentHandler) ListExperimentsMDDInactive(ctx echo.Context) error {
	experimentMDDs, err := h.u.ListMDDInactive(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, experimentMDDs)
}

// Represents a experiment by specified ID.
// (GET /experiment/mdd/{id})
func (h *ExperimentHandler) GetExperimentMDDById(ctx echo.Context, id openapi.ExperimentId) error {
	// TODO
	return nil
}

// Register the result of the specified experiment ID.
// (POST /experiment/mdd/{id}/results)
func (h *ExperimentHandler) RegisterResultOfExperimentMDDById(ctx echo.Context, id openapi.ExperimentId) error {
	// TODO
	return nil
}
