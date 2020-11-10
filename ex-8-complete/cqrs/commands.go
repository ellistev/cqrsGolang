package cqrs

import (
	"database/sql"
	"errors"
	"github.com/satori/go.uuid"
)

type Command interface {
	AggregateId() uuid.UUID
}

type CommandFactory interface {
	CreateCommand(name string) (Command, error)
	GetCommandType(Command) (string, error)
}

type CommandSender interface {
	SendCommand(cmd Command) error
}

type CommandHandler interface {
	HandleCommand(Command, *sql.DB) error
}

type CommandHandlerFunc func(Command, *sql.DB) error

func (f CommandHandlerFunc) HandleCommand(cmd Command, db *sql.DB) error {
	return f(cmd, db)
}

var CommandHandlerAlreadyRegistered = errors.New("Command handler already registered")