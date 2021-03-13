package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tetsuzawa/web-spat/interfaces/server/openapi"
	"github.com/tetsuzawa/web-spat/usecase"
)

type ExperimentHandler struct {
	u usecase.IExperimentUseCase
}

func NewExperimentsHandler(u usecase.IExperimentUseCase) *ExperimentHandler {
	return &ExperimentHandler{u: u}
}

// Create a experiment by specified conditions.
// (POST /experiment/mdd)
func (h *ExperimentHandler) CreateExperimentMDD(ctx echo.Context) error {
	reqBody := &openapi.ExperimentMDD{}
	if err := ctx.Bind(reqBody); err != nil {
		return err
	}
	e := NewExperimentMDDData(reqBody)
	e, err := h.u.CreateMDD(ctx.Request().Context(), e)
	if err != nil {
		return err
	}
	respBody := NewExperimentMDDOAPI(e)
	return ctx.JSON(http.StatusOK, respBody)

}

// Returns a list of active experiments.
// (GET /experiment/mdd/active)
func (h *ExperimentHandler) ListExperimentsMDDActive(ctx echo.Context) error {
	experimentMDDs, err := h.u.ListMDDActive(ctx.Request().Context())
	if err != nil {
		return err
	}
	experimentMDDOAPIs := make(openapi.ExperimentsMDD, len(experimentMDDs))
	for i, v := range experimentMDDs {
		experimentMDDOAPIs[i] = *NewExperimentMDDOAPI(v)
	}
	return ctx.JSON(http.StatusOK, experimentMDDOAPIs)
}

// Returns a list of inactive experiments.
// (GET /experiment/mdd/inactive)
func (h *ExperimentHandler) ListExperimentsMDDInactive(ctx echo.Context) error {
	experimentMDDs, err := h.u.ListMDDInactive(ctx.Request().Context())
	if err != nil {
		return err
	}
	experimentMDDOAPIs := make(openapi.ExperimentsMDD, len(experimentMDDs))
	for i, v := range experimentMDDs {
		experimentMDDOAPIs[i] = *NewExperimentMDDOAPI(v)
	}
	return ctx.JSON(http.StatusOK, experimentMDDOAPIs)
}

// Represents a experiment by specified ID.
// (GET /experiment/mdd/{id})
func (h *ExperimentHandler) GetExperimentMDDById(ctx echo.Context, id openapi.ExperimentId) error {
	idData := NewExperimentIdData(id)
	experimentMDD, err := h.u.FindMDDById(ctx.Request().Context(), idData)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, NewExperimentMDDOAPI(experimentMDD))
}

// Register the result of the specified experiment ID.
// (POST /experiment/mdd/{id}/results)
func (h *ExperimentHandler) RegisterExperimentMDDResultById(ctx echo.Context, id openapi.ExperimentId) error {
	reqBody := &openapi.ResultMDD{}
	if err := ctx.Bind(reqBody); err != nil {
		return err
	}
	r := NewResultMDDData(reqBody)
	r, err := h.u.CreateMDDResult(ctx.Request().Context(), r)
	if err != nil {
		return err
	}
	respBody := NewResultMDDOAPI(r)
	return ctx.JSON(http.StatusOK, respBody)
}
