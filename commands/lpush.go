package commands

import (
	"errors"

	"github.com/bmqube/VolatileDB/models"
	"github.com/bmqube/VolatileDB/resp"
	"github.com/bmqube/VolatileDB/store"
)

type LPushCommand struct{}

func (cmd *LPushCommand) Execute(args []models.Message, store *store.Storage, mb *resp.MessageBuilder) models.Message {
	if err := cmd.ValidateArgs(args); err != nil {
		return mb.Error(err.Error()).Build()
	}

	key := args[0]
	values := make([]string, len(args)-1)
	for i := 1; i < len(args); i++ {
		values[i-1] = string(args[i].BulkString)
	}

	length := store.LPush(string(key.BulkString), values)

	return mb.Integer(int64(length)).Build()
}

func (cmd *LPushCommand) ValidateArgs(args []models.Message) error {
	if len(args) < 2 {
		return errors.New("ERR wrong number of arguments for 'lpush' command")
	}

	return nil
}
