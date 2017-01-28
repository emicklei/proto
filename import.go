package proto3

import "fmt"

// Import holds a filename to another .proto definition.
type Import struct {
	Line     int
	Filename string
	Kind     string // weak, public, <empty>
}

func (i *Import) parse(p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	i.Line = p.s.line
	switch tok {
	case WEAK:
		i.Kind = lit
		return i.parse(p)
	case PUBLIC:
		i.Kind = lit
		return i.parse(p)
	case QUOTE:
		i.Filename = p.s.scanUntil('"')
	default:
		return fmt.Errorf("found %q, expected weak|public|quoted identifier", lit)
	}
	return nil
}
