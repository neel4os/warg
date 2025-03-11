package storage

import (
	"errors"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type StorageRepo struct {
	dbcon *gorm.DB
}

func NewStorageRepo(dbcon *gorm.DB) *StorageRepo {
	return &StorageRepo{
		dbcon: dbcon,
	}
}

func (sr *StorageRepo) Create(sin *Storage) (*Storage, error) {
	result := sr.dbcon.Create(sin)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			log.Error().Err(result.Error).Caller().Msg("database operation failed because of duplicate eentry")
			return nil, StorageError{
				ErrorCode: DuplicateKeyExists,
				Message:   "file with same hash key exist",
			}
		} else {
			return nil, result.Error
		}
	}
	return sin, nil
}

func (sr *StorageRepo) UpdateTaskStatus(storageID uuid.UUID, status TaskStatus, publish_event bool, event *InternalEvent) error {
	sr.dbcon.Transaction(func(tx *gorm.DB) error {
		result := tx.Model(Storage{}).Where("id = ?", storageID).Update("task_status", status)
		if result.Error != nil {
			return result.Error
		}
		if publish_event {
			event := tx.Create(event)
			if event.Error != nil {
				return event.Error
			}
			return nil
		}
		return nil
	})
	return nil
}

func (sr *StorageRepo) GetStorageById(storageID uuid.UUID) (*Storage, error) {
	var _st Storage
	result := sr.dbcon.First(&_st, "id = ?", storageID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, StorageError{
				ErrorCode: NoRecordFound,
				Message:   "no storage exists with id " + storageID.String(),
			}
		} else {
			return nil, result.Error
		}
	}
	return &_st, nil
}
