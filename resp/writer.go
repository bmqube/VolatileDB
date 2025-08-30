package resp

import (
	"fmt"
	"io"

	"github.com/bmqube/VolatileDB/models"
)

type ResponseWriter interface {
	WriteResponse(writer io.Writer, message models.Message)
	WriteError(writer io.Writer, errorMsg string)
}

type RESPResponseWriter struct{}

func NewRESPResponseWriter() *RESPResponseWriter {
	return &RESPResponseWriter{}
}

func (responseWriter *RESPResponseWriter) WriteResponse(writer io.Writer, message models.Message) {
	fmt.Fprint(writer, Serialize(message))
}

func (responseWriter *RESPResponseWriter) WriteError(writer io.Writer, errorMsg string) {
	fmt.Fprint(writer, SerializeErrorMessage(errorMsg))
}
