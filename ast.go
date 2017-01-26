package proto3parser

import (
	"bytes"
	"fmt"
)

type Proto struct {
	Syntax   string
	Services []*Service
	Messages []*Message
}

func NewProto() *Proto {
	return &Proto{Services: []*Service{}}
}

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

type Service struct {
	Name    string
	RPCalls []*RPCall
}

func (s Service) String() string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "service %s {\n", s.Name)
	for _, each := range s.RPCalls {
		fmt.Fprintln(buf, each)
	}
	fmt.Fprintf(buf, "}\n")
	return buf.String()
}

func ParseService(p *Parser) (*Service, error) {
	s := new(Service)
	tok, lit := p.scanIgnoreWhitespace()
	if tok != IDENT {
		return nil, fmt.Errorf("found %q, expected string", lit)
	}
	s.Name = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok != LEFTCURLY {
		return nil, fmt.Errorf("found %q, expected {", lit)
	}
	for {
		tok, lit = p.scanIgnoreWhitespace()
		if tok == RPC {
			if rpc, err := ParseRPC(p); err != nil {
				return nil, err
			} else {
				s.RPCalls = append(s.RPCalls, rpc)
			}
		} else {
			p.unscan()
			break
		}
	}
	tok, lit = p.scanIgnoreWhitespace()
	if tok != RIGHTCURLY {
		return nil, fmt.Errorf("found %q, expected }", lit)
	}
	return s, nil
}

type RPCall struct {
	Method      string
	RequestType string
	Streaming   bool
	ReturnsType string
}

func (r RPCall) String() string {
	return fmt.Sprintf("rpc %s (%s) returns (%s) {}", r.Method, r.RequestType, r.ReturnsType)
}

// rpc CreateAccount (CreateAccount) returns    (ServiceFault) {}
func ParseRPC(p *Parser) (*RPCall, error) {
	rpc := new(RPCall)
	tok, lit := p.scanIgnoreWhitespace()
	if tok != IDENT {
		return nil, fmt.Errorf("found %q, expected method", lit)
	}
	rpc.Method = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok != LEFTPAREN {
		return nil, fmt.Errorf("found %q, expected (", lit)
	}
	tok, lit = p.scanIgnoreWhitespace()
	if tok != IDENT {
		return nil, fmt.Errorf("found %q, expected request type", lit)
	}
	rpc.RequestType = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok != RIGHTPAREN {
		return nil, fmt.Errorf("found %q, expected )", lit)
	}
	tok, lit = p.scanIgnoreWhitespace()
	if tok != RETURNS {
		return nil, fmt.Errorf("found %q, expected returns", lit)
	}
	tok, lit = p.scanIgnoreWhitespace()
	if tok != LEFTPAREN {
		return nil, fmt.Errorf("found %q, expected (", lit)
	}
	tok, lit = p.scanIgnoreWhitespace()
	if tok != IDENT {
		return nil, fmt.Errorf("found %q, expected returns type", lit)
	}
	rpc.ReturnsType = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok != RIGHTPAREN {
		return nil, fmt.Errorf("found %q, expected )", lit)
	}
	tok, lit = p.scanIgnoreWhitespace()
	if tok != LEFTCURLY {
		return nil, fmt.Errorf("found %q, expected {", lit)
	}
	tok, lit = p.scanIgnoreWhitespace()
	if tok != RIGHTCURLY {
		return nil, fmt.Errorf("found %q, expected }", lit)
	}
	return rpc, nil
}

type Message struct {
	Name string
}

func (m Message) String() string {
	return "message"
}

func ParseMessage(p *Parser) (*Message, error) {
	m := new(Message)
	return m, nil
}
