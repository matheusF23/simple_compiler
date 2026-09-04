package main

import "simple_compiler/parser"

func main() {
	input := "let a = 42 + 5 - 8;"

	p := parser.NewParser([]byte(input))

	p.Parse()
}
