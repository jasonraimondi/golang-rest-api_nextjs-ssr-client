package lib

type CommandBus interface {
	Handle(c Command) error
}
