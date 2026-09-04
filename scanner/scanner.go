package scanner

import (
	"fmt"
	"unicode"

	"simple_compiler/token"
)

type Scanner struct {
	input   []byte
	current int
}

var keywords = map[string]token.TokenType{
	"let":   token.LET,
	"print": token.PRINT,
}

func NewScanner(input []byte) *Scanner {
	return &Scanner{
		input: input,
	}
}

func (s *Scanner) peek() byte {
	if s.current < len(s.input) {
		return s.input[s.current]
	}

	return '\x00'
}

func (s *Scanner) advance() {
	ch := s.peek()

	if ch != '\x00' {
		s.current++
	}
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || unicode.IsDigit(rune(c))
}

func (s *Scanner) number() token.Token {
	start := s.current

	for unicode.IsDigit(rune(s.peek())) {
		s.advance()
	}

	n := string(s.input[start:s.current])

	return token.NewToken(token.NUMBER, n)
}

func (s *Scanner) identifier() token.Token {
	start := s.current

	for isAlphaNumeric(s.peek()) {
		s.advance()
	}

	id := string(s.input[start:s.current])

	typeToken, ok := keywords[id]

	if !ok {
		typeToken = token.IDENT
	}

	return token.NewToken(typeToken, id)
}

func (s *Scanner) NextToken() token.Token {
	s.skipWhitespace()

	ch := s.peek()

	if isAlpha(ch) {
		return s.identifier()
	}

	if ch == '0' {
		s.advance()

		return token.NewToken(
			token.NUMBER,
			string(ch),
		)
	} else if unicode.IsDigit(rune(ch)) {
		return s.number()
	}

	switch ch {
	case '+':
		s.advance()
		return token.NewToken(token.PLUS, "+")

	case '-':
		s.advance()
		return token.NewToken(token.MINUS, "-")

	case '*':
		s.advance()
		return token.NewToken(token.MULT, "*")

	case '/':
		s.advance()
		return token.NewToken(token.DIV, "/")

	case '=':
		s.advance()
		return token.NewToken(token.EQ, "=")

	case ';':
		s.advance()
		return token.NewToken(token.SEMICOLON, ";")

	case '\x00':
		return token.NewToken(token.EOF, "EOF")

	default:
		panic(fmt.Sprintf("lexical error at %c", ch))
	}
}

func (s *Scanner) skipWhitespace() {
	ch := s.peek()

	for unicode.IsSpace(rune(ch)) {
		s.advance()
		ch = s.peek()
	}
}
