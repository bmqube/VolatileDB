package commands

import (
	"errors"

	"github.com/bmqube/VolatileDB/models"
	"github.com/bmqube/VolatileDB/resp"
	"github.com/bmqube/VolatileDB/store"
)

type DelCommand struct{}

func (cmd *DelCommand) Execute(args []models.Message, store *store.Storage, mb *resp.MessageBuilder) models.Message {
	if err := cmd.ValidateArgs(args); err != nil {
		return mb.Error(err.Error()).Build()
	}

	var ok bool = false

	for _, key := range args {
		ok = store.Del(string(key.BulkString)) || ok
	}

	if !ok {
		return mb.Integer(0).Build()
	}

	return mb.Integer(1).Build()
}

func (cmd *DelCommand) ValidateArgs(args []models.Message) error {
	if len(args) < 1 {
		return errors.New("ERR wrong number of arguments for 'del' command")
	}

	return nil
}
