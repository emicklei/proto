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

import "strconv"

// Oneof is a field alternate.
type Oneof struct {
	Comment  *Comment
	Name     string
	Elements []Visitee
}

// addElement is part of elementContainer
func (o *Oneof) addElement(v Visitee) {
	o.Elements = append(o.Elements, v)
}

// elements is part of elementContainer
func (o *Oneof) elements() []Visitee {
	return o.Elements
}

// takeLastComment is part of elementContainer
// removes and returns the last element of the list if it is a Comment.
func (o *Oneof) takeLastComment() (last *Comment) {
	last, o.Elements = takeLastComment(o.Elements)
	return last
}

// parse expects:
// oneofName "{" { oneofField | emptyStatement } "}"
func (o *Oneof) parse(p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != tIDENT {
		if !isKeyword(tok) {
			return p.unexpected(lit, "oneof identifier", o)
		}
	}
	o.Name = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tLEFTCURLY {
		return p.unexpected(lit, "oneof opening {", o)
	}
	for {
		tok, lit = p.scanIgnoreWhitespace()
		switch tok {
		case tIDENT:
			f := newOneOfField()
			f.Type = lit
			if err := parseFieldAfterType(f.Field, p); err != nil {
				return err
			}
			o.Elements = append(o.Elements, f)
		case tGROUP:
			g := new(Group)
			if err := g.parse(p); err != nil {
				return err
			}
			o.Elements = append(o.Elements, g)
		case tSEMICOLON:
			maybeScanInlineComment(p, o)
			// continue
		default:
			goto done
		}
	}
done:
	if tok != tRIGHTCURLY {
		return p.unexpected(lit, "oneof closing }", o)
	}
	return nil
}

// Accept dispatches the call to the visitor.
func (o *Oneof) Accept(v Visitor) {
	v.VisitOneof(o)
}

// OneOfField is part of Oneof.
type OneOfField struct {
	*Field
}

func newOneOfField() *OneOfField { return &OneOfField{Field: new(Field)} }

// Accept dispatches the call to the visitor.
func (o *OneOfField) Accept(v Visitor) {
	v.VisitOneofField(o)
}

// columns returns printable source tokens
func (o *OneOfField) columns() (cols []aligned) {
	cols = append(cols,
		rightAligned(o.Type),
		alignedSpace,
		leftAligned(o.Name),
		alignedEquals,
		rightAligned(strconv.Itoa(o.Sequence)))
	if len(o.Options) > 0 {
		cols = append(cols, leftAligned(" ["))
		for i, each := range o.Options {
			if i > 0 {
				cols = append(cols, alignedComma)
			}
			cols = append(cols, each.keyValuePair(true)...)
		}
		cols = append(cols, leftAligned("]"))
	}
	cols = append(cols, alignedSemicolon)
	if o.InlineComment != nil {
		cols = append(cols, notAligned(" //"), notAligned(o.InlineComment.Message()))
	}
	return
}
