package lib

import (
	"time"
)

type Command struct {
	Time      time.Time
	CommandId string
}

type CommandHandler interface {
	Handle(c *Command) (err error)
}
