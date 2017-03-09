// Copyright (c) 2017 Ernest Micklei
//
// MIT License
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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

// addElement is part of elementContainer
func (s *Service) addElement(v Visitee) {
	s.Elements = append(s.Elements, v)
}

// elements is part of elementContainer
func (s *Service) elements() []Visitee {
	return s.Elements
}

// parse continues after reading "service"
func (s *Service) parse(p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != tIDENT {
		if !isKeyword(tok) {
			return p.unexpected(lit, "service identifier", s)
		}
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
			maybeScanInlineComment(p, s)
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
	Comment        *Comment
	Options        []*Option
}

// Accept dispatches the call to the visitor.
func (r *RPC) Accept(v Visitor) {
	v.VisitRPC(r)
}

// inlineComment is part of commentInliner.
func (r *RPC) inlineComment(c *Comment) {
	r.Comment = c
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
	cols = append(cols, alignedSemicolon)
	if r.Comment != nil {
		cols = append(cols, notAligned(" //"), notAligned(r.Comment.Message))
	}
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
	tok, lit = p.scanIgnoreWhitespace()
	if tSEMICOLON == tok {
		p.s.unread(';') // allow for inline comment parsing
		return nil
	}
	if tLEFTCURLY == tok {
		// parse options
		for {
			tok, lit = p.scanIgnoreWhitespace()
			if tRIGHTCURLY == tok {
				break
			}
			if tCOMMENT == tok {
				// TODO put comment in the next option
				continue
			}
			if tOPTION != tok {
				return p.unexpected(lit, "rpc option", r)
			}
			o := new(Option)
			if err := o.parse(p); err != nil {
				return err
			}
			r.Options = append(r.Options, o)
		}
	}
	return nil
}
