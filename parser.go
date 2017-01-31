package proto3

import (
	"fmt"
	"io"
	"runtime"
)

// Parser represents a parser.
type Parser struct {
	s   *scanner
	buf struct {
		tok token  // last read token
		lit string // last read literal
		n   int    // buffer size (max=1)
	}
	debug bool
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: newScanner(r)}
}

// Parse parses a proto definition.
func (p *Parser) Parse() (*Proto, error) {
	proto := new(Proto)
	return proto, proto.parse(p)
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok token, lit string) {
	tok, lit = p.scan()
	if tok == tWS {
		tok, lit = p.scan()
	}
	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() { p.buf.n = 1 }

// newComment returns a comment with line indication.
func (p *Parser) newComment(lit string) *Comment {
	return &Comment{Message: lit}
}

func (p *Parser) unexpected(found, expected string) error {
	debug := ""
	if p.debug {
		_, file, line, _ := runtime.Caller(1)
		debug = fmt.Sprintf(" at %s:%d", file, line)
	}
	return fmt.Errorf("found %q on line %d, expected %s%s", found, p.s.line, expected, debug)
}
