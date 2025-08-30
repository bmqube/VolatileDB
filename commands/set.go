package commands

import (
	"errors"

	"github.com/bmqube/VolatileDB/models"
	"github.com/bmqube/VolatileDB/resp"
	"github.com/bmqube/VolatileDB/store"
)

type SetCommand struct{}

func (cmd *SetCommand) Execute(args []models.Message, store *store.Storage, mb *resp.MessageBuilder) models.Message {
	if err := cmd.ValidateArgs(args); err != nil {
		return mb.Error(err.Error()).Build()
	}

	key := args[0]
	val := args[1]

	store.Set(string(key.BulkString), string(val.BulkString))

	return mb.SimpleString("OK").Build()
}

func (cmd *SetCommand) ValidateArgs(args []models.Message) error {
	if len(args) != 2 {
		return errors.New("ERR wrong number of arguments for 'set' command")
	}

	return nil
}
