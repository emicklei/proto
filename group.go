package proto

// Group represents a (proto2 only) group.
// https://developers.google.com/protocol-buffers/docs/reference/proto2-spec#group_field
type Group struct {
	Name     string
	Optional bool
	Sequence int
	Elements []Visitee
}

// Accept dispatches the call to the visitor.
func (g *Group) Accept(v Visitor) {
	v.VisitGroup(g)
}

// addElement is part of elementContainer
func (g *Group) addElement(v Visitee) {
	g.Elements = append(g.Elements, v)
}

// elements is part of elementContainer
func (g *Group) elements() []Visitee {
	return g.Elements
}

// parse expects:
// groupName "=" fieldNumber { messageBody }
func (g *Group) parse(p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != tIDENT {
		if !isKeyword(tok) {
			return p.unexpected(lit, "group name", g)
		}
	}
	g.Name = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tEQUALS {
		return p.unexpected(lit, "group =", g)
	}
	i, err := p.s.scanInteger()
	if err != nil {
		return p.unexpected(lit, "group sequence number", g)
	}
	g.Sequence = i
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tLEFTCURLY {
		return p.unexpected(lit, "group opening {", g)
	}
	parseMessageBody(p, g)
	return nil
}
