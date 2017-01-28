package proto3

import "fmt"

// Message consists of a message name and a message body.
type Message struct {
	Line     int
	Comments []*Comment

	Name   string
	Fields []*Field
	Enums  []*Enum
}

func (m *Message) parse(p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != tIDENT {
		return fmt.Errorf("found %q, expected identifier", lit)
	}
	m.Name = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tLEFTCURLY {
		return fmt.Errorf("found %q, expected {", lit)
	}
	for {
		tok, lit = p.scanIgnoreWhitespace()
		switch tok {
		case tCOMMENT:
			m.Comments = append(m.Comments, p.newComment(lit))
		case tRIGHTCURLY:
			goto done
		case tSEMICOLON:
		case tENUM:
			e := new(Enum)
			err := e.parse(p)
			if err != nil {
				return err
			}
			m.Enums = append(m.Enums, e)
		default:
			p.unscan()
			f := new(Field)
			err := f.parse(p)
			if err != nil {
				return err
			}
			m.Fields = append(m.Fields, f)
		}
	}
done:
	if tok != tRIGHTCURLY {
		return fmt.Errorf("found %q, expected }", lit)
	}
	return nil
}
