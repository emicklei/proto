package proto3

import (
	"fmt"
	"strconv"
)

// Oneof is a field alternate.
type Oneof struct {
	Name     string
	Elements []Visitee
}

func (o *Oneof) parse(p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != tIDENT {
		return fmt.Errorf("found %q, expected identifier", lit)
	}
	o.Name = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tLEFTCURLY {
		return fmt.Errorf("found %q, expected {", lit)
	}
	for {
		tok, lit := p.scanIgnoreWhitespace()
		if tRIGHTCURLY == tok {
			break
		}
		if tIDENT == tok {
			f := new(OneOfField)
			f.Type = lit
			err := f.parse(p)
			if err != nil {
				return err
			}
			o.Elements = append(o.Elements, f)
		}
	}
	return nil
}

// Accept dispatches the call to the visitor.
func (o *Oneof) Accept(v Visitor) {
	v.VisitOneof(o)
}

// OneOfField is part of Oneof.
type OneOfField struct {
	Name     string
	Type     string
	Sequence int
	Options  []*Option
}

// Accept dispatches the call to the visitor.
func (o *OneOfField) Accept(v Visitor) {
	v.VisitOneofField(o)
}

func (o *OneOfField) parse(p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != tIDENT {
		return fmt.Errorf("found %q, expected identifier", lit)
	}
	o.Name = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tEQUALS {
		return p.unexpected(lit, "=")
	}
	_, lit = p.scanIgnoreWhitespace()
	i, err := strconv.Atoi(lit)
	if err != nil {
		return fmt.Errorf("found %q, expected sequence number", lit)
	}
	o.Sequence = i
	tok, lit = p.scanIgnoreWhitespace()
	if tLEFTSQUARE == tok {
		// TODO
		p.s.scanUntil(']')
	} else {
		p.unscan()
	}
	return nil
}
