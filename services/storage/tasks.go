package storage

import (
	"errors"
	"mime/multipart"
	"strconv"

	"github.com/avast/retry-go"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type UploadTask struct {
	repo             *StorageRepo
	s3               *S3Client
	Filename         string
	Content          multipart.File
	continueNextStep bool
	TaskStatus       TaskStatus
}

func NewUploadTask(repo *StorageRepo, s3 *S3Client, filename string, content multipart.File) *UploadTask {
	return &UploadTask{
		repo:             repo,
		s3:               s3,
		Filename:         filename,
		Content:          content,
		continueNextStep: true,
		TaskStatus:       TaskPending,
	}
}

func (u *UploadTask) Execute(storageId uuid.UUID) {
	log.Debug().Str("tasks_id", storageId.String()).Msg("initiating async job")
	// first check that if storageID is valid or not
	err := retry.Do(func() error {
		return u.IfStorageIDValid(storageId)
	}, retry.Attempts(3), retry.OnRetry(func(n uint, err error) {
		log.Debug().Caller().Msg("Retrying to execute IfStorageIDValid: " + strconv.Itoa(int(n)+1) + "/3")
	}))
	if err != nil {
		log.Error().Caller().Msg("async job terminating...")
		u.continueNextStep = false
	}
	if u.continueNextStep {
		err = retry.Do(func() error {
			return u.uploadStatusInitiated(storageId)
		}, retry.Attempts(3), retry.OnRetry(func(n uint, err error) {
			log.Debug().Caller().Msg("Retrying to execute uploadStatusInitiated: " + strconv.Itoa(int(n)+1) + "/3")
		}))
		if err != nil {
			log.Error().Caller().Msg("async job terminating...")
			u.continueNextStep = false
		}
	}
}

func (u *UploadTask) IfStorageIDValid(storageId uuid.UUID) error {
	st, err := u.repo.GetStorageById(storageId)
	if err != nil {
		if errors.Is(err, StorageError{}) {
			// if we do not find the record there is no point in retry
			// but we failed as a task
			log.Error().Err(err).Caller().Msg("could not find storage id " + storageId.String())
			u.continueNextStep = false
			return nil
		} else {
			// we need to retry
			log.Error().Err(err).Caller().Msg("unable to fetch details of storage")
			return err
		}
	}
	u.TaskStatus = st.TaskStatus
	return nil
}

// func (u *UploadTask) UploadfileInBucket(storageId uuid.UUID) (bool, error) {
// 	if u.Status == "failed" {
// 		// there is no point in retry
// 		return false, nil
// 	}
// 	err := u.uploadStatusInitiated(storageId)
// 	if err != nil {
// 		// status change failed but need to retyr
// 		return false, err
// 	}
// 	//now try to upload the file in bucket
// 	err = u.putContentInBucket()
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil

// }

// func (u *UploadTask) putContentInBucket() error {
// 	isFileExist, err := u.s3.IsFileExistOnBucket(u.Filename)
// 	if err != nil {
// 		// this need retries, may be not but we will do it anyway
// 		return err
// 	}
// 	if !isFileExist {
// 		err := u.s3.PutFileInBucket(u.Filename, u.Content)
// 		if err != nil {
// 			// retry
// 			return err
// 		}
// 		return nil

// 	}
// 	// file already exist so dont need to anything
// 	return nil

// }

func (u *UploadTask) uploadStatusInitiated(storageId uuid.UUID) error {
	isAlreadyInitiated := u.TaskStatus == TaskInitiated
	if !isAlreadyInitiated {
		// this job have never been success still
		err := u.repo.UpdateTaskStatus(storageId, TaskInitiated)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

// func InitiateTask() {

// }
