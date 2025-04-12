package value

type AccountStatus string

const (
	AccountStatusActive   AccountStatus = "active"
	AccountStatusInactive AccountStatus = "inactive"
	AccountStatusPending  AccountStatus = "created"
	AccountStatusDeleted  AccountStatus = "deleted"
)

func (s AccountStatus) IsValid() bool {
	switch s {
	case AccountStatusActive, AccountStatusInactive, AccountStatusPending, AccountStatusDeleted:
		return true
	}
	return false
}

func (s AccountStatus) String() string {
	switch s {
	case AccountStatusActive:
		return "active"
	case AccountStatusInactive:
		return "inactive"
	case AccountStatusPending:
		return "pending"
	case AccountStatusDeleted:
		return "deleted"
	default:
		return "unknown"
	}
}