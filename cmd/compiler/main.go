package main

import "simple_compiler/parser"

func main() {
	input := `
				let a = 42 + 5 - 8;
				let b = 56 + 8;
				print a + b + 6;
			`

	p := parser.NewParser([]byte(input))

	p.Parse()
}
