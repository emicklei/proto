package proto

// Syntax should have value "proto"
type Syntax struct {
	Value string
}

// Accept dispatches the call to the visitor.
func (s *Syntax) Accept(v Visitor) {
	v.VisitSyntax(s)
}

func (s *Syntax) parse(p *Parser) error {
	if tok, lit := p.scanIgnoreWhitespace(); tok != tEQUALS {
		return p.unexpected(lit, "syntax =", s)
	}
	lit, ok := p.s.scanLiteral()
	if !ok {
		return p.unexpected(lit, "syntax string constant", s)
	}
	s.Value = lit
	return nil
}
