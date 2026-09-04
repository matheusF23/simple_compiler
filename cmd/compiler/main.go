package main

import "simple_compiler/parser"

func main() {
	input := "45  + preco - 876"

	p := parser.NewParser([]byte(input))

	p.Parse()
}
