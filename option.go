package proto3

import "fmt"

// Option is a protoc compiler option
type Option struct {
	Line    int
	Name    string
	String  string
	Boolean bool
}

// accept dispatches the call to the visitor.
func (o *Option) accept(v Visitor) {
	v.VisitOption(o)
}

func (o *Option) parse(p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	switch tok {
	case IDENT:
		o.Line = p.s.line
		o.Name = lit
	case LEFTPAREN:
		tok, lit = p.scanIgnoreWhitespace()
		if tok != IDENT {
			return fmt.Errorf("found %q, expected identifier", lit)
		}
		o.Name = lit
		tok, lit = p.scanIgnoreWhitespace()
		if tok != RIGHTPAREN {
			return fmt.Errorf("found %q, expected )", lit)
		}
	default:
		return fmt.Errorf("found %q, expected identifier or (", lit)
	}
	tok, lit = p.scanIgnoreWhitespace()
	if tok != EQUALS {
		return fmt.Errorf("found %q, expected =", lit)
	}
	tok, lit = p.scanIgnoreWhitespace()
	if tok == QUOTE {
		ident := p.s.scanUntil('"')
		if len(ident) == 0 {
			return fmt.Errorf("unexpected end of quoted string") // TODO create constant for this
		}
		o.String = ident
		return nil
	}
	if TRUE == tok || FALSE == tok {
		o.Boolean = lit == "true"
	} else {
		return fmt.Errorf("found %q, expected true or false", lit)
	}
	return nil
}
