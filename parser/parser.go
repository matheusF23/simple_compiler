package parser

import (
	"fmt"
	"unicode"
)

type Parser struct {
	input   []byte
	current int
}

func NewParser(input []byte) *Parser {
	return &Parser{
		input: input,
	}
}

func (p *Parser) Parse() {
	p.expr()
}

func (p *Parser) peek() byte {
	if p.current < len(p.input) {
		return p.input[p.current]
	}

	return '\x00'
}

func (p *Parser) match(c byte) {
	if c == p.peek() {
		p.current++
	} else {
		panic("syntax error")
	}
}

func (p *Parser) expr() {
	p.digit()
	p.oper()
}

func (p *Parser) digit() {
	if unicode.IsDigit(rune(p.peek())) {
		fmt.Println("push", string(p.peek()))
		p.match(p.peek())
	} else {
		panic("syntax error")
	}
}

func (p *Parser) oper() {
	if p.peek() == '+' {
		p.match('+')
		p.digit()
		fmt.Println("add")
		p.oper()
	} else if p.peek() == '-' {
		p.match('-')
		p.digit()
		fmt.Println("sub")
		p.oper()
	}
}
