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
		return p.unexpected(lit, "=")
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != tQUOTE && tok != tSINGLEQUOTE {
		return p.unexpected(lit, "\" or '")
	}
	tok, lit := p.scanIgnoreWhitespace()
	if tok != tIDENT {
		return p.unexpected(lit, "proto")
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != tQUOTE && tok != tSINGLEQUOTE {
		return p.unexpected(lit, "\" or '")
	}
	s.Value = lit
	return nil
}
