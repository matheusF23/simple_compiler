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

func (s *Scanner) number() token.Token {
	start := s.current

	for unicode.IsDigit(rune(s.peek())) {
		s.advance()
	}

	n := string(s.input[start:s.current])

	return token.NewToken(token.NUMBER, n)
}

func (s *Scanner) NextToken() token.Token {
	ch := s.peek()

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

	case '\x00':
		return token.NewToken(token.EOF, "EOF")

	default:
		panic(fmt.Sprintf("lexical error at %c", ch))
	}
}
