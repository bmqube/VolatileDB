package resp

import (
	"net"

	"github.com/bmqube/VolatileDB/models"
)

type ResponseWriter interface {
	WriteResponse(conn net.Conn, message models.Message)
	WriteError(conn net.Conn, errorMsg string)
}

type RESPResponseWriter struct{}

func NewRESPResponseWriter() *RESPResponseWriter {
	return &RESPResponseWriter{}
}

func (responseWriter *RESPResponseWriter) WriteResponse(conn net.Conn, message models.Message) {
	conn.Write([]byte(Serialize(message)))
}

func (responseWriter *RESPResponseWriter) WriteError(conn net.Conn, errorMsg string) {
	conn.Write([]byte(SerializeErrorMessage(errorMsg)))

}
