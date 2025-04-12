package handler

import (
	"github.com/neel4os/warg/internal/common/config"
	"github.com/neel4os/warg/internal/common/database"
)


type Handler struct {
	cfg *config.Config
	dbcon *database.DataConn
	
}

func NewHandler(cfg *config.Config, dbcon *database.DataConn) *Handler {
	return &Handler{
		cfg: cfg,
		dbcon: dbcon,
	}
}
