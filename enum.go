package proto3

import (
	"fmt"
	"strconv"
)

type Enum struct {
	Line       int
	Name       string
	Options    []*Option
	EnumFields []*EnumField
}

type EnumField struct {
	Name        string
	Integer     int
	ValueOption *Option
}

func (f *EnumField) parse(p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != IDENT {
		return fmt.Errorf("found %q, expected identifier", lit)
	}
	tok, lit = p.scanIgnoreWhitespace()
	if tok != EQUALS {
		return fmt.Errorf("found %q, expected =", lit)
	}
	is := p.s.scanIntegerString()
	i, err := strconv.Atoi(is)
	if err != nil {
		return fmt.Errorf("found %q, expected integer", is)
	}
	f.Integer = i
	tok, lit = p.scanIgnoreWhitespace()
	if tok == LEFTSQUARE {
		o := new(Option)
		err := o.parse(p)
		if err != nil {
			return err
		}
		f.ValueOption = o
		tok, lit = p.scanIgnoreWhitespace()
		if tok != RIGHTSQUARE {
			return fmt.Errorf("found %q, expected ]", lit)
		}
	}
	return nil
}

func (e *Enum) parse(p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != IDENT {
		return fmt.Errorf("found %q, expected identifier", lit)
	}
	e.Name = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok != LEFTCURLY {
		return fmt.Errorf("found %q, expected {", lit)
	}
	for {
		tok, lit = p.scanIgnoreWhitespace()
		switch tok {
		case RIGHTCURLY:
			goto done
		case SEMICOLON:
		case OPTION:
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
	if tok != RIGHTCURLY {
		return fmt.Errorf("found %q, expected }", lit)
	}
	return nil
}
