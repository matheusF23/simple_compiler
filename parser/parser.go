package parser

import (
	"fmt"

	"simple_compiler/scanner"
	"simple_compiler/token"
)

type Parser struct {
	scan         *scanner.Scanner
	currentToken token.Token
}

func NewParser(input []byte) *Parser {
	scan := scanner.NewScanner(input)

	return &Parser{
		scan:         scan,
		currentToken: scan.NextToken(),
	}
}

func (p *Parser) Parse() {
	p.statements()
}

func (p *Parser) expr() {
	p.term()
	p.oper()
}

func (p *Parser) number() {
	fmt.Println("push", p.currentToken.Lexeme)
	p.match(token.NUMBER)
}

func (p *Parser) letStatement() {
	p.match(token.LET)

	id := p.currentToken.Lexeme

	p.match(token.IDENT)
	p.match(token.EQ)

	p.expr()

	fmt.Println("pop", id)

	p.match(token.SEMICOLON)
}

func (p *Parser) printStatement() {
	p.match(token.PRINT)
	p.expr()
	fmt.Println("print")
	p.match(token.SEMICOLON)
}

func (p *Parser) statement() {
	switch p.currentToken.Type {
	case token.LET:
		p.letStatement()
	case token.PRINT:
		p.printStatement()
	default:
		panic("syntax error")
	}
}

func (p *Parser) statements() {
	for p.currentToken.Type != token.EOF {
		p.statement()
	}
}

func (p *Parser) term() {
	switch p.currentToken.Type {
	case token.NUMBER:
		p.number()
	case token.IDENT:
		fmt.Println("push", p.currentToken.Lexeme)
		p.match(token.IDENT)
	default:
		panic("syntax error")
	}
}

func (p *Parser) oper() {
	switch p.currentToken.Type {
	case token.PLUS:
		p.match(token.PLUS)
		p.term()
		fmt.Println("add")
		p.oper()
	case token.MINUS:
		p.match(token.MINUS)
		p.term()
		fmt.Println("sub")
		p.oper()
	}
}

func (p *Parser) match(t token.TokenType) {
	if p.currentToken.Type == t {
		p.nextToken()
	} else {
		panic("syntax error")
	}
}

func (p *Parser) nextToken() {
	p.currentToken = p.scan.NextToken()
}
