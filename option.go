package proto

import "fmt"

// Option is a protoc compiler option
type Option struct {
	Name       string
	Constant   Literal
	IsEmbedded bool
	IsCustom   bool // TODO needed?
}

// Accept dispatches the call to the visitor.
func (o *Option) Accept(v Visitor) {
	v.VisitOption(o)
}

// columns returns printable source tokens
func (o *Option) columns() (cols []aligned) {
	if !o.IsEmbedded {
		cols = append(cols, leftAligned("option "))
	} else {
		cols = append(cols, leftAligned(" ["))
	}
	cols = append(cols, o.keyValuePair()...)
	if o.IsEmbedded {
		cols = append(cols, leftAligned("]"))
	}
	return
}

// keyValuePair returns key = value or "value"
func (o *Option) keyValuePair() (cols []aligned) {
	return append(cols, leftAligned(o.Name), alignedShortEquals, rightAligned(o.Constant.String()))
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
	l := new(Literal)
	if err := l.parse(p); err != nil {
		return err
	}
	o.Constant = *l
	return nil
}

// Literal represents intLit,floatLit,strLit or boolLit
type Literal struct {
	Source   string
	IsString bool
}

// String returns the source (if quoted then use double quote).
func (l Literal) String() string {
	if l.IsString {
		return "\"" + l.Source + "\""
	}
	return l.Source
}

func (l *Literal) parse(p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	// stringLiteral?
	if tok == tQUOTE {
		ident := p.s.scanUntil('"')
		if len(ident) == 0 {
			return p.unexpected(lit, "quoted string")
		}
		l.Source, l.IsString = ident, true
		return nil
	}
	// stringLiteral?
	if tok == tSINGLEQUOTE {
		ident := p.s.scanUntil('\'')
		if len(ident) == 0 {
			return p.unexpected(lit, "single quoted string")
		}
		l.Source, l.IsString = ident, true
		return nil
	}
	// float, bool or intLit ?
	if lit == "-" { // TODO token?
		_, rem := p.s.scanIdent()
		l.Source = "-" + rem
		return nil
	}
	l.Source = lit
	return nil
}
