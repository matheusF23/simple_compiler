package main

import (
	"simple_compiler/interpreter"
	"simple_compiler/parser"
)

func main() {
	input := `
				let a =7 - 10 / 2;
				let b = 2 + 3 * 4;
				print b / a;
			`

	p := parser.NewParser([]byte(input))
	p.Parse()

	i := interpreter.NewInterpreter(p.Output())
	i.Run()
}
