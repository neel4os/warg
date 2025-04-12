package decorators

type CommandHandler interface {
	Handle() error
}
