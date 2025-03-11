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
	repo       *StorageRepo
	s3         *S3Client
	Filename   string
	Content    multipart.File
	TaskStatus TaskStatus
}

func NewUploadTask(repo *StorageRepo, s3 *S3Client, filename *multipart.FileHeader) *UploadTask {
	content, _ := filename.Open() //we can ignore this because its done earlier
	return &UploadTask{
		repo:       repo,
		s3:         s3,
		Filename:   filename.Filename,
		Content:    content,
		TaskStatus: TaskPending,
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
		msg := err.Error()
		log.Error().Caller().Msg("async job terminating...")
		u.repo.UpdateTaskStatus(storageId, TaskFailed, true, &InternalEvent{
			StorageId: storageId,
			Message:   &msg,
			State:     TaskPending,
		})
		return
	}
	// then we set the status of the task to initiated
	err = retry.Do(func() error {
		return u.uploadStatusInitiated(storageId)
	}, retry.Attempts(3), retry.OnRetry(func(n uint, err error) {
		log.Debug().Caller().Msg("Retrying to execute uploadStatusInitiated: " + strconv.Itoa(int(n)+1) + "/3")
	}))
	if err != nil {
		msg := err.Error()
		log.Error().Caller().Msg("async job terminating...")
		u.repo.UpdateTaskStatus(storageId, TaskFailed, true, &InternalEvent{
			StorageId: storageId,
			Message:   &msg,
			State:     TaskInitiated,
		})
		return
	}
	// we upload the file in bucket
	err = retry.Do(func() error {
		return u.putContentInBucket()
	}, retry.Attempts(3), retry.OnRetry(func(n uint, err error) {
		log.Debug().Caller().Msg("Retrying to execute putContentInBucket: " + strconv.Itoa(int(n)+1) + "/3")
	}))
	if err != nil {
		msg := err.Error()
		log.Error().Caller().Msg("async job terminating...")
		u.repo.UpdateTaskStatus(storageId, TaskFailed, true, &InternalEvent{
			StorageId: storageId,
			Message:   &msg,
			State:     TaskInitiated,
		})
		return
	}
	u.Content.Close()
	// if everything is successfull then we update the status of the task to success
	err = retry.Do(func() error {
		return u.repo.UpdateTaskStatus(storageId, TaskSucceded, true, &InternalEvent{
			StorageId: storageId,
			Message:   nil,
			State:     TaskSucceded,
		})
	}, retry.Attempts(3), retry.OnRetry(func(n uint, err error) {
		log.Debug().Caller().Msg("Retrying to execute UpdateTaskSucceded: " + strconv.Itoa(int(n)+1) + "/3")
	}))
	if err != nil {
		msg := err.Error()
		log.Error().Caller().Msg("async job terminating...")
		u.repo.UpdateTaskStatus(storageId, TaskFailed, true, &InternalEvent{
			StorageId: storageId,
			Message:   &msg,
			State:     TaskInitiated,
		})
		return
	}
	log.Info().Caller().Msg("async job completed")
}

func (u *UploadTask) IfStorageIDValid(storageId uuid.UUID) error {
	st, err := u.repo.GetStorageById(storageId)
	if err != nil {
		if errors.Is(err, StorageError{}) {
			// if we do not find the record there is no point in retry
			// but we failed as a task
			log.Error().Err(err).Caller().Msg("could not find storage id " + storageId.String())
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

func (u *UploadTask) putContentInBucket() error {
	log.Debug().Caller().Interface("content", u.Content).Msg("uploading file in bucket")
	err := u.s3.PutFileInBucket(u.Filename, u.Content)
	if err != nil {
		log.Error().Err(err).Caller().Msg("unable to put file in bucket")
		return err
	}
	log.Debug().Caller().Msg("file uploaded in bucket")
	return nil
}

func (u *UploadTask) uploadStatusInitiated(storageId uuid.UUID) error {
	isAlreadyInitiated := u.TaskStatus == TaskInitiated
	if !isAlreadyInitiated {
		// this job have never been success still
		err := u.repo.UpdateTaskStatus(storageId, TaskInitiated, false, nil)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}
