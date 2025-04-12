package value

type (
	OrganizationStatus string
)

const (
	OrganizationStatusActive   OrganizationStatus = "active"
	OrganizationStatusPending  OrganizationStatus = "pending"
	OrganizationStatusDeleted  OrganizationStatus = "deleted"
)

func (s OrganizationStatus) String() string {
	return string(s)
}
