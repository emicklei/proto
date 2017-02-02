package proto

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
		return p.unexpected(lit, "service identifier", s)
	}
	s.Name = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tLEFTCURLY {
		return p.unexpected(lit, "service opening {", s)
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
			return p.unexpected(lit, "service comment|rpc", s)
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

// columns returns printable source tokens
func (r *RPC) columns() (cols []aligned) {
	cols = append(cols,
		leftAligned("rpc "),
		leftAligned(r.Name),
		leftAligned(" ("))
	if r.StreamsRequest {
		cols = append(cols, leftAligned("stream "))
	} else {
		cols = append(cols, alignedEmpty)
	}
	cols = append(cols,
		leftAligned(r.RequestType),
		leftAligned(") "),
		leftAligned("returns"),
		leftAligned(" ("))
	if r.StreamsReturns {
		cols = append(cols, leftAligned("stream "))
	} else {
		cols = append(cols, alignedEmpty)
	}
	cols = append(cols,
		leftAligned(r.ReturnsType),
		leftAligned(")"))
	return cols
}

// parse continues after reading "rpc"
func (r *RPC) parse(p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != tIDENT {
		return p.unexpected(lit, "rpc method", r)
	}
	r.Name = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tLEFTPAREN {
		return p.unexpected(lit, "rpc type opening (", r)
	}
	tok, lit = p.scanIgnoreWhitespace()
	if iSTREAM == lit {
		r.StreamsRequest = true
		tok, lit = p.scanIgnoreWhitespace()
	}
	if tok != tIDENT {
		return p.unexpected(lit, "rpc stream | request type", r)
	}
	r.RequestType = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tRIGHTPAREN {
		return p.unexpected(lit, "rpc type closing )", r)
	}
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tRETURNS {
		return p.unexpected(lit, "rpc returns", r)
	}
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tLEFTPAREN {
		return p.unexpected(lit, "rpc type opening (", r)
	}
	tok, lit = p.scanIgnoreWhitespace()
	if iSTREAM == lit {
		r.StreamsReturns = true
		tok, lit = p.scanIgnoreWhitespace()
	}
	if tok != tIDENT {
		return p.unexpected(lit, "rpc stream | returns type", r)
	}
	r.ReturnsType = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tRIGHTPAREN {
		return p.unexpected(lit, "rpc type closing )", r)
	}
	return nil
}
