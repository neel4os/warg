package storage

import (
	// "fmt"
	"time"

	"github.com/google/uuid"
)

type Health struct {
	Status string `json:"status"`
}

type UploadStatus string

const (
	UploadInitiated UploadStatus = "initiated"
	UploadPending   UploadStatus = "pending"
	UploadSucceded  UploadStatus = "succeded"
	UploadFailed    UploadStatus = "failed"
)

// func (u UploadStatus) String() string {
// 	switch u {
// 	case UploadInitiated:
// 		return "Initiated"
// 	case UploadPending:
// 		return "Pending"
// 	case UploadSucceded:
// 		return "Succeded"
// 	case UploadFailed:
// 		return "Failed"
// 	}
// 	return "Unknown"
// }

// func GetUploadStatus(u UploadStatus) string {
// 	return fmt.Sprintf("%d", u)
// }

type Base struct {
	Id        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time `json:"created_at" gorm:"not null,autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null, autoCreateTime"`
}

type Storage struct {
	Base
	StorageCreationRequest
	FileHash     string       `json:"file_hash" gorm:"unique;not null"`
	FileSize     int64        `json:"file_size" gorm:"not null"`
	UploadStatus UploadStatus `json:"upload_status" gorm:"not null;default:0"`
}

type StorageCreationRequest struct {
	Filename string `form:"filename" gorm:"not null"`
}
