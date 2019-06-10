package lib

import (
	"time"
)

type Command struct {
	Time *time.Time
}

type CommandHandler interface {
	Handle(c *Command) (err error)
}