package storage

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/neel4os/warg/libs/boilerplate"
)

type StorageHandler struct {
	deps map[string]boilerplate.Dependent
	cfg  boilerplate.ConfigIface
}

func NewStorageHandler(deps map[string]boilerplate.Dependent, cfg boilerplate.ConfigIface) *StorageHandler {
	return &StorageHandler{
		deps: deps,
		cfg:  cfg,
	}
}

func (h *StorageHandler) GetHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, Health{Status: "OK"})
}


