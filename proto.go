package proto3parser

import "fmt"

type Proto struct {
	Syntax   string
	Services []*Service
	Messages []*Message
}

func ParseSyntax(p *Parser) (string, error) {
	if tok, lit := p.scanIgnoreWhitespace(); tok != EQUALS {
		return "", fmt.Errorf("found %q, expected EQUALS", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != QUOTE {
		return "", fmt.Errorf("found %q, expected QUOTE", lit)
	}
	tok, lit := p.scanIgnoreWhitespace()
	if tok != IDENT {
		return "", fmt.Errorf("found %q, expected string", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != QUOTE {
		return "", fmt.Errorf("found %q, expected QUOTE", lit)
	}
	return lit, nil
}
