package proto3

import (
	"fmt"
	"strconv"
)

// Enum definition consists of a name and an enum body.
type Enum struct {
	Line       int
	Name       string
	Options    []*Option
	EnumFields []*EnumField
}

// EnumField is part of the body of an Enum.
type EnumField struct {
	Name        string
	Integer     int
	ValueOption *Option
}

func (f *EnumField) parse(p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != tIDENT {
		return fmt.Errorf("found %q, expected identifier", lit)
	}
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tEQUALS {
		return fmt.Errorf("found %q, expected =", lit)
	}
	is := p.s.scanIntegerString()
	i, err := strconv.Atoi(is)
	if err != nil {
		return fmt.Errorf("found %q, expected integer", is)
	}
	f.Integer = i
	tok, lit = p.scanIgnoreWhitespace()
	if tok == tLEFTSQUARE {
		o := new(Option)
		err := o.parse(p)
		if err != nil {
			return err
		}
		f.ValueOption = o
		tok, lit = p.scanIgnoreWhitespace()
		if tok != tRIGHTSQUARE {
			return fmt.Errorf("found %q, expected ]", lit)
		}
	}
	return nil
}

func (e *Enum) parse(p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != tIDENT {
		return fmt.Errorf("found %q, expected identifier", lit)
	}
	e.Name = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tLEFTCURLY {
		return fmt.Errorf("found %q, expected {", lit)
	}
	for {
		tok, lit = p.scanIgnoreWhitespace()
		switch tok {
		case tRIGHTCURLY:
			goto done
		case tSEMICOLON:
		case tOPTION:
			v := new(Option)
			err := v.parse(p)
			if err != nil {
				return err
			}
			e.Options = append(e.Options, v)
		default:
			p.unscan()
			f := new(EnumField)
			err := f.parse(p)
			if err != nil {
				return err
			}
			e.EnumFields = append(e.EnumFields, f)
		}
	}
done:
	if tok != tRIGHTCURLY {
		return fmt.Errorf("found %q, expected }", lit)
	}
	return nil
}
