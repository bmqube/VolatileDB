package commands

import (
	"errors"

	"github.com/bmqube/VolatileDB/models"
	"github.com/bmqube/VolatileDB/resp"
	"github.com/bmqube/VolatileDB/store"
)

type ExistsCommand struct{}

func (cmd *ExistsCommand) Execute(args []models.Message, store *store.Storage, mb *resp.MessageBuilder) models.Message {
	if err := cmd.ValidateArgs(args); err != nil {
		return mb.Error(err.Error()).Build()
	}

	key := args[0]

	_, ok := store.Get(string(key.BulkString))
	if !ok {
		return mb.Integer(0).Build()
	}

	return mb.Integer(1).Build()
}

func (cmd *ExistsCommand) ValidateArgs(args []models.Message) error {
	if len(args) != 1 {
		return errors.New("ERR wrong number of arguments for 'exists' command")
	}

	return nil
}
