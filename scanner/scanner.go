package scanner

import "unicode"

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

func (s *Scanner) NextToken() string {
	ch := s.peek()

	if ch == '0' {
		s.advance()
		return string(ch)
	} else if unicode.IsDigit(rune(ch)) {
		return s.number()
	}

	switch ch {
	case '+', '-':
		s.advance()
		return string(ch)
	}

	panic("lexical error")
}

func (s *Scanner) number() string {
	start := s.current

	for unicode.IsDigit(rune(s.peek())) {
		s.advance()
	}

	return string(s.input[start:s.current])
}
