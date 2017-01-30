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

// parse expects:
// [ "repeated" ] type fieldName "=" fieldNumber [ "[" fieldOptions "]" ] ";"
func (f *Field) parse(p *Parser) error {
	for {
		tok, lit := p.scanIgnoreWhitespace()
		switch tok {
		case tREPEATED:
			f.Repeated = true
			return f.parse(p)
		case tIDENT:
			f.Type = lit
			return parseFieldAfterType(f, p)
		default:
			goto done
		}
	}
done:
	return nil
}

// parseMapField continues after scanning "map"
func parseMapField(f *Field, p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	if tLESS != tok {
		return p.unexpected(lit, "<")
	}
	// TODO proper parsing key-value type
	kvtypes := p.s.scanUntil('>')
	f.Type = fmt.Sprintf("map<%s>", kvtypes)
	return parseFieldAfterType(f, p)
}

// parseFieldAfterType expects:
// fieldName "=" fieldNumber [ "[" fieldOptions "]" ] ";
func parseFieldAfterType(f *Field, p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != tIDENT {
		return p.unexpected(lit, "identifier")
	}
	f.Name = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tEQUALS {
		return p.unexpected(lit, "=")
	}
	lit = p.s.scanIntegerString()
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
	// consume options
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
