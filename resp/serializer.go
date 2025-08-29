package resp

import (
	"fmt"

	"github.com/bmqube/VolatileDB/models"
)

func Serialize(message models.Message) string {
	switch message.DataType {
	case "simple_string":
		return fmt.Sprintf("+%s\r\n", message.SimpleString)
	case "error":
		return fmt.Sprintf("-%s\r\n", message.ErrorMessage)
	case "bulk_string":
		if message.BulkString == nil {
			return "$-1\r\n"
		}
		return fmt.Sprintf("$%d\r\n%s\r\n", len(message.BulkString), string(message.BulkString))
	case "int":
		return fmt.Sprintf(":%d\r\n", message.Int)
	case "array":
		result := fmt.Sprintf("*%d\r\n", len(message.Array))
		for _, elem := range message.Array {
			result += Serialize(elem)
		}
		return result
	default:
		return ""
	}
}

func SerializeErrorMessage(errorMessage string) string {
	return fmt.Sprintf("-%s\r\n", errorMessage)
}
