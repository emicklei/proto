package proto3

import "fmt"

// Message consists of a message name and a message body.
type Message struct {
	Line     int
	Comments []*Comment

	Name   string
	Fields []*Field
}

func (m *Message) parse(p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != IDENT {
		return fmt.Errorf("found %q, expected identifier", lit)
	}
	m.Name = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok != LEFTCURLY {
		return fmt.Errorf("found %q, expected {", lit)
	}
	for {
		tok, lit = p.scanIgnoreWhitespace()
		switch tok {
		case COMMENT:
			m.Comments = append(m.Comments, p.newComment(lit))
		case RIGHTCURLY:
			goto done
		case SEMICOLON:
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
	if tok != RIGHTCURLY {
		return fmt.Errorf("found %q, expected }", lit)
	}
	return nil
}
