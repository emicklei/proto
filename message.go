package proto3parser

import (
	"bytes"
	"fmt"
	"log"
	"strings"
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
		if tok != RIGHTCURLY {
			f, err := ParseField(p)
			if err != nil {
				return nil, err
			}
			m.Fields = append(m.Fields, f)
		} else {
			// NEEDED?
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
	Type     string
	Repeated bool
	Messages []*Message
}

func ParseField(p *Parser) (*Field, error) {
	f := new(Field)
	for {
		tok, lit := p.scanIgnoreWhitespace()
		switch tok {
		case IDENT:
			// normal type?
			if strings.Contains(TypeTokens, lit) {
				f.Type = lit
				return f, ParseNormalField(f, p)
			}
		case MESSAGE:
			m, err := ParseMessage(p)
			if err != nil {
				return f, err
			}
			f.Messages = append(f.Messages, m)
		case REPEATED:
			f.Repeated = true
			return f, ParseNormalField(f, p)
		default:
			log.Println("default", tok, lit)
			goto done
		}
	}
done:
	return f, nil
}

func ParseNormalField(f *Field, p *Parser) error {
	return nil
}
