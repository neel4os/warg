package keycloak

type DomainRepresensation struct {
	Name     string `json:"name"`
	Verified bool   `json:"verified"`
}

type OrganizationRepresentation struct {
	Name    string                 `json:"name"`
	Enabled bool                   `json:"enabled"`
	Domains []DomainRepresensation `json:"domains"`
}

type UserRepresentation struct {
	Email         string `json:"email"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Enabled       bool   `json:"enabled"`
	EmailVerified bool   `json:"emailVerified"`
}

func NewUserRepresentation(email string, firstName string, lastName string) *UserRepresentation {
	return &UserRepresentation{
		Email:         email,
		FirstName:     firstName,
		LastName:      lastName,
		Enabled:       true,
		EmailVerified: false,
	}
}

func NewOrganizationRepresentation(name string) *OrganizationRepresentation {
	return &OrganizationRepresentation{
		Name:    name,
		Enabled: true,
		Domains: []DomainRepresensation{
			{
				Name:     name,
				Verified: false,
			},
		},
	}
}
