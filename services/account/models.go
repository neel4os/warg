package account

import "github.com/neel4os/warg/services/storage"

type UsersTypes string

const (
	UsersTypesStandalone UsersTypes = "standalone"
	UsersTypesFederated  UsersTypes = "federated"
)

type Users struct {
	storage.Base
	UserCreationRequest
	UsersType UsersTypes `json:"users_type"`
	GroupId   *string    `json:"group_id"`
	OrgId     *string    `json:"org_id"`
}

type UserCreationRequest struct {
	Email     string `json:"email" gorm:"unique,not null"`
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`
}
