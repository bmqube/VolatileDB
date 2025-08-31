package handlers

import (
	"bytes"
	"net"
	"strings"

	"github.com/bmqube/VolatileDB/commands"
	"github.com/bmqube/VolatileDB/models"
	"github.com/bmqube/VolatileDB/resp"
	"github.com/bmqube/VolatileDB/store"
)

type ConnectionHandler struct {
	store           *store.Storage
	responseWriter  resp.ResponseWriter
	commandRegistry *commands.CommandRegistry
}

func NewConnectionHandler(memStore *store.Storage) *ConnectionHandler {
	return &ConnectionHandler{
		store:           memStore,
		responseWriter:  resp.NewRESPResponseWriter(),
		commandRegistry: commands.NewCommandRegistry(),
	}
}

func (handler *ConnectionHandler) Handle(conn net.Conn) {
	defer conn.Close()

	for {
		if err := handler.processRequest(conn); err != nil {
			break
		}
	}
}

func (handler *ConnectionHandler) processRequest(conn net.Conn) error {
	mb := resp.NewMessageBuilder()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return err
	}

	message, err := resp.Deserialize(bytes.NewReader(buffer[:n]))
	if err != nil {
		return err
	}

	// println(len(message.Array))
	// println(message.String())
	response := handler.executeCommand(message, mb)
	handler.responseWriter.WriteResponse(conn, response)

	return nil
}

func (handler *ConnectionHandler) executeCommand(message models.Message, mb *resp.MessageBuilder) models.Message {
	if message.DataType != "array" || len(message.Array) == 0 {
		return mb.Error("ERR Invalid expression").Build()
	}

	temp := message.Array[0]
	if temp.DataType != "bulk_string" || len(temp.BulkString) == 0 {
		return mb.Error("ERR Invalid command").Build()
	}

	command := strings.ToLower(string(temp.BulkString))
	executor, ok := handler.commandRegistry.Get(command)

	args := message.Array[1:]

	if !ok {
		stringifiedArgs := ""
		for _, b := range args {
			stringifiedArgs += "'" + string(b.BulkString) + "' "
		}

		if stringifiedArgs == "" {
			stringifiedArgs = command
		}

		return mb.Error("ERR unknown command '" + command + "', with args beginning with: " + stringifiedArgs).Build()
	}

	if err := executor.ValidateArgs(args); err != nil {
		return mb.Error(err.Error()).Build()
	}

	return executor.Execute(args, handler.store, mb)
}
