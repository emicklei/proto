package proto3

import "fmt"

// Option is a protoc compiler option
type Option struct {
	Name              string
	Identifier        string
	String            string
	Boolean           bool
	PartOfFieldOrEnum bool
}

// Accept dispatches the call to the visitor.
func (o *Option) Accept(v Visitor) {
	v.VisitOption(o)
}

// parse reads an Option body
// ( ident | "(" fullIdent ")" ) { "." ident } "=" constant ";"
func (o *Option) parse(p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	switch tok {
	case tIDENT:
		o.Name = lit
	case tLEFTPAREN:
		tok, lit = p.scanIgnoreWhitespace()
		if tok != tIDENT {
			return p.unexpected(lit, "identifier")
		}
		o.Name = lit
		tok, lit = p.scanIgnoreWhitespace()
		if tok != tRIGHTPAREN {
			return p.unexpected(lit, ")")
		}
	default:
		return p.unexpected(lit, "identifier or (")
	}
	tok, lit = p.scanIgnoreWhitespace()
	if tok == tDOT {
		// extend identifier
		tok, lit = p.scanIgnoreWhitespace()
		if tok != tIDENT {
			return p.unexpected(lit, "postfix identifier")
		}
		o.Name = fmt.Sprintf("%s.%s", o.Name, lit)
		tok, lit = p.scanIgnoreWhitespace()
	}
	if tok != tEQUALS {
		return p.unexpected(lit, "=")
	}
	tok, lit = p.scanIgnoreWhitespace()
	// stringLiteral?
	if tok == tQUOTE {
		ident := p.s.scanUntil('"')
		if len(ident) == 0 {
			return fmt.Errorf("unexpected end of quoted string") // TODO create constant for this
		}
		o.String = ident
		return nil
	}
	if tIDENT != tok {
		return p.unexpected(lit, "constant")
	}
	o.Identifier = lit
	return nil
}
