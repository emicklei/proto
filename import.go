package proto

// Import holds a filename to another .proto definition.
type Import struct {
	Filename string
	Kind     string // weak, public, <empty>
}

// Accept dispatches the call to the visitor.
func (i *Import) Accept(v Visitor) {
	v.VisitImport(i)
}

func (i *Import) parse(p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	switch tok {
	case tWEAK:
		i.Kind = lit
		return i.parse(p)
	case tPUBLIC:
		i.Kind = lit
		return i.parse(p)
	case tQUOTE:
		i.Filename = p.s.scanUntil('"')
	default:
		return p.unexpected(lit, "weak|public|quoted identifier")
	}
	return nil
}
