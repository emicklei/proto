package proto3

import "fmt"

// Service defines a set of RPC calls.
type Service struct {
	Name     string
	Elements []Visitee
}

// Accept dispatches the call to the visitor.
func (s *Service) Accept(v Visitor) {
	v.VisitService(s)
}

func (s *Service) parse(p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != tIDENT {
		return p.unexpected(lit, "identifier")
	}
	s.Name = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tLEFTCURLY {
		return p.unexpected(lit, "{")
	}
	for {
		tok, lit = p.scanIgnoreWhitespace()
		switch tok {
		case tCOMMENT:
			s.Elements = append(s.Elements, p.newComment(lit))
		case tRPC:
			rpc := new(RPcall)
			err := rpc.parse(p)
			if err != nil {
				return err
			}
			s.Elements = append(s.Elements, rpc)
		case tSEMICOLON:
			goto done
		default:
			return p.unexpected(lit, "comment|rpc|;")
		}
	}
done:
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tRIGHTCURLY {
		return p.unexpected(lit, "}")
	}
	return nil
}

// RPcall represents an rpc entry in a message.
type RPcall struct {
	Name        string
	RequestType string
	Streaming   bool
	ReturnsType string
}

// Accept dispatches the call to the visitor.
func (r *RPcall) Accept(v Visitor) {
	v.VisitRPcall(r)
}

func (r *RPcall) parse(p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != tIDENT {
		return fmt.Errorf("found %q, expected method", lit)
	}
	r.Name = lit
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
