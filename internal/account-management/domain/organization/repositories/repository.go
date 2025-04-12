package repositories

type OrganizationRepositoryInterface interface {
	CreateOrganization(name string) (string, error)
}
