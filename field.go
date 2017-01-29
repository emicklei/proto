package proto3

import (
	"fmt"
	"strconv"
)

// Field is a message field.
type Field struct {
	Name     string
	Type     string
	Repeated bool
	Sequence int
	Options  []*Option
}

// Accept dispatches the call to the visitor.
func (f *Field) Accept(v Visitor) {
	v.VisitField(f)
}

func (f *Field) parse(p *Parser) error {
	for {
		tok, lit := p.scanIgnoreWhitespace()
		switch tok {
		case tIDENT:
			f.Type = lit
			return parseNormalField(f, p)
		case tREPEATED:
			f.Repeated = true
			return f.parse(p)
		case tMAP:
			tok, lit := p.scanIgnoreWhitespace()
			if tLESS != tok {
				return p.unexpected(lit, "<")
			}
			kvtypes := p.s.scanUntil('>')
			f.Type = fmt.Sprintf("map<%s>", kvtypes)
			return parseNormalField(f, p)
		default:
			goto done
		}
	}
done:
	return nil
}

// parseNormalField expects:
// fieldName "=" fieldNumber [ "[" fieldOptions "]" ] ";
func parseNormalField(f *Field, p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != tIDENT {
		return p.unexpected(lit, "identifier")
	}
	f.Name = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tEQUALS {
		return p.unexpected(lit, "=")
	}
	_, lit = p.scanIgnoreWhitespace()
	i, err := strconv.Atoi(lit)
	if err != nil {
		return p.unexpected(lit, "sequence number")
	}
	f.Sequence = i
	// see if there are options
	tok, lit = p.scanIgnoreWhitespace()
	if tLEFTSQUARE != tok {
		p.unscan()
		return nil
	}
	for {
		o := new(Option)
		o.PartOfFieldOrEnum = true
		err := o.parse(p)
		if err != nil {
			return err
		}
		f.Options = append(f.Options, o)

		tok, lit = p.scanIgnoreWhitespace()
		if tRIGHTSQUARE == tok {
			break
		}
		if tCOMMA != tok {
			return p.unexpected(lit, ",")
		}
	}
	return nil
}
