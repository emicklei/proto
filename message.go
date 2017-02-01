package proto

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
		return p.unexpected(lit, "identifier")
	}
	m.Name = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tLEFTCURLY {
		return p.unexpected(lit, "{")
	}
	for {
		tok, lit = p.scanIgnoreWhitespace()
		switch tok {
		case tCOMMENT:
			m.Elements = append(m.Elements, p.newComment(lit))
		case tENUM:
			e := new(Enum)
			if err := e.parse(p); err != nil {
				return err
			}
			m.Elements = append(m.Elements, e)
		case tMESSAGE:
			msg := new(Message)
			if err := msg.parse(p); err != nil {
				return err
			}
			m.Elements = append(m.Elements, msg)
		case tOPTION:
			o := new(Option)
			if err := o.parse(p); err != nil {
				return err
			}
			m.Elements = append(m.Elements, o)
		case tONEOF:
			o := new(Oneof)
			if err := o.parse(p); err != nil {
				return err
			}
			m.Elements = append(m.Elements, o)
		case tMAP:
			f := newMapField()
			if err := f.parse(p); err != nil {
				return err
			}
			m.Elements = append(m.Elements, f)
		case tRESERVED:
			r := new(Reserved)
			if err := r.parse(p); err != nil {
				return err
			}
			m.Elements = append(m.Elements, r)
		case tRIGHTCURLY:
			goto done
		case tSEMICOLON:
			// continue
		default:
			// tFIELD
			p.unscan()
			f := newNormalField()
			if err := f.parse(p); err != nil {
				return err
			}
			m.Elements = append(m.Elements, f)
		}
	}
done:
	if tok != tRIGHTCURLY {
		return p.unexpected(lit, "}")
	}
	return nil
}
