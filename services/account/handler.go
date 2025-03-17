package account

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/neel4os/warg/libs/boilerplate"
	"github.com/neel4os/warg/services/storage"
)

type AccountHandler struct {
	deps map[string]boilerplate.Dependent
	cfg  boilerplate.ConfigIface
}

func NewAccountHandler(deps map[string]boilerplate.Dependent, cfg boilerplate.ConfigIface) *AccountHandler {
	return &AccountHandler{
		deps: deps,
		cfg:  cfg,
	}
}

func (a *AccountHandler) GetHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, storage.Health{Status: "OK"})
}

func (a *AccountHandler) CreateUser(c echo.Context) error {
	var user UserCreationRequest
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "could not bind the request")
	}
	return c.JSON(http.StatusAccepted, user)
}
