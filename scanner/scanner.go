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

func (s *Scanner) NextToken() byte {
	ch := s.peek()

	if unicode.IsDigit(rune(ch)) {
		s.advance()
		return ch
	}

	switch ch {
	case '+', '-':
		s.advance()
		return ch
	}

	return '\x00'
}
