package events

import (
	"github.com/google/uuid"
	"github.com/neel4os/warg/internal/account-management/domain/account/aggregates/value"
)

type AccountOnboarded struct {
	AccountId   uuid.UUID
	AccountName string
	FirstName   string
	LastName    string
	Email       string
	Status      value.AccountStatus
}


