package commands

import (
	"errors"

	"github.com/bmqube/VolatileDB/models"
	"github.com/bmqube/VolatileDB/resp"
	"github.com/bmqube/VolatileDB/store"
)

type DecrCommand struct{}

func (cmd *DecrCommand) Execute(args []models.Message, store *store.Storage, mb *resp.MessageBuilder) models.Message {
	if err := cmd.ValidateArgs(args); err != nil {
		return mb.Error(err.Error()).Build()
	}

	key := args[0]

	val, err := store.Decr(string(key.BulkString))
	if err != nil {
		return mb.Error(err.Error()).Build()
	}

	return mb.Integer(val).Build()
}

func (cmd *DecrCommand) ValidateArgs(args []models.Message) error {
	if len(args) != 1 {
		return errors.New("ERR wrong number of arguments for 'decr' command")
	}

	return nil
}
