package proto

import "fmt"
import "strings"

// Option is a protoc compiler option
type Option struct {
	Name       string
	Constant   Literal
	IsEmbedded bool
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
	cols = append(cols, o.keyValuePair(o.IsEmbedded)...)
	if o.IsEmbedded {
		cols = append(cols, leftAligned("]"))
	}
	return
}

// keyValuePair returns key = value or "value"
func (o *Option) keyValuePair(embedded bool) (cols []aligned) {
	equals := alignedEquals
	name := o.Name
	if strings.Contains(name, ".") {
		name = fmt.Sprintf("(%s)", name)
	}
	if embedded {
		return append(cols, notAligned(name), equals, rightAligned(o.Constant.String()))
	}
	return append(cols, rightAligned(name), equals, rightAligned(o.Constant.String()))
}

// parse reads an Option body
// ( ident | "(" fullIdent ")" ) { "." ident } "=" constant ";"
func (o *Option) parse(p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	if tLEFTPAREN == tok {
		tok, lit = p.scanIgnoreWhitespace()
		if tok != tIDENT {
			if !isKeyword(tok) {
				return p.unexpected(lit, "option full identifier", o)
			}
		}
		o.Name = lit
		tok, _ = p.scanIgnoreWhitespace()
		if tok != tRIGHTPAREN {
			return p.unexpected(lit, "full identifier closing )", o)
		}
	} else {
		// non full ident
		if tIDENT != tok {
			if !isKeyword(tok) {
				return p.unexpected(lit, "option identifier", o)
			}
		}
		o.Name = lit
	}
	tok, lit = p.scanIgnoreWhitespace()
	if tDOT == tok {
		// extend identifier
		tok, lit = p.scanIgnoreWhitespace()
		if tok != tIDENT {
			return p.unexpected(lit, "option postfix identifier", o)
		}
		o.Name = fmt.Sprintf("%s.%s", o.Name, lit)
		tok, lit = p.scanIgnoreWhitespace()
	}
	if tEQUALS != tok {
		return p.unexpected(lit, "option constant =", o)
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

// parse expects to read a literal constant after =.
func (l *Literal) parse(p *Parser) error {
	l.Source, l.IsString = p.s.scanLiteral()
	return nil
}
