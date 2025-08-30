package commands

import (
	"errors"

	"github.com/bmqube/VolatileDB/models"
	"github.com/bmqube/VolatileDB/resp"
	"github.com/bmqube/VolatileDB/store"
)

type PingCommand struct{}

func (cmd *PingCommand) Execute(args []models.Message, store *store.Storage, mb *resp.MessageBuilder) models.Message {
	if err := cmd.ValidateArgs(args); err != nil {
		return mb.Error(err.Error()).Build()
	}

	if len(args) == 0 {
		return mb.SimpleString("PONG").Build()
	}

	return args[1]
}

func (cmd *PingCommand) ValidateArgs(args []models.Message) error {
	if len(args) > 1 {
		return errors.New("ERR wrong number of arguments for 'ping' command")
	}

	return nil
}
