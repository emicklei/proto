package proto3parser

import (
	"bytes"
	"fmt"
)

type Message struct {
	Name   string
	Fields []*Field
}

func (m Message) String() string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "message %s {\n", m.Name)
	for _, each := range m.Fields {
		fmt.Fprintln(buf, each)
	}
	fmt.Fprintf(buf, "}\n")
	return buf.String()
}

func ParseMessage(p *Parser) (*Message, error) {
	m := new(Message)
	tok, lit := p.scanIgnoreWhitespace()
	if tok != IDENT {
		return nil, fmt.Errorf("found %q, expected name", lit)
	}
	m.Name = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok != LEFTCURLY {
		return nil, fmt.Errorf("found %q, expected {", lit)
	}
	for {
		tok, lit = p.scanIgnoreWhitespace()
		// TODO rewrite this
		if tok == OPTIONAL {
			if f, err := ParseField(p); err != nil {
				return nil, err
			} else {
				f.Optional = true
				m.Fields = append(m.Fields, f)
			}
		} else if tok == REPEATED {
			if f, err := ParseField(p); err != nil {
				return nil, err
			} else {
				f.Repeated = true
				m.Fields = append(m.Fields, f)
			}
		} else {
			p.unscan()
			break
		}
	}
	tok, lit = p.scanIgnoreWhitespace()
	if tok != RIGHTCURLY {
		return nil, fmt.Errorf("found %q, expected }", lit)
	}
	return m, nil
}

type Field struct {
	Name     string
	Optional bool
	Repeated bool
}

func ParseField(p *Parser) (*Field, error) {
	return new(Field), nil
}
