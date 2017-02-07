package proto

import "strings"

// Proto represents a .proto definition
type Proto struct {
	Elements []Visitee
}

// addElement is part of elementContainer
func (proto *Proto) addElement(v Visitee) {
	proto.Elements = append(proto.Elements, v)
}

// elements is part of elementContainer
func (proto *Proto) elements() []Visitee {
	return proto.Elements
}

// parse parsers a complete .proto definition source.
func (proto *Proto) parse(p *Parser) error {
	for {
		tok, lit := p.scanIgnoreWhitespace()
		switch tok {
		case tCOMMENT:
			proto.Elements = append(proto.Elements, p.newComment(lit))
		case tOPTION:
			o := new(Option)
			if err := o.parse(p); err != nil {
				return err
			}
			proto.Elements = append(proto.Elements, o)
		case tSYNTAX:
			s := new(Syntax)
			if err := s.parse(p); err != nil {
				return err
			}
			proto.Elements = append(proto.Elements, s)
		case tIMPORT:
			im := new(Import)
			if err := im.parse(p); err != nil {
				return err
			}
			proto.Elements = append(proto.Elements, im)
		case tENUM:
			enum := new(Enum)
			if err := enum.parse(p); err != nil {
				return err
			}
			proto.Elements = append(proto.Elements, enum)
		case tSERVICE:
			service := new(Service)
			err := service.parse(p)
			if err != nil {
				return err
			}
			proto.Elements = append(proto.Elements, service)
		case tPACKAGE:
			pkg := new(Package)
			if err := pkg.parse(p); err != nil {
				return err
			}
			proto.Elements = append(proto.Elements, pkg)
		case tMESSAGE:
			msg := new(Message)
			if err := msg.parse(p); err != nil {
				return err
			}
			proto.Elements = append(proto.Elements, msg)
		// BEGIN proto2
		case tEXTEND:
			msg := new(Message)
			msg.IsExtend = true
			if err := msg.parse(p); err != nil {
				return err
			}
			proto.Elements = append(proto.Elements, msg)
		// END proto2
		case tSEMICOLON:
			maybeScanInlineComment(p, proto)
			// continue
		case tEOF:
			goto done
		default:
			return p.unexpected(lit, ".proto element {comment|option|import|syntax|enum|service|package|message}", p)
		}
	}
done:
	return nil
}

// Comment holds a message and line number.
type Comment struct {
	Message string
}

// Accept dispatches the call to the visitor.
func (c *Comment) Accept(v Visitor) {
	v.VisitComment(c)
}

// IsMultiline returns whether its message has one or more lineends.
func (c Comment) IsMultiline() bool {
	return strings.Contains(c.Message, "\n")
}

// commentInliner is for types that can have an inline comment.
type commentInliner interface {
	inlineComment(c *Comment)
}

// elementContainer unifies types that have elements.
type elementContainer interface {
	addElement(v Visitee)
	elements() []Visitee
}

// maybeScanInlineComment tries to scan comment on the current line ; if present then set it for the last element added.
func maybeScanInlineComment(p *Parser, c elementContainer) {
	currentLine := p.s.line
	// see if there is an inline Comment
	tok, lit := p.scanIgnoreWhitespace()
	esize := len(c.elements())
	// seen comment and on same line and elements have been added
	if tCOMMENT == tok && p.s.line == currentLine+1 && esize > 0 {
		// if the last added element can have an inline comment then set it
		last := c.elements()[esize-1]
		if inliner, ok := last.(commentInliner); ok {
			// TODO skip multiline?
			inliner.inlineComment(p.newComment(lit))
		}
	} else {
		p.unscan()
	}
}
