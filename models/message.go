package models

import (
	"bytes"
	"fmt"
)

type Message struct {
	DataType     string // int, error, simple_string, bulk_string, array
	Int          int64
	SimpleString string
	BulkString   []byte
	Array        []Message
	ErrorMessage string
}

func (message Message) String() string {
	switch message.DataType {
	case "simple_string":
		return fmt.Sprintf("SimpleString(%q)", message.SimpleString)
	case "error":
		return fmt.Sprintf("Error(%q)", message.ErrorMessage)
	case "int":
		return fmt.Sprintf("Integer(%d)", message.Int)
	case "bulk_string":
		if message.BulkString == nil {
			return "BulkString(nil)"
		}
		return fmt.Sprintf("BulkString(%q)", string(message.BulkString))
	case "array":
		if message.Array == nil {
			return "BulkString(nil)"
		}
		var buf bytes.Buffer
		buf.WriteString("Array[")
		for i, el := range message.Array {
			if i > 0 {
				buf.WriteString(", ")
			}
			buf.WriteString(el.String())
		}
		buf.WriteString("]")
		return buf.String()
	default:
		return "Unknown"
	}
}
