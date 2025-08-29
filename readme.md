# VolatileDB

A lightweight, in-memory key-value database implementation written in Go, compatible with Redis protocol (RESP).

## Features

- **In-Memory Storage**: Fast key-value operations stored in memory
- **Redis Protocol Compatible**: Uses RESP (Redis Serialization Protocol) for communication
- **Concurrent Safe**: Thread-safe operations with proper synchronization
- **Extensible Architecture**: Clean design patterns make it easy to add new commands
- **TCP Server**: Handles multiple concurrent client connections

## Supported Commands

- `PING [message]` - Test server connectivity
- `ECHO <message>` - Echo back the given message
- `SET <key> <value>` - Set a key-value pair
- `GET <key>` - Get value by key

## Quick Start

### Prerequisites

- Go 1.19 or higher

### Installation

```bash
git clone https://github.com/bmqube/VolatileDB.git
cd VolatileDB
go build -o volatiledb
```

### Running the Server

```bash
./volatiledb
```

The server will start listening on port `6969` by default.

### Connecting to the Server

You can connect using any Redis client or tools like `redis-cli`, `telnet`, or `nc`:

```bash
# Using redis-cli (if installed)
redis-cli -p 6969

# Using telnet
telnet localhost 6969

# Using netcat
nc localhost 6969
```

### Example Usage

```bash
# Connect to the server
$ redis-cli -p 6969

# Test connectivity
127.0.0.1:6969> PING
PONG

127.0.0.1:6969> PING Hello
Hello

# Echo command
127.0.0.1:6969> ECHO test
test

# Set and get values
127.0.0.1:6969> SET mykey myValue
OK

127.0.0.1:6969> GET mykey
myValue

127.0.0.1:6969> GET nonexistent
(nil)
```

## Project Structure

```
VolatileDB/
├── commands/           # Command implementations
│   ├── registry.go     # Command registry
│   ├── ping.go         # PING command
│   ├── echo.go         # ECHO command
│   ├── set.go          # SET command
│   ├── get.go          # GET command
│   └── command.go      # base command interface
├── handlers/           # Connection handling
│   └── handler.go      # TCP connection handler
├── models/             # Data models
│   └── message.go      # RESP message structures
├── resp/               # RESP protocol implementation
│   ├── serializer.go   # Message serialization
│   ├── deserializer.go # Message deserialization
│   ├── builder.go      # Message building
│   └── writer.go       # Response writing
├── store/              # Storage implementations
│   └── storage.go      # Storage interface & In-memory storage
└── main.go             # Server entry point
```

## Development

### Adding New Commands

1. Create a new command struct implementing the `Command` interface:

```go
type YourCommand struct{}

func (c *YourCommand) Execute(args []models.Message, store store.Storage) models.Message {
    // Your implementation
}

func (c *YourCommand) ValidateArgs(args []models.Message) error {
    // Validation logic
}
```

2. Register it in the command registry:

```go
registry.Register("yourcommand", &YourCommand{})
```

## Performance

- **Concurrency**: Handles multiple simultaneous connections
- **Memory**: All data stored in memory for maximum speed
- **Protocol**: Efficient RESP protocol for minimal overhead

## Limitations

- **Persistence**: Data is not persisted to disk (volatile by design)
- **Memory**: Limited by available system memory
- **Commands**: Subset of Redis commands implemented

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is open source and available under the [MIT License](LICENSE).

## Acknowledgments

- Inspired by Redis and its RESP protocol
- Built as a learning project to understand database internals