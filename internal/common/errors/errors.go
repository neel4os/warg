package errors

import "encoding/json"

type errorCode string

const (
	// ErrorCodeBindError is the error code for binding errors.
	ErrorCodeBindError     errorCode = "WARG-400-001"
	ErrorJSONMarshalError  errorCode = "WARG-400-002"
	ErrorCodeBadRequest    errorCode = "WARG-400-003"
	ErrorDatabaseOperation errorCode = "WARG-500-001"
)

type wargError struct {
	ErrorCode       errorCode `json:"error_code"`
	Message         string    `json:"message"`
	DetailedMessage string    `json:"detailed_message"`
}

func (e *wargError) Error() string {
	res, err := json.MarshalIndent(e, "", "4")
	if err != nil {
		return NewJSONMarhsalError(err.Error()).Error()
	}
	return string(res)
}

func NewJSONMarhsalError(message string) *wargError {
	return &wargError{
		ErrorCode:       ErrorJSONMarshalError,
		Message:         "unable to marhsall output",
		DetailedMessage: message,
	}
}

func NewBindError(message string) *wargError {
	return &wargError{
		ErrorCode:       ErrorCodeBindError,
		Message:         "failed to bind request body",
		DetailedMessage: message,
	}
}

func NewBadRequestError(message string) *wargError {
	return &wargError{
		ErrorCode:       ErrorCodeBadRequest,
		Message:         "bad request",
		DetailedMessage: message,
	}
}

func NewInternalServerError(message string) *wargError {
	return &wargError{
		ErrorCode:       "WARG-500-001",
		Message:         "internal server error",
		DetailedMessage: message,
	}
}

func NewDatabaseOperationError(message string) *wargError {
	return &wargError{
		ErrorCode:       ErrorDatabaseOperation,
		Message:         "database operation error",
		DetailedMessage: message,
	}
}
