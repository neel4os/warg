package storage

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *StorageHandler) CreateStorage(c echo.Context) error {
	filename, err := c.FormFile("filename")
	if err != nil {
		return c.JSON(http.StatusBadRequest, "could not able to read the form filename")
	}
	_storage := Storage{}
	_storage.Filename = filename.Filename
	_storage.FileSize = filename.Size
	_hash := sha256.New()
	src, err := filename.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, "could not open file "+filename.Filename)
	}
	defer src.Close()
	_, err = io.Copy(_hash, src)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "could not compute hash of the file "+_storage.Filename)
	}
	_storage.FileHash = hex.EncodeToString(_hash.Sum(nil))
	_storage.TaskStatus = TaskPending
	domainObj := NewStorageDomain(h.deps)
	asyncTask := NewUploadTask(domainObj.repo,
		h.deps["objectstorage"].(*S3Client),
		filename)
	sin, err := domainObj.Create(&_storage)
	if err != nil {
		var customErr StorageError
		isCustomErr := errors.As(err, &customErr)
		if isCustomErr {
			switch customErr.ErrorCode {
			case DuplicateKeyExists:
				return c.JSON(http.StatusConflict, customErr)
			}
		} else {
			return c.JSON(http.StatusInternalServerError, err)
		}
	}

	go asyncTask.Execute(sin.Id)
	return c.JSON(http.StatusAccepted, sin)
}
