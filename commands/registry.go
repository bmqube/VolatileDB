package commands

type CommandRegistry struct {
	commands map[string]Command
}

func NewCommandRegistry() *CommandRegistry {
	registry := &CommandRegistry{
		commands: make(map[string]Command),
	}

	// Register all commands
	registry.Register("ping", &PingCommand{})
	registry.Register("echo", &EchoCommand{})
	registry.Register("set", &SetCommand{})
	registry.Register("get", &GetCommand{})
	registry.Register("get", &GetCommand{})

	return registry
}

func (r *CommandRegistry) Register(name string, cmd Command) {
	r.commands[name] = cmd
}

func (r *CommandRegistry) Get(name string) (Command, bool) {
	cmd, exists := r.commands[name]
	return cmd, exists
}
