package proto3parser

import (
	"bytes"
	"fmt"
)

type Service struct {
	Line    int
	Name    string
	RPCalls []*rpcall
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

// Accept dispatches the call to the visitor.
func (s *Service) Accept(v Visitor) {
	v.VisitService(s)
}

func parseService(p *Parser) (*Service, error) {
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
			if rpc, err := parseRPC(p); err != nil {
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

type rpcall struct {
	Method      string
	RequestType string
	Streaming   bool
	ReturnsType string
}

func (r rpcall) String() string {
	return fmt.Sprintf("rpc %s (%s) returns (%s) {}", r.Method, r.RequestType, r.ReturnsType)
}

// rpc CreateAccount (CreateAccount) returns    (ServiceFault) {}
func parseRPC(p *Parser) (*rpcall, error) {
	rpc := new(rpcall)
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
