package main

import "simple_compiler/parser"

func main() {
	input := "89+508-7+99"

	p := parser.NewParser([]byte(input))

	p.Parse()
}
