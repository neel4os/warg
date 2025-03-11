package boilerplate

type Dependent interface {
	Ping()
	Close() error
}
