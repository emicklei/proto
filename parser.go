package proto3parser

import "io"

// Parser represents a parser.
type Parser struct {
	s   *Scanner
	buf struct {
		tok Token  // last read token
		lit string // last read literal
		n   int    // buffer size (max=1)
	}
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

// Parse parses a proto definition.
func (p *Parser) Parse() (*Proto, error) {
	proto := new(Proto)
	tok, _ := p.scanIgnoreWhitespace()
	switch tok {
	case SYNTAX:
		if syntax, err := ParseSyntax(p); err != nil {
			return nil, err
		} else {
			proto.Syntax = syntax
		}
	case SERVICE:
		if service, err := ParseService(p); err != nil {
			return nil, err
		} else {
			proto.Services = append(proto.Services, service)
		}
	case MESSAGE:
		if msg, err := ParseMessage(p); err != nil {
			return nil, err
		} else {
			proto.Messages = append(proto.Messages, msg)
		}
	}
	return proto, nil
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() { p.buf.n = 1 }
