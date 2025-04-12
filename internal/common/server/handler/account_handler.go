package handler

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	//"github.com/neel4os/warg/internal/account-management/domain/account/app/commands_events"
	"github.com/neel4os/warg/internal/account-management/domain/account/app/commands"
	"github.com/neel4os/warg/internal/common/errors"
	"github.com/neel4os/warg/internal/eventstore/domain/app"
	"github.com/neel4os/warg/pkg"
)

func (h *Handler) OnboardAccount(c echo.Context) error {
	_account := &commands.OnBoardAccount{}
	_account.AccountId = uuid.New()
	if err := c.Bind(_account); err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewBindError(err.Error()))
	}
	if err := c.Validate(_account); err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewBadRequestError(err.Error()))
	}
	ep := app.GetEventPlatform()
	err := ep.CommandBus.Send(context.Background(), _account)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.NewInternalServerError(err.Error()))
	}
	return c.JSON(http.StatusAccepted, pkg.OnboardingResponse{
		OnboardingId: _account.AccountId.String(),
	})
}
