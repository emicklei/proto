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

// Proto represents a .proto definition
type Proto struct {
	Elements []Visitee
}

// addElement is part of elementContainer
func (proto *Proto) addElement(v Visitee) {
	proto.Elements = append(proto.Elements, v)
}

// elements is part of elementContainer
func (proto *Proto) elements() []Visitee {
	return proto.Elements
}

// takeLastComment is part of elementContainer
// removes and returns the last element of the list if it is a Comment.
func (proto *Proto) takeLastComment() (last *Comment) {
	last, proto.Elements = takeLastComment(proto.Elements)
	return
}

// parse parsers a complete .proto definition source.
func (proto *Proto) parse(p *Parser) error {
	for {
		line, tok, lit := p.scanIgnoreWhitespace()
		switch tok {
		case tCOMMENT:
			if com := mergeOrReturnComment(proto.Elements, lit, line); com != nil { // not merged?
				proto.Elements = append(proto.Elements, com)
			}
		case tOPTION:
			o := new(Option)
			o.LineNumber = line
			o.Comment, proto.Elements = takeLastComment(proto.Elements)
			if err := o.parse(p); err != nil {
				return err
			}
			proto.Elements = append(proto.Elements, o)
		case tSYNTAX:
			s := new(Syntax)
			s.LineNumber = line
			s.Comment, proto.Elements = takeLastComment(proto.Elements)
			if err := s.parse(p); err != nil {
				return err
			}
			proto.Elements = append(proto.Elements, s)
		case tIMPORT:
			im := new(Import)
			im.LineNumber = line
			im.Comment, proto.Elements = takeLastComment(proto.Elements)
			if err := im.parse(p); err != nil {
				return err
			}
			proto.Elements = append(proto.Elements, im)
		case tENUM:
			enum := new(Enum)
			enum.LineNumber = line
			enum.Comment, proto.Elements = takeLastComment(proto.Elements)
			if err := enum.parse(p); err != nil {
				return err
			}
			proto.Elements = append(proto.Elements, enum)
		case tSERVICE:
			service := new(Service)
			service.LineNumber = line
			service.Comment, proto.Elements = takeLastComment(proto.Elements)
			err := service.parse(p)
			if err != nil {
				return err
			}
			proto.Elements = append(proto.Elements, service)
		case tPACKAGE:
			pkg := new(Package)
			pkg.LineNumber = line
			pkg.Comment, proto.Elements = takeLastComment(proto.Elements)
			if err := pkg.parse(p); err != nil {
				return err
			}
			proto.Elements = append(proto.Elements, pkg)
		case tMESSAGE:
			msg := new(Message)
			msg.LineNumber = line
			msg.Comment, proto.Elements = takeLastComment(proto.Elements)
			if err := msg.parse(p); err != nil {
				return err
			}
			proto.Elements = append(proto.Elements, msg)
		// BEGIN proto2
		case tEXTEND:
			msg := new(Message)
			msg.LineNumber = line
			msg.Comment, proto.Elements = takeLastComment(proto.Elements)
			msg.IsExtend = true
			if err := msg.parse(p); err != nil {
				return err
			}
			proto.Elements = append(proto.Elements, msg)
		// END proto2
		case tSEMICOLON:
			maybeScanInlineComment(p, proto)
			// continue
		case tEOF:
			goto done
		default:
			return p.unexpected(lit, ".proto element {comment|option|import|syntax|enum|service|package|message}", p)
		}
	}
done:
	return nil
}

// elementContainer unifies types that have elements.
type elementContainer interface {
	addElement(v Visitee)
	elements() []Visitee
	takeLastComment() *Comment
}
