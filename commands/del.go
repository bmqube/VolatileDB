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

	cnt := int64(0)

	for _, key := range args {
		if store.Del(string(key.BulkString)) {
			cnt++
		}
	}

	return mb.Integer(cnt).Build()
}

func (cmd *DelCommand) ValidateArgs(args []models.Message) error {
	if len(args) < 1 {
		return errors.New("ERR wrong number of arguments for 'del' command")
	}

	return nil
}
