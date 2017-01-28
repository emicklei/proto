package proto3

import "fmt"

// Syntax should have value "proto3"
type Syntax struct {
	Value string
}

// Accept dispatches the call to the visitor.
func (s *Syntax) Accept(v Visitor) {
	v.VisitSyntax(s)
}

func (s *Syntax) parse(p *Parser) error {
	if tok, lit := p.scanIgnoreWhitespace(); tok != tEQUALS {
		return fmt.Errorf("found %q, expected EQUALS", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != tQUOTE {
		return fmt.Errorf("found %q, expected QUOTE", lit)
	}
	tok, lit := p.scanIgnoreWhitespace()
	if tok != tIDENT {
		return fmt.Errorf("found %q, expected string", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != tQUOTE {
		return fmt.Errorf("found %q, expected QUOTE", lit)
	}
	s.Value = lit
	return nil
}
