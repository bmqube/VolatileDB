package commands

import (
	"errors"

	"github.com/bmqube/VolatileDB/models"
	"github.com/bmqube/VolatileDB/resp"
	"github.com/bmqube/VolatileDB/store"
)

type IncrCommand struct{}

func (cmd *IncrCommand) Execute(args []models.Message, store *store.Storage, mb *resp.MessageBuilder) models.Message {
	if err := cmd.ValidateArgs(args); err != nil {
		return mb.Error(err.Error()).Build()
	}

	key := args[0]

	val, err := store.Incr(string(key.BulkString))
	if err != nil {
		return mb.Error(err.Error()).Build()
	}

	return mb.Integer(val).Build()
}

func (cmd *IncrCommand) ValidateArgs(args []models.Message) error {
	if len(args) != 1 {
		return errors.New("ERR wrong number of arguments for 'incr' command")
	}

	return nil
}
