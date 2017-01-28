package proto3

import "strings"

// Proto represents a .proto definition
type Proto struct {
	Elements []Visitee
}

// Comment holds a message and line number.
type Comment struct {
	Line    int
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

// parse parsers a complete .proto definition source.
func (proto *Proto) parse(p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	switch tok {
	case tCOMMENT:
		proto.Elements = append(proto.Elements, p.newComment(lit))
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
		err := pkg.parse(p)
		if err != nil {
			return err
		}
		proto.Elements = append(proto.Elements, pkg)
	case tMESSAGE:
		msg := new(Message)
		if err := msg.parse(p); err != nil {
			return err
		}
		proto.Elements = append(proto.Elements, msg)
	case tEOF:
		return nil
	}
	return proto.parse(p)
}
