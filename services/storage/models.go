package storage

import (
	// "fmt"
	"time"

	"github.com/google/uuid"
)

type Health struct {
	Status string `json:"status"`
}

type TaskStatus string

const (
	TaskPending   TaskStatus = "pending"
	TaskInitiated TaskStatus = "initiated"
	TaskSucceded  TaskStatus = "succeded"
	TaskFailed    TaskStatus = "failed"
)

type Base struct {
	Id        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time `json:"created_at" gorm:"not null,autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null, autoCreateTime"`
}

type Storage struct {
	Base
	StorageCreationRequest
	FileHash   string     `json:"file_hash" gorm:"unique;not null"`
	FileSize   int64      `json:"file_size" gorm:"not null"`
	TaskStatus TaskStatus `json:"task_status" gorm:"not null;default:0"`
}

type StorageCreationRequest struct {
	Filename string `json:"filename" gorm:"not null"`
}

type InternalEvent struct {
	Id        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time `gorm:"not null,autoCreateTime"`
	StorageId uuid.UUID
	Storage   Storage `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Message   *string
	State     TaskStatus `gorm:"not null"`
}
