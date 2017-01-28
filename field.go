package proto3

import (
	"fmt"
	"strconv"
	"strings"
)

// Field is a message field.
type Field struct {
	Name     string
	Type     string
	Repeated bool
	Messages []*Message
	Sequence int
}

func (f *Field) parse(p *Parser) error {
	for {
		tok, lit := p.scanIgnoreWhitespace()
		switch tok {
		case tIDENT:
			// normal type?
			if strings.Contains(typeTokens, lit) {
				f.Type = lit
				return parseNormalField(f, p)
			}
			//if tok == ONEOF {}
			//if tok == ONEOFFIELD {}
		case tMESSAGE:
			m := new(Message)
			err := m.parse(p)
			if err != nil {
				return err
			}
			f.Messages = append(f.Messages, m)
		case tREPEATED:
			f.Repeated = true
			return f.parse(p)
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
