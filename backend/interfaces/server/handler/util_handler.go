package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UtilHandler struct{}

func NewUtilHandler() *UtilHandler {
	return &UtilHandler{}
}

// Ping returns "OK"
// (GET /ping)
func (h *UtilHandler) Ping(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "OK")
}
