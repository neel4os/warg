package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/neel4os/warg/pkg"
)

func (h *Handler) CheckHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, pkg.Health{
		Status: pkg.Up,
	})
}
