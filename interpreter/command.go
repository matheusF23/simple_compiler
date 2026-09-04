package interpreter

import (
	"strings"
)

type CommandType string

const (
	ADD   CommandType = "ADD"
	SUB   CommandType = "SUB"
	PUSH  CommandType = "PUSH"
	POP   CommandType = "POP"
	PRINT CommandType = "PRINT"
)

type Command struct {
	Type CommandType
	Arg  string
}

func NewCommand(command string) Command {
	parts := strings.Fields(command)

	cmd := Command{
		Type: CommandType(strings.ToUpper(parts[0])),
	}

	if len(parts) > 1 {
		cmd.Arg = parts[1]
	}

	return cmd
}

func (c Command) String() string {
	return string(c.Type) + " " + c.Arg
}
