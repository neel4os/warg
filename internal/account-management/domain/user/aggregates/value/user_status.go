package value

type UserStatus string

const (
	UserStatusActive   UserStatus = "active"
	UserStatusInactive UserStatus = "inactive"
	UserStatusPending  UserStatus = "pending"
	UserStatusDeleted  UserStatus = "deleted"
)

func (s UserStatus) String() string {
	return string(s)
}
