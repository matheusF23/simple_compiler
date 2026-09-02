package parser

import (
	"fmt"
	"unicode"

	"simple_compiler/scanner"
)

type Parser struct {
	scan         *scanner.Scanner
	currentToken byte
}

func NewParser(input []byte) *Parser {
	scan := scanner.NewScanner(input)

	return &Parser{
		scan:         scan,
		currentToken: scan.NextToken(),
	}
}

func (p *Parser) nextToken() {
	p.currentToken = p.scan.NextToken()
}

func (p *Parser) Parse() {
	p.expr()
}

func (p *Parser) expr() {
	p.digit()
	p.oper()
}

func (p *Parser) match(t byte) {
	if p.currentToken == t {
		p.nextToken()
	} else {
		panic("syntax error")
	}
}

func (p *Parser) digit() {
	if unicode.IsDigit(rune(p.currentToken)) {
		fmt.Println("push", string(p.currentToken))
		p.match(p.currentToken)
	} else {
		panic("syntax error")
	}
}

func (p *Parser) oper() {
	switch p.currentToken {
	case '+':
		p.match('+')
		p.digit()
		fmt.Println("add")
		p.oper()
	case '-':
		p.match('-')
		p.digit()
		fmt.Println("sub")
		p.oper()
	}
}
