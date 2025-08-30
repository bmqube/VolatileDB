package commands

import (
	"github.com/bmqube/VolatileDB/models"
	"github.com/bmqube/VolatileDB/resp"
	"github.com/bmqube/VolatileDB/store"
)

type Command interface {
	Execute(args []models.Message, store *store.Storage, mb *resp.MessageBuilder) models.Message
	ValidateArgs(args []models.Message) error
}
