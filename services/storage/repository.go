package storage

import (
	"errors"

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
			log.Error().Err(result.Error).Caller().Msg("database operation failed because of duplicat eentry")
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
