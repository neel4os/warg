package aggregates

import (
	"time"

	"github.com/google/uuid"
	"github.com/neel4os/warg/internal/account-management/domain/account/aggregates/value"
)


type AccountDAO struct {
	ID        uuid.UUID           `json:"id"`
	UpdatedAt time.Time           `json:"updated_at"`
	Name      string              `json:"name"`
	Status    value.AccountStatus `json:"status"`
}

func NewAccountDAO(name, firstName, lastName, email string) *AccountDAO {
	return &AccountDAO{
		Name:   name,
		Status: value.AccountStatusPending,
	}
}
