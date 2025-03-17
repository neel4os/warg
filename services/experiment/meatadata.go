package experiment

import (
	"encoding/json"
)

type Metadata struct {
	RollbackStrategy    RollbackStrategy   `json:"rollback-strategy" validate:"required,oneof=default always deviated never"`
	HypothesisStrategy  HypothesisStrategy `json:"hypothesis-strategy" validate:"required,oneof=default before-method-only after-method-only during-method-only continuously"`
	HypothesisFrequency float64            `json:"hypothesis-frequency" `
	FailFast            bool               `json:"fail-fast"`
}

func NewMetadata(rollback, hypothesis string, frequency float64, failFast bool) *Metadata {
	return &Metadata{
		RollbackStrategy:    GetRollbackStrategy(rollback),
		HypothesisStrategy:  GetHypothesisStrategy(hypothesis),
		HypothesisFrequency: frequency,
		FailFast:            failFast,
	}
}

func (m *Metadata) ConvertToBytes() []byte {
	byteFormat, err := json.Marshal(m)
	if err != nil {
		return nil
	}
	return byteFormat
}

// func (m *Metadata) Validate() error {
// 	validate := validator.New(validator.WithPrivateFieldValidation())
// 	err := validate.Struct(m)
// 	if err != nil {
// 		return helpers.CustomWargError{
// 			ErrorCode:       InvalidRequest,
// 			Message:         "Invalid request",
// 			DetailedMessage: err.Error(),
// 		}
// 	}
// 	return nil
// }

type RollbackStrategy string

const (
	// RollbackStrategyNone is a constant for none rollback strategy
	RollbackStrategyDefault    RollbackStrategy = "default"
	RollbackStrategyAlways     RollbackStrategy = "always"
	RollbackStrategyOnDeviated RollbackStrategy = "deviated"
	RollbackStrategyNever      RollbackStrategy = "never"
)

func GetRollbackStrategy(rollback string) RollbackStrategy {
	switch rollback {
	case "always":
		return RollbackStrategyAlways
	case "deviated":
		return RollbackStrategyOnDeviated
	case "never":
		return RollbackStrategyNever
	case "default":
		return RollbackStrategyDefault
	case "":
		return RollbackStrategyAlways
	default:
		return ""
	}
}

type HypothesisStrategy string

const (
	HypothesisStrategyDefault HypothesisStrategy = "default"
	HypothesisStrategyBMO     HypothesisStrategy = "before-method-only"
	HypothesisStrategyAMO     HypothesisStrategy = "after-method-only"
	HypothesisStrategyDM      HypothesisStrategy = "during-method-only"
	HypothesisStrategyCon     HypothesisStrategy = "continuously"
)

func GetHypothesisStrategy(hypothesis string) HypothesisStrategy {
	switch hypothesis {
	case "before-method-only":
		return HypothesisStrategyBMO
	case "after-method-only":
		return HypothesisStrategyAMO
	case "during-method-only":
		return HypothesisStrategyDM
	case "continuously":
		return HypothesisStrategyCon
	case "default":
		return HypothesisStrategyDefault
	case "":
		return HypothesisStrategyDefault
	default:
		return ""
	}
}

type ValidationStatus string

const (
	ValidationStatusSuccess ValidationStatus = "success"
	ValidationStatusFailed  ValidationStatus = "failed"
	ValidationStatusSkipped ValidationStatus = "skipped"
	ValidationStatusPending ValidationStatus = "pending"
	ValidationStatusUnknown ValidationStatus = "unknown"
	ValidationStatusError   ValidationStatus = "error"
	ValidationStatusRunning ValidationStatus = "running"
)

func GetValidationStatus(status string) ValidationStatus {
	switch status {
	case "success":
		return ValidationStatusSuccess
	case "failed":
		return ValidationStatusFailed
	case "skipped":
		return ValidationStatusSkipped
	case "pending":
		return ValidationStatusPending
	case "unknown":
		return ValidationStatusUnknown
	case "error":
		return ValidationStatusError
	case "running":
		return ValidationStatusRunning
	default:
		return ""
	}
}
