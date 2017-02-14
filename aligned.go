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

import "strings"

type aligned struct {
	source  string
	left    bool
	padding bool
}

var (
	alignedEquals      = leftAligned(" = ")
	alignedShortEquals = leftAligned("=")
	alignedSpace       = leftAligned(" ")
	alignedComma       = leftAligned(", ")
	alignedEmpty       = leftAligned("")
	alignedSemicolon   = leftAligned(";")
)

func leftAligned(src string) aligned  { return aligned{src, true, true} }
func rightAligned(src string) aligned { return aligned{src, false, true} }
func notAligned(src string) aligned   { return aligned{src, false, false} }

func (a aligned) preferredWidth() int { return len(a.source) }

func (a aligned) formatted(width int) string {
	if len(a.source) > width {
		return a.source[:width]
	}
	if !a.padding {
		return a.source
	}
	if a.left {
		return a.source + strings.Repeat(" ", width-len(a.source))
	}
	return strings.Repeat(" ", width-len(a.source)) + a.source
}
