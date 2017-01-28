package proto3

import "strings"

// Proto represents a .proto definition
type Proto struct {
	Syntax   *Syntax
	Imports  []*Import
	Enums    []*Enum
	Services []*Service
	Messages []*Message
	Comments []*Comment
}

// Comment holds a message and line number.
type Comment struct {
	Line    int
	Message string
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
		proto.Comments = append(proto.Comments, p.newComment(lit))
	case tSYNTAX:
		s := new(Syntax)
		if err := s.parse(p); err != nil {
			return err
		}
		proto.Syntax = s
	case tIMPORT:
		im := new(Import)
		if err := im.parse(p); err != nil {
			return err
		}
		proto.Imports = append(proto.Imports, im)
	case tENUM:
		enum := new(Enum)
		if err := enum.parse(p); err != nil {
			return err
		}
		proto.Enums = append(proto.Enums, enum)
	case tSERVICE:
		service := new(Service)
		err := service.parse(p)
		if err != nil {
			return err
		}
		proto.Services = append(proto.Services, service)
	case tMESSAGE:
		msg := new(Message)
		if err := msg.parse(p); err != nil {
			return err
		}
		proto.Messages = append(proto.Messages, msg)
	case tEOF:
		return nil
	}
	return proto.parse(p)
}
