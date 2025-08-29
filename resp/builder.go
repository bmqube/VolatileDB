package resp

import "github.com/bmqube/VolatileDB/models"

type MessageBuilder struct {
	message models.Message
}

func NewMessageBuilder() *MessageBuilder {
	return &MessageBuilder{}
}

func (b *MessageBuilder) SimpleString(value string) *MessageBuilder {
	b.message = models.Message{
		DataType:     "simple_string",
		SimpleString: value,
	}
	return b
}

func (b *MessageBuilder) Error(value string) *MessageBuilder {
	b.message = models.Message{
		DataType:     "error",
		ErrorMessage: value,
	}
	return b
}

func (b *MessageBuilder) BulkString(value []byte) *MessageBuilder {
	b.message = models.Message{
		DataType:   "bulk_string",
		BulkString: value,
	}
	return b
}

func (b *MessageBuilder) Integer(value int64) *MessageBuilder {
	b.message = models.Message{
		DataType: "int",
		Int:      value,
	}
	return b
}

func (b *MessageBuilder) Array(values []models.Message) *MessageBuilder {
	b.message = models.Message{
		DataType: "array",
		Array:    values,
	}
	return b
}

func (b *MessageBuilder) Build() models.Message {
	return b.message
}
