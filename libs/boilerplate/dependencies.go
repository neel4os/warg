package boilerplate

type Dependent interface {
	Ping() error
	Close() error
}
