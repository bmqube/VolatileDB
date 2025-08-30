package commands

import (
	"errors"

	"github.com/bmqube/VolatileDB/models"
	"github.com/bmqube/VolatileDB/resp"
	"github.com/bmqube/VolatileDB/store"
)

type EchoCommand struct{}

func (cmd *EchoCommand) Execute(args []models.Message, store *store.Storage, mb *resp.MessageBuilder) models.Message {
	if err := cmd.ValidateArgs(args); err != nil {
		return mb.Error(err.Error()).Build()
	}

	message := args[0]

	return mb.BulkString(message.BulkString).Build()
}

func (cmd *EchoCommand) ValidateArgs(args []models.Message) error {
	if len(args) != 1 {
		return errors.New("ERR wrong number of arguments for 'echo' command")
	}

	return nil
}
