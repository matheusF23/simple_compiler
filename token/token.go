package token

import "fmt"

type TokenType int

const (
	PLUS TokenType = iota
	MINUS
	NUMBER
	EOF
)

func (t TokenType) String() string {
	switch t {
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case NUMBER:
		return "NUMBER"
	case EOF:
		return "EOF"
	default:
		return "UNKNOWN"
	}
}

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
