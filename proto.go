package proto

import "strings"

// Proto represents a .proto definition
type Proto struct {
	Elements []Visitee
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

type elementContainer interface {
	addElement(v Visitee)
}
