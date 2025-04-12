package aggregates

import (
	"testing"

	"github.com/google/uuid"
	"github.com/neel4os/warg/internal/account-management/domain/account/aggregates/value"
)

func TestNewAccountDAO(t *testing.T) {
	name := "Test Account"
	firstName := "John"
	lastName := "Doe"
	email := "john.doe@example.com"

	account := NewAccountDAO(name, firstName, lastName, email)

	if account.Name != name {
		t.Errorf("expected Name to be '%s', got '%s'", name, account.Name)
	}

	if account.Status != value.AccountStatusPending {
		t.Errorf("expected Status to be '%s', got '%s'", value.AccountStatusPending, account.Status)
	}

	if account.ID == uuid.Nil {
		t.Error("expected ID to be a valid UUID, got nil UUID")
	}

	if account.UpdatedAt.IsZero() {
		t.Error("expected UpdatedAt to be a valid timestamp, got zero value")
	}
}