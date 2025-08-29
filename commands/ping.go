package commands

import (
	"errors"

	"github.com/bmqube/VolatileDB/models"
	"github.com/bmqube/VolatileDB/resp"
	"github.com/bmqube/VolatileDB/store"
)

type PingCommand struct{}

func (cmd *PingCommand) Execute(args []models.Message, store *store.Storage) models.Message {
	messageBuilder := resp.NewMessageBuilder()
	if err := cmd.ValidateArgs(args); err != nil {
		return messageBuilder.Error(err.Error()).Build()
	}

	if len(args) == 0 {
		return messageBuilder.SimpleString("PONG").Build()
	}

	return args[1]
}

func (cmd *PingCommand) ValidateArgs(args []models.Message) error {
	if len(args) > 1 {
		return errors.New("ERR wrong number of arguments for 'ping' command")
	}

	return nil
}
