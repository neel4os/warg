package persistence

import (
	"github.com/neel4os/warg/internal/common/database"
	"github.com/neel4os/warg/internal/eventstore/domain/aggregates"
	"gorm.io/gorm"
)

type eventDatabaseRepository struct{
	dbcon *database.DataConn
}

func NewEventDatabaseRepository(dbcon *database.DataConn) *eventDatabaseRepository {
	return &eventDatabaseRepository{
		dbcon: dbcon,
	}
}
// we have decided to use transactional repository pattern
func (r *eventDatabaseRepository) CreateEvent(event *aggregates.Event) (*gorm.DB, error) {
	// Implement the logic to create an event in the database
	tx := r.dbcon.GetDb().Begin()
	if err := tx.Create(event).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return tx, nil
}
