package proto3parser

import "fmt"

type Option struct {
	Name     string
	Constant string
}

func ParseOption(p *Parser) (*Option, error) {
	o := new(Option)
	tok, lit := p.scanIgnoreWhitespace()
	switch tok {
	case IDENT:
		o.Name = lit
	case LEFTPAREN:
		tok, lit = p.scanIgnoreWhitespace()
		if tok != IDENT {
			return nil, fmt.Errorf("found %q, expected identifier", lit)
		}
		o.Name = lit
		tok, lit = p.scanIgnoreWhitespace()
		if tok != RIGHTPAREN {
			return nil, fmt.Errorf("found %q, expected )", lit)
		}
	default:
		return nil, fmt.Errorf("found %q, expected identifier or (", lit)
	}
	tok, lit = p.scanIgnoreWhitespace()
	if tok != EQUALS {
		return nil, fmt.Errorf("found %q, expected =", lit)
	}
	ident, err := p.scanQuotedIdent()
	if err != nil {
		return nil, err
	}
	o.Constant = ident
	return o, nil
}
