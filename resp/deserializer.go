package resp

import (
	"bytes"
	"errors"
	"strconv"

	"github.com/bmqube/VolatileDB/models"
)

func Deserialize(reader *bytes.Reader) (models.Message, error) {
	op, err := reader.ReadByte()
	if err != nil {
		return models.Message{}, err
	}

	switch op {
	case '+':
		line, err := readLine(reader)
		if err != nil {
			return models.Message{}, err
		}

		return models.Message{
			DataType:     "simple_string",
			SimpleString: string(line),
		}, nil
	case '-':
		line, err := readLine(reader)
		if err != nil {
			return models.Message{}, err
		}

		return models.Message{
			DataType:     "error",
			ErrorMessage: string(line),
		}, nil
	case ':':
		line, err := readLine(reader)
		if err != nil {
			return models.Message{}, err
		}

		val, err := strconv.ParseInt(string(line), 10, 64)

		if err != nil {
			return models.Message{}, err
		}

		return models.Message{
			DataType: "int",
			Int:      val,
		}, nil
	case '$':
		line, err := readLine(reader)
		if err != nil {
			return models.Message{}, err
		}

		length, err := strconv.ParseInt(string(line), 10, 32)

		if err != nil {
			return models.Message{}, err
		}

		line, err = readLine(reader)
		if len(line) != int(length) {
			return models.Message{}, err
		}

		return models.Message{
			DataType:   "bulk_string",
			BulkString: line,
		}, err
	case '*':
		line, err := readLine(reader)
		if err != nil {
			return models.Message{}, err
		}

		length, err := strconv.ParseInt(string(line), 10, 64)

		if err != nil {
			return models.Message{}, err
		}

		messages := make([]models.Message, 0, length)

		for range int(length) {
			message, err := Deserialize(reader)

			if err != nil {
				return models.Message{}, err
			}

			messages = append(messages, message)
		}

		return models.Message{
			DataType: "array",
			Array:    messages,
		}, nil
	default:
		return models.Message{}, errors.New("Unknown Operation")
	}
}

func readLine(reader *bytes.Reader) ([]byte, error) {
	var line []byte
	for {
		ch, err := reader.ReadByte()
		if err != nil {
			return nil, err
		}
		if ch == '\r' {
			break
		}
		line = append(line, ch)
	}
	// Expecting '\n' after '\r'
	if ch, err := reader.ReadByte(); err != nil || ch != '\n' {
		return nil, err
	}
	return line, nil
}
