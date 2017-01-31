package proto3

// Service defines a set of RPC calls.
type Service struct {
	Name     string
	Elements []Visitee
}

// Accept dispatches the call to the visitor.
func (s *Service) Accept(v Visitor) {
	v.VisitService(s)
}

// parse continues after reading "service"
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
			rpc := new(RPC)
			err := rpc.parse(p)
			if err != nil {
				return err
			}
			s.Elements = append(s.Elements, rpc)
		case tSEMICOLON:
		case tRIGHTCURLY:
			goto done
		default:
			return p.unexpected(lit, "comment|rpc|;}")
		}
	}
done:
	return nil
}

// RPC represents an rpc entry in a message.
type RPC struct {
	Name           string
	RequestType    string
	StreamsRequest bool
	ReturnsType    string
	StreamsReturns bool
}

// Accept dispatches the call to the visitor.
func (r *RPC) Accept(v Visitor) {
	v.VisitRPC(r)
}

// parse continues after reading "rpc"
func (r *RPC) parse(p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != tIDENT {
		return p.unexpected(lit, "method")
	}
	r.Name = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tLEFTPAREN {
		return p.unexpected(lit, "(")
	}
	tok, lit = p.scanIgnoreWhitespace()
	if iSTREAM == lit {
		r.StreamsRequest = true
		tok, lit = p.scanIgnoreWhitespace()
	}
	if tok != tIDENT {
		return p.unexpected(lit, "stream | request type")
	}
	r.RequestType = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tRIGHTPAREN {
		return p.unexpected(lit, ")")
	}
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tRETURNS {
		return p.unexpected(lit, "returns")
	}
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tLEFTPAREN {
		return p.unexpected(lit, "(")
	}
	tok, lit = p.scanIgnoreWhitespace()
	if iSTREAM == lit {
		r.StreamsReturns = true
		tok, lit = p.scanIgnoreWhitespace()
	}
	if tok != tIDENT {
		return p.unexpected(lit, "stream | returns type")
	}
	r.ReturnsType = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tRIGHTPAREN {
		return p.unexpected(lit, ")")
	}
	return nil
}
