package repositories

import (
	//"github.com/neel4os/warg/internal/account-management/domain/account/aggregates"
	"github.com/neel4os/warg/internal/account-management/domain/account/aggregates/value"
)

type AccountRepositoryInterface interface {
	CreateAccount(value.AccountCreationRequest) error
	// GetAccountByID(string) (aggregates.AccountDAO, error)
	// GetAccountByEmail(string) (aggregates.AccountDAO, error)
	// GetAccounts() ([]aggregates.AccountDAO, error)
	// UpdateAccount(aggregates.AccountDAO) error
	// DeleteAccount(string) error
}
