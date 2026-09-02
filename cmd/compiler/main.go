package main

import (
	"simple_compiler/parser"
)

func main() {
	input := "8+5-7+9"

	p := parser.NewParser([]byte(input))

	p.Parse()
}
