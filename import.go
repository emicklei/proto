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

import "fmt"

// Import holds a filename to another .proto definition.
type Import struct {
	Filename string
	Kind     string // weak, public, <empty>
	Comment  *Comment
}

func (i *Import) parse(p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	switch tok {
	case tWEAK:
		i.Kind = lit
		return i.parse(p)
	case tPUBLIC:
		i.Kind = lit
		return i.parse(p)
	case tQUOTE:
		i.Filename = p.s.scanUntil('"')
	case tSINGLEQUOTE:
		i.Filename = p.s.scanUntil('\'')
	default:
		return p.unexpected(lit, "import classifier weak|public|quoted", i)
	}
	return nil
}

// Accept dispatches the call to the visitor.
func (i *Import) Accept(v Visitor) {
	v.VisitImport(i)
}

// inlineComment is part of commentInliner.
func (i *Import) inlineComment(c *Comment) {
	i.Comment = c
}

// columns returns printable source tokens
func (i *Import) columns() (cols []aligned) {
	cols = append(cols, leftAligned("import"), alignedSpace)
	if len(i.Kind) > 0 {
		cols = append(cols, leftAligned(i.Kind), alignedSpace)
	}
	cols = append(cols, notAligned(fmt.Sprintf("%q", i.Filename)), alignedSemicolon)
	if i.Comment != nil {
		cols = append(cols, notAligned(" //"), notAligned(i.Comment.Message))
	}
	return
}
