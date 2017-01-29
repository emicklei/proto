package proto3

import "fmt"

// Message consists of a message name and a message body.
type Message struct {
	Name     string
	Elements []Visitee
}

// Accept dispatches the call to the visitor.
func (m *Message) Accept(v Visitor) {
	v.VisitMessage(m)
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
			m.Elements = append(m.Elements, p.newComment(lit))
		case tRIGHTCURLY:
			goto done
		case tSEMICOLON:
		case tENUM:
			e := new(Enum)
			err := e.parse(p)
			if err != nil {
				return err
			}
			m.Elements = append(m.Elements, e)
		case tMESSAGE:
			msg := new(Message)
			err := msg.parse(p)
			if err != nil {
				return err
			}
			m.Elements = append(m.Elements, msg)
		case tOPTION:
			o := new(Option)
			err := o.parse(p)
			if err != nil {
				return err
			}
			m.Elements = append(m.Elements, o)
		case tONEOF:
			o := new(Oneof)
			err := o.parse(p)
			if err != nil {
				return err
			}
			m.Elements = append(m.Elements, o)
		default:
			// tFIELD
			// tMAP
			p.unscan()
			f := new(Field)
			err := f.parse(p)
			if err != nil {
				return err
			}
			m.Elements = append(m.Elements, f)
		}
	}
done:
	if tok != tRIGHTCURLY {
		return fmt.Errorf("found %q, expected }", lit)
	}
	return nil
}
