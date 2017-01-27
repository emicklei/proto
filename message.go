package proto3parser

import (
	"bytes"
	"fmt"
)

type Message struct {
	Line   int
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

func (m *Message) parse(p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != IDENT {
		return fmt.Errorf("found %q, expected name", lit)
	}
	m.Name = lit
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
		default:
			p.unscan()
			f := new(Field)
			err := parseField(f, p)
			if err != nil {
				return err
			}
			m.Fields = append(m.Fields, f)
		}
	}
done:
	if tok != RIGHTCURLY {
		return fmt.Errorf("found %q, expected }", lit)
	}
	return nil
}
