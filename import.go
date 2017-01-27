package proto3parser

import (
	"fmt"
	"strings"
)

type Import struct {
	Line     int
	Filename string
	Kind     string // weak, public
}

func (i *Import) parse(p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != IDENT || !strings.Contains("weak public", lit) {
		return fmt.Errorf("found %q, expected kind (weak|public)", lit)
	}
	i.Line = p.s.line
	i.Kind = lit
	name := p.s.scanUntil('\n')
	if len(name) == 0 {
		return fmt.Errorf("unexpected end of quoted string")
	}
	i.Filename = name
	return nil
}
