package storage

import (
	"errors"

	"gorm.io/gorm"
)


func (sr *StorageRepo) CreateInternalEvent(event *InternalEvent) error {
	result := sr.dbcon.Create(event)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrForeignKeyViolated) {
			return StorageError{
				ErrorCode: NotValidStorageIdInEventInternal,
				Message: event.StorageId.String() +  " is not a valid storage id",
			}
		}
		return result.Error
	}
	return nil
}