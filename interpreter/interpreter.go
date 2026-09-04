package interpreter

import (
	"fmt"
	"strconv"
	"strings"
)

type Interpreter struct {
	commands  []Command
	stack     []int
	variables map[string]int
}

func (i *Interpreter) push(value int) {
	i.stack = append(i.stack, value)
}

func (i *Interpreter) pop() int {
	last := len(i.stack) - 1
	value := i.stack[last]

	i.stack = i.stack[:last]

	return value
}

func NewInterpreter(input string) *Interpreter {
	lines := strings.Split(input, "\n")

	commands := make([]Command, 0)

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}

		commands = append(commands, NewCommand(line))
	}

	return &Interpreter{
		commands:  commands,
		stack:     []int{},
		variables: make(map[string]int),
	}
}

func (i *Interpreter) hasMoreCommands() bool {
	return len(i.commands) != 0
}

func (i *Interpreter) nextCommand() Command {
	command := i.commands[0]
	i.commands = i.commands[1:]

	return command
}

func (i *Interpreter) Run() {
	for i.hasMoreCommands() {
		command := i.nextCommand()

		switch command.Type {
		case ADD:
			arg2 := i.pop()
			arg1 := i.pop()
			i.push(arg1 + arg2)

		case SUB:
			arg2 := i.pop()
			arg1 := i.pop()
			i.push(arg1 - arg2)

		case PUSH:
			value, ok := i.variables[command.Arg]

			if ok {
				i.push(value)
			} else {
				value, err := strconv.Atoi(command.Arg)
				if err != nil {
					panic(err)
				}

				i.push(value)
			}

		case POP:
			value := i.pop()
			i.variables[command.Arg] = value

		case PRINT:
			value := i.pop()
			fmt.Println(value)
		}
	}
}
