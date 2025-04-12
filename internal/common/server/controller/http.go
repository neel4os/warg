package controller

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/neel4os/warg/internal/common/config"
	"github.com/neel4os/warg/internal/common/database"
	"github.com/neel4os/warg/internal/common/server/handler"
	"github.com/neel4os/warg/internal/common/util"
	"github.com/neel4os/warg/pkg"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type httpComponent struct {
	e   *echo.Echo
	cfg *config.Config
	dbcon *database.DataConn
}

func NewHTTPComponent(cfg *config.Config, dbcon *database.DataConn) *httpComponent {
	return &httpComponent{cfg: cfg, dbcon: dbcon}
}

func (h *httpComponent) Name() string {
	return "http"
}

func (h *httpComponent) Init() {
	h.e = echo.New()
	st := util.NewStaticFileLocation(nil)
	h.e.StaticFS("/", echo.MustSubFS(st.GetStaticFiles(), "console/.output/public"))
	h.customize()
	handler := handler.NewHandler(h.cfg, h.dbcon)
	pkg.RegisterHandlers(h.e, handler)
}

func (h *httpComponent) Run() {
	s := http.Server{
		Addr:         ":" + h.cfg.ServerConfig.Port,
		Handler:      h.e,
		ReadTimeout:  time.Duration(h.cfg.ServerConfig.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(h.cfg.ServerConfig.WriteTimeout) * time.Second,
	}
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal().Err(err).Caller().Msg("failed to start server")
	}
}

func (h *httpComponent) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(h.cfg.ServerConfig.GraceFullShutdownTime)*time.Second)
	defer cancel()
	if err := h.e.Shutdown(ctx); err != nil {
		h.e.Logger.Fatal(err)
	}
}

func (h *httpComponent) customize() {
	h.e.HideBanner = false
	h.e.HidePort = h.cfg.ServerConfig.HidePortInStdOut
	h.e.Validator = &CustomValidator{}
	// logger middleware
	logger := zerolog.New(os.Stdout)
	h.e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogMethod: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Str("Method", v.Method).
				Msg("request")

			return nil
		},
	}))
	// recover middleware
	h.e.Use(middleware.Recover())
	// request ID
	h.e.Use(middleware.RequestID())
}
