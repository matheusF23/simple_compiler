package token

import "fmt"

type TokenType string

const (
	PLUS   TokenType = "PLUS"
	MINUS  TokenType = "MINUS"
	NUMBER TokenType = "NUMBER"
	IDENT  TokenType = "IDENT"
	EOF    TokenType = "EOF"
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
