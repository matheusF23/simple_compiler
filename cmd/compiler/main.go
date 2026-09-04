package main

import (
	"simple_compiler/interpreter"
	"simple_compiler/parser"
)

func main() {
	input := `
				let a = 42 + 2;
				let b = 15 + 3;
				print a + b;
			`

	p := parser.NewParser([]byte(input))
	p.Parse()

	i := interpreter.NewInterpreter(p.Output())
	i.Run()
}
