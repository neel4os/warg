package repositories

type UserRepositoryInterface interface {
	CreateUser(firstname string, lastname string, email string) (string, error)
}


