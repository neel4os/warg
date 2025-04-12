package repositories

import (
	"github.com/neel4os/warg/internal/eventstore/domain/aggregates"
	"gorm.io/gorm"
)

type EventRepositories interface {
	CreateEvent(*aggregates.Event) (*gorm.DB, error)
}
