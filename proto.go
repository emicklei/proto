package proto3parser

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

// parseProto parsers a complete .proto definition source.
func parseProto(proto *Proto, p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	//log.Println(tok, lit)
	switch tok {
	case COMMENT:
		proto.Comments = append(proto.Comments, &Comment{
			Line:    p.s.line - 1, // line number before EOL was seen
			Message: lit,
		})
	case SYNTAX:
		s := new(Syntax)
		if err := s.parse(p); err != nil {
			return err
		}
		proto.Syntax = s
	case IMPORT:
		im := new(Import)
		if err := im.parse(p); err != nil {
			return err
		}
		proto.Imports = append(proto.Imports, im)
	case ENUM:
		enum := new(Enum)
		if err := enum.parse(p); err != nil {
			return err
		}
		proto.Enums = append(proto.Enums, enum)
	case SERVICE:
		if service, err := parseService(p); err != nil {
			return err
		} else {
			proto.Services = append(proto.Services, service)
		}
	case MESSAGE:
		msg := new(Message)
		if err := msg.parse(p); err != nil {
			return err
		}
		proto.Messages = append(proto.Messages, msg)
	case EOF:
		return nil
	}
	return parseProto(proto, p)
}
