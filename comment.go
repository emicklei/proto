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

import (
	"strings"
	"text/scanner"
)

// Comment one or more comment text lines, either in c- or c++ style.
type Comment struct {
	Position scanner.Position
	// Lines are comment text lines without prefixes //, ///, /* or suffix */
	Lines      []string
	Cstyle     bool // refers to /* ... */,  C++ style is using //
	ExtraSlash bool // is true if the comment starts with 3 slashes
}

// newComment returns a comment.
func newComment(pos scanner.Position, lit string) *Comment {
	extraSlash := strings.HasPrefix(lit, "///")
	isCstyle := strings.HasPrefix(lit, "/*") && strings.HasSuffix(lit, "*/")
	var lines []string
	if isCstyle {
		withoutMarkers := strings.TrimRight(strings.TrimLeft(lit, "/*"), "*/")
		lines = strings.Split(withoutMarkers, "\n")
	} else {
		lines = strings.Split(strings.TrimLeft(lit, "/"), "\n")
	}
	return &Comment{Position: pos, Lines: lines, Cstyle: isCstyle, ExtraSlash: extraSlash}
}

type inlineComment struct {
	line       string
	extraSlash bool
}

// Accept dispatches the call to the visitor.
func (c *Comment) Accept(v Visitor) {
	v.VisitComment(c)
}

// Merge appends all lines from the argument comment.
func (c *Comment) Merge(other *Comment) {
	c.Lines = append(c.Lines, other.Lines...)
	c.Cstyle = c.Cstyle || other.Cstyle
}

// Message returns the first line or empty if no lines.
func (c Comment) Message() string {
	if len(c.Lines) == 0 {
		return ""
	}
	return c.Lines[0]
}

// commentInliner is for types that can have an inline comment.
type commentInliner interface {
	inlineComment(c *Comment)
}

// maybeScanInlineComment tries to scan comment on the current line ; if present then set it for the last element added.
func maybeScanInlineComment(p *Parser, c elementContainer) {
	currentPos := p.scanner.Position
	// see if there is an inline Comment
	pos, tok, lit := p.next()
	esize := len(c.elements())
	// seen comment and on same line and elements have been added
	if tCOMMENT == tok && pos.Line == currentPos.Line && esize > 0 {
		// if the last added element can have an inline comment then set it
		last := c.elements()[esize-1]
		if inliner, ok := last.(commentInliner); ok {
			// TODO skip multiline?
			inliner.inlineComment(newComment(pos, lit))
		}
	} else {
		p.nextPut(pos, tok, lit)
	}
}

// takeLastComment removes and returns the last element of the list if it is a Comment.
func takeLastComment(list []Visitee) (*Comment, []Visitee) {
	if len(list) == 0 {
		return nil, list
	}
	if last, ok := list[len(list)-1].(*Comment); ok {
		return last, list[:len(list)-1]
	}
	return nil, list
}

// mergeOrReturnComment creates a new comment and tries to merge it with the last element (if is a comment and is on the next line).
func mergeOrReturnComment(elements []Visitee, lit string, pos scanner.Position) *Comment {
	com := newComment(pos, lit)
	// last element must be a comment to merge +
	// do not merge c-style comments +
	// last comment line was on previous line
	if esize := len(elements); esize > 0 {
		if last, ok := elements[esize-1].(*Comment); ok &&
			!last.Cstyle &&
			pos.Line <= last.Position.Line+len(last.Lines) { // less than because last line of file could be inline comment
			last.Merge(com)
			// mark as merged
			com = nil
		}
	}
	return com
}
