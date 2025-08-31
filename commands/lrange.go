package commands

import (
	"errors"
	"strconv"

	"github.com/bmqube/VolatileDB/models"
	"github.com/bmqube/VolatileDB/resp"
	"github.com/bmqube/VolatileDB/store"
)

type LRangeCommand struct{}

func (cmd *LRangeCommand) Execute(args []models.Message, store *store.Storage, mb *resp.MessageBuilder) models.Message {
	if err := cmd.ValidateArgs(args); err != nil {
		return mb.Error(err.Error()).Build()
	}

	key := string(args[0].BulkString)

	start, err := strconv.ParseInt(string(args[1].BulkString), 10, 64)
	if err != nil {
		return mb.Error("ERR value is not an integer or out of range").Build()
	}

	stop, err := strconv.ParseInt(string(args[2].BulkString), 10, 64)
	if err != nil {
		return mb.Error("ERR value is not an integer or out of range").Build()
	}

	vals := store.LRange(key, start, stop)

	messages := make([]models.Message, len(vals))
	for i, v := range vals {
		messages[i] = models.Message{DataType: "bulk_string", BulkString: []byte(v)}
	}

	return mb.Array(messages).Build()
}

func (cmd *LRangeCommand) ValidateArgs(args []models.Message) error {
	if len(args) != 3 {
		return errors.New("ERR wrong number of arguments for 'lrange' command")
	}

	return nil
}
