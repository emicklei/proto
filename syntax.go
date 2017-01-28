package proto3

import "fmt"

// Syntax should have value "proto3"
type Syntax struct {
	Value string
}

func (s *Syntax) parse(p *Parser) error {
	if tok, lit := p.scanIgnoreWhitespace(); tok != EQUALS {
		return fmt.Errorf("found %q, expected EQUALS", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != QUOTE {
		return fmt.Errorf("found %q, expected QUOTE", lit)
	}
	tok, lit := p.scanIgnoreWhitespace()
	if tok != IDENT {
		return fmt.Errorf("found %q, expected string", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != QUOTE {
		return fmt.Errorf("found %q, expected QUOTE", lit)
	}
	s.Value = lit
	return nil
}
