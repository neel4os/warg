package experiment

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/neel4os/warg/libs/boilerplate"
	"github.com/neel4os/warg/services/storage"
)


type ExperimentHandler struct {
	deps map[string]boilerplate.Dependent
	cfg boilerplate.ConfigIface
}

func NewExperimentHandler(deps map[string]boilerplate.Dependent, cfg boilerplate.ConfigIface) *ExperimentHandler {
	return &ExperimentHandler{
		deps: deps,
		cfg: cfg,
	}
}

func (e *ExperimentHandler) GetHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, storage.Health{Status: "OK"})
}