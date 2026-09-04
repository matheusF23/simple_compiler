package token

import "fmt"

type TokenType string

const (
	PLUS      TokenType = "PLUS"
	MINUS     TokenType = "MINUS"
	MULT      TokenType = "MULT"
	DIV       TokenType = "DIV"
	EQ        TokenType = "EQ"
	SEMICOLON TokenType = "SEMICOLON"

	// Literals.
	NUMBER TokenType = "NUMBER"
	IDENT  TokenType = "IDENT"

	LET   TokenType = "LET"
	PRINT TokenType = "PRINT"

	EOF TokenType = "EOF"
)

type Token struct {
	Type   TokenType
	Lexeme string
}

func NewToken(tokenType TokenType, lexeme string) Token {
	return Token{
		Type:   tokenType,
		Lexeme: lexeme,
	}
}

func (t Token) String() string {
	return fmt.Sprintf("<%s>%s</%s>", t.Type, t.Lexeme, t.Type)
}
