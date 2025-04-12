package account

import (
	"github.com/neel4os/warg/internal/account-management/domain/account/aggregates/value"
)

type accountDatabaseRepository struct{}

func NewAccountDatabaseRepository() *accountDatabaseRepository {
	return &accountDatabaseRepository{}
}

func (r *accountDatabaseRepository) CreateAccount(req value.AccountCreationRequest) error {
	return nil
}
