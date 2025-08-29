package commands

import (
	"github.com/bmqube/VolatileDB/models"
	"github.com/bmqube/VolatileDB/store"
)

type Command interface {
	Execute(args []models.Message, store *store.Storage) models.Message
	ValidateArgs(args []models.Message) error
}
