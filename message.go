// Copyright (c) 2017 Ernest Micklei
//
// MIT License
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package proto

// Message consists of a message name and a message body.
type Message struct {
	Position Position
	Comment  *Comment
	Name     string
	IsExtend bool
	Elements []Visitee
}

func (m *Message) groupName() string {
	if m.IsExtend {
		return "extend"
	}
	return "message"
}

// parse expects ident { messageBody
func (m *Message) parse(p *Parser) error {
	_, tok, lit := p.scanIgnoreWhitespace()
	if tok != tIDENT {
		if !isKeyword(tok) {
			return p.unexpected(lit, m.groupName()+" identifier", m)
		}
	}
	m.Name = lit
	_, tok, lit = p.scanIgnoreWhitespace()
	if tok != tLEFTCURLY {
		return p.unexpected(lit, m.groupName()+" opening {", m)
	}
	return parseMessageBody(p, m)
}

// parseMessageBody parses elements after {. It consumes the closing }
func parseMessageBody(p *Parser, c elementContainer) error {
	var (
		pos Position
		tok token
		lit string
	)
	for {
		pos, tok, lit = p.scanIgnoreWhitespace()
		switch tok {
		case tCOMMENT:
			if com := mergeOrReturnComment(c.elements(), lit, pos); com != nil { // not merged?
				c.addElement(com)
			}
		case tENUM:
			e := new(Enum)
			e.Position = pos
			e.Comment = c.takeLastComment()
			if err := e.parse(p); err != nil {
				return err
			}
			c.addElement(e)
		case tMESSAGE:
			msg := new(Message)
			msg.Position = pos
			msg.Comment = c.takeLastComment()
			if err := msg.parse(p); err != nil {
				return err
			}
			c.addElement(msg)
		case tOPTION:
			o := new(Option)
			o.Position = pos
			o.Comment = c.takeLastComment()
			if err := o.parse(p); err != nil {
				return err
			}
			c.addElement(o)
		case tONEOF:
			o := new(Oneof)
			o.Position = pos
			o.Comment = c.takeLastComment()
			if err := o.parse(p); err != nil {
				return err
			}
			c.addElement(o)
		case tMAP:
			f := newMapField()
			f.Position = pos
			f.Comment = c.takeLastComment()
			if err := f.parse(p); err != nil {
				return err
			}
			c.addElement(f)
		case tRESERVED:
			r := new(Reserved)
			r.Position = pos
			r.Comment = c.takeLastComment()
			if err := r.parse(p); err != nil {
				return err
			}
			c.addElement(r)
		// BEGIN proto2
		case tOPTIONAL, tREPEATED, tREQUIRED:
			// look ahead
			prevTok := tok
			_, tok, lit = p.scanIgnoreWhitespace()
			if tGROUP == tok {
				g := new(Group)
				g.Position = pos
				g.Comment = c.takeLastComment()
				g.Optional = prevTok == tOPTIONAL
				if err := g.parse(p); err != nil {
					return err
				}
				c.addElement(g)
			} else {
				// not a group, will be tFIELD
				p.unscan()
				f := newNormalField()
				f.Position = pos
				f.Comment = c.takeLastComment()
				f.Optional = prevTok == tOPTIONAL
				f.Repeated = prevTok == tREPEATED
				f.Required = prevTok == tREQUIRED
				if err := f.parse(p); err != nil {
					return err
				}
				c.addElement(f)
			}
		case tGROUP:
			g := new(Group)
			g.Position = pos
			g.Comment = c.takeLastComment()
			if err := g.parse(p); err != nil {
				return err
			}
			c.addElement(g)
		case tEXTENSIONS:
			e := new(Extensions)
			e.Position = pos
			e.Comment = c.takeLastComment()
			if err := e.parse(p); err != nil {
				return err
			}
			c.addElement(e)
		case tEXTEND:
			e := new(Message)
			e.Position = pos
			e.Comment = c.takeLastComment()
			e.IsExtend = true
			if err := e.parse(p); err != nil {
				return err
			}
			c.addElement(e)
		// END proto2 only
		case tRIGHTCURLY, tEOF:
			goto done
		case tSEMICOLON:
			maybeScanInlineComment(p, c)
			// continue
		default:
			// tFIELD
			p.unscan()
			f := newNormalField()
			f.Position = pos
			f.Comment = c.takeLastComment()
			if err := f.parse(p); err != nil {
				return err
			}
			c.addElement(f)
		}
	}
done:
	if tok != tRIGHTCURLY {
		return p.unexpected(lit, "extend|message|group closing }", c)
	}
	return nil
}

// Accept dispatches the call to the visitor.
func (m *Message) Accept(v Visitor) {
	v.VisitMessage(m)
}

// addElement is part of elementContainer
func (m *Message) addElement(v Visitee) {
	m.Elements = append(m.Elements, v)
}

// elements is part of elementContainer
func (m *Message) elements() []Visitee {
	return m.Elements
}

func (m *Message) takeLastComment() (last *Comment) {
	last, m.Elements = takeLastComment(m.Elements)
	return
}

// Doc is part of Documented
func (m *Message) Doc() *Comment {
	return m.Comment
}
