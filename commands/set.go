package commands

import (
	"errors"
	"strconv"
	"strings"
	"time"

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

	timeToLive := time.Duration(0)
	for i := 2; i < len(args); i++ {
		arg := args[i]

		comm := strings.ToLower(string(arg.BulkString))

		switch comm {
		case "ex", "px":
			if timeToLive > time.Duration(0) {
				return mb.Error("ERR syntax error").Build()
			}

			if i+1 > len(args) {
				return mb.Error("ERR syntax error").Build()
			}

			ttl, err := strconv.ParseInt(string(args[i+1].BulkString), 10, 64)

			if err != nil {
				return mb.Error("ERR value is not an integer or out of range").Build()
			}

			if ttl <= 0 {
				return mb.Error("ERR value is not an integer or out of range").Build()
			}

			i++
			if comm == "ex" {
				timeToLive = time.Duration(ttl) * time.Second
			} else {
				timeToLive = time.Duration(ttl) * time.Millisecond
			}
		default:
			return mb.Error("ERR syntax error").Build()
		}
	}

	store.Set(string(key.BulkString), string(val.BulkString), timeToLive)

	return mb.SimpleString("OK").Build()
}

func (cmd *SetCommand) ValidateArgs(args []models.Message) error {
	if len(args) < 2 {
		return errors.New("ERR wrong number of arguments for 'set' command")
	}

	return nil
}
