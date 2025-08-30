package commands

import (
	"errors"

	"github.com/bmqube/VolatileDB/models"
	"github.com/bmqube/VolatileDB/resp"
	"github.com/bmqube/VolatileDB/store"
)

type GetCommand struct{}

func (cmd *GetCommand) Execute(args []models.Message, store *store.Storage, mb *resp.MessageBuilder) models.Message {
	if err := cmd.ValidateArgs(args); err != nil {
		return mb.Error(err.Error()).Build()
	}

	key := args[0]

	val, _ := store.Get(string(key.BulkString))

	return mb.BulkString([]byte(val)).Build()
}

func (cmd *GetCommand) ValidateArgs(args []models.Message) error {
	if len(args) != 1 {
		return errors.New("ERR wrong number of arguments for 'get' command")
	}

	return nil
}
