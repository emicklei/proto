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
	Messages []*Message
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
		case tMESSAGE: // TODO here?
			m := new(Message)
			err := m.parse(p)
			if err != nil {
				return err
			}
			f.Messages = append(f.Messages, m)
		case tREPEATED:
			f.Repeated = true
			return f.parse(p)
		case tMAP:
			tok, lit := p.scanIgnoreWhitespace()
			if tLESS != tok {
				return fmt.Errorf("found %q, expected <", lit)
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

// parseNormalField proceeds after reading the type of f.
func parseNormalField(f *Field, p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != tIDENT {
		return fmt.Errorf("found %q, expected identifier", lit)
	}
	f.Name = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tEQUALS {
		return fmt.Errorf("found %q, expected =", lit)
	}
	_, lit = p.scanIgnoreWhitespace()
	i, err := strconv.Atoi(lit)
	if err != nil {
		return fmt.Errorf("found %q, expected sequence number", lit)
	}
	f.Sequence = i
	return nil
}
