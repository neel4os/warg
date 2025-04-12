package organization

import (
	"github.com/google/uuid"
	"github.com/neel4os/warg/internal/account-management/domain/organization/aggregates/value"
)

type Organization struct {
	AccountId  uuid.UUID                `json:"account_id"`
	DomainName string                   `json:"domain_name"`
	ID         uuid.UUID                `json:"id"`
	Name       string                   `json:"name"`
	Status     value.OrganizationStatus `json:"status"`
}

type OrganizationCreated struct {
	AccountId      uuid.UUID
	OrganizationId uuid.UUID
	OwnerEmail     string
	OwnerFirstName string
	OwnerLastName  string
}
