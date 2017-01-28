package proto3

import "fmt"

// Service defines a set of RPC calls.
type Service struct {
	Line    int
	Name    string
	RPCalls []*RPcall
}

// accept dispatches the call to the visitor.
func (s *Service) accept(v Visitor) {
	v.VisitService(s)
}

func (s *Service) parse(p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != tIDENT {
		return fmt.Errorf("found %q, expected string", lit)
	}
	s.Name = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tLEFTCURLY {
		return fmt.Errorf("found %q, expected {", lit)
	}
	for {
		tok, lit = p.scanIgnoreWhitespace()
		if tok == tRPC {
			rpc := new(RPcall)
			err := rpc.parse(p)
			if err != nil {
				return err
			}
			s.RPCalls = append(s.RPCalls, rpc)
		} else {
			p.unscan()
			break
		}
	}
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tRIGHTCURLY {
		return fmt.Errorf("found %q, expected }", lit)
	}
	return nil
}

// RPcall represents an rpc entry in a message.
type RPcall struct {
	Method      string
	RequestType string
	Streaming   bool
	ReturnsType string
}

func (r *RPcall) parse(p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != tIDENT {
		return fmt.Errorf("found %q, expected method", lit)
	}
	r.Method = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tLEFTPAREN {
		return fmt.Errorf("found %q, expected (", lit)
	}
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tIDENT {
		return fmt.Errorf("found %q, expected request type", lit)
	}
	r.RequestType = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tRIGHTPAREN {
		return fmt.Errorf("found %q, expected )", lit)
	}
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tRETURNS {
		return fmt.Errorf("found %q, expected returns", lit)
	}
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tLEFTPAREN {
		return fmt.Errorf("found %q, expected (", lit)
	}
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tIDENT {
		return fmt.Errorf("found %q, expected returns type", lit)
	}
	r.ReturnsType = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tRIGHTPAREN {
		return fmt.Errorf("found %q, expected )", lit)
	}
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tLEFTCURLY {
		return fmt.Errorf("found %q, expected {", lit)
	}
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tRIGHTCURLY {
		return fmt.Errorf("found %q, expected }", lit)
	}
	return nil
}
