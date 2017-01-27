package proto3parser

import "fmt"

type Proto struct {
	Syntax   string
	Services []*Service
	Messages []*Message
	Comments []*Comment
}

// ParseSyntax returns the syntax value. Parser has seen "syntax".
func ParseSyntax(p *Parser) (string, error) {
	if tok, lit := p.scanIgnoreWhitespace(); tok != EQUALS {
		return "", fmt.Errorf("found %q, expected EQUALS", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != QUOTE {
		return "", fmt.Errorf("found %q, expected QUOTE", lit)
	}
	tok, lit := p.scanIgnoreWhitespace()
	if tok != IDENT {
		return "", fmt.Errorf("found %q, expected string", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != QUOTE {
		return "", fmt.Errorf("found %q, expected QUOTE", lit)
	}
	return lit, nil
}

// Comment holds a message and line number.
type Comment struct {
	Line    int
	Message string
}

// ParseProto parsers a complete .proto definition source.
func ParseProto(proto *Proto, p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	//log.Println(tok, lit)
	switch tok {
	case COMMENT:
		proto.Comments = append(proto.Comments, &Comment{
			Line:    p.s.Line() - 1, // line number before EOL was seen
			Message: lit,
		})
	case SYNTAX:
		if syntax, err := ParseSyntax(p); err != nil {
			return err
		} else {
			proto.Syntax = syntax
		}
	case SERVICE:
		if service, err := ParseService(p); err != nil {
			return err
		} else {
			proto.Services = append(proto.Services, service)
		}
	case MESSAGE:
		if msg, err := ParseMessage(p); err != nil {
			return err
		} else {
			proto.Messages = append(proto.Messages, msg)
		}
	case EOF:
		return nil
	}
	return ParseProto(proto, p)
}
