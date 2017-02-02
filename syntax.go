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
	lit, err := p.scanStringLiteral()
	if err != nil {
		return err
	}
	s.Value = lit
	return nil
}
