package storage

import "fmt"

type StorageError struct {
	ErrorCode       StorageErrorCode `json:"error_code"`
	Message         string           `json:"message"`
	DetailedMessage string           `json:"detailed_message"`
}

func (se StorageError) Error() string {
	return fmt.Sprintf("error_code = %s, message = %s detailed_message=%s", se.ErrorCode, se.Message, se.DetailedMessage)
}

type StorageErrorCode string

const (
	DuplicateKeyExists StorageErrorCode = "STO-1000"
	UnableToUploadFile StorageErrorCode = "STO-2000"
	NoRecordFound StorageErrorCode = "STO-2001"
	NotValidStorageIdInEventInternal StorageErrorCode = "STO-2002"
)
