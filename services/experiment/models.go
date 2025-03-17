package experiment

import (
	"github.com/google/uuid"
	"github.com/neel4os/warg/services/storage"
	"gorm.io/datatypes"
)

type ExperimentCreationResponse struct {
	storage.Base
	Metadata         datatypes.JSON   `json:"metadata" gorm:"not null"`
	ValidationStatus ValidationStatus `json:"validation_status" gorm:"not null"`
}

type Experiment struct {
	ExperimentCreationResponse
	StorageId uuid.UUID `json:"storage_id" gorm:"not null"`
}
