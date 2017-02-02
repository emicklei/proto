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

// addElement is part of elementContainer
func (m *Message) addElement(v Visitee) {
	m.Elements = append(m.Elements, v)
}

// parse expects ident { messageBody
func (m *Message) parse(p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != tIDENT {
		if !isKeyword(tok) {
			return p.unexpected(lit, "message identifier", m)
		}
	}
	m.Name = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tLEFTCURLY {
		return p.unexpected(lit, "message opening {", m)
	}
	return parseMessageBody(p, m)
}

// parseMessageBody parses elements after {. It consumes the closing }
func parseMessageBody(p *Parser, c elementContainer) error {
	var (
		tok token
		lit string
	)
	for {
		tok, lit = p.scanIgnoreWhitespace()
		switch tok {
		case tCOMMENT:
			c.addElement(p.newComment(lit))
		case tENUM:
			e := new(Enum)
			if err := e.parse(p); err != nil {
				return err
			}
			c.addElement(e)
		case tMESSAGE:
			msg := new(Message)
			if err := msg.parse(p); err != nil {
				return err
			}
			c.addElement(msg)
		case tOPTION:
			o := new(Option)
			if err := o.parse(p); err != nil {
				return err
			}
			c.addElement(o)
		case tONEOF:
			o := new(Oneof)
			if err := o.parse(p); err != nil {
				return err
			}
			c.addElement(o)
		case tMAP:
			f := newMapField()
			if err := f.parse(p); err != nil {
				return err
			}
			c.addElement(f)
		case tRESERVED:
			r := new(Reserved)
			if err := r.parse(p); err != nil {
				return err
			}
			c.addElement(r)
		// BEGIN proto2
		case tOPTIONAL, tREPEATED:
			// look ahead
			prevTok := tok
			tok, lit = p.scanIgnoreWhitespace()
			if tGROUP == tok {
				g := new(Group)
				g.Optional = prevTok == tOPTIONAL
				if err := g.parse(p); err != nil {
					return err
				}
				c.addElement(g)
			} else {
				// not a group, will be tFIELD
				p.unscan()
				f := newNormalField()
				f.Optional = prevTok == tOPTIONAL
				f.Repeated = prevTok == tREPEATED
				if err := f.parse(p); err != nil {
					return err
				}
				c.addElement(f)
			}
		case tGROUP:
			g := new(Group)
			if err := g.parse(p); err != nil {
				return err
			}
			c.addElement(g)
		// END proto2 only
		case tRIGHTCURLY, tEOF:
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
			c.addElement(f)
		}
	}
done:
	if tok != tRIGHTCURLY {
		return p.unexpected(lit, "message|group closing }", c)
	}
	return nil
}
