// Copyright (c) 2019 Ernest Micklei
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

import "testing"

func TestUnQuoteCases(t *testing.T) {
	singleQuoteRune := rune('\'')
	for i, each := range []struct {
		input, output string
		quoteRune     rune
	}{
		{"thanos", "thanos", doubleQuoteRune},
		{"`bucky`", "`bucky`", doubleQuoteRune},
		{"'nat", "'nat", doubleQuoteRune},
		{"'bruce'", "bruce", singleQuoteRune},
		{"\"tony\"", "tony", doubleQuoteRune},
		{"\"'\"\"'  -> \"\"\"\"\"\"", `'""'  -> """""`, doubleQuoteRune},
		{`"''"`, "''", doubleQuoteRune},
		{"''", "", singleQuoteRune},
		{"", "", doubleQuoteRune},
	} {
		got, gotRune := unQuote(each.input)
		if gotRune != each.quoteRune {
			t.Errorf("[%d] got [%v] want [%v]", i, gotRune, each.quoteRune)
		}
		want := each.output
		if got != want {
			t.Errorf("[%d] got [%s] want [%s]", i, got, want)
		}
	}
}

func TestIsNumber(t *testing.T) {
	for i, each := range []struct {
		input    string
		isNumber bool
	}{
		{`1`, true},
		{`1.2`, true},
		{`-1.02`, true},
		{`a1`, false},
		{`0x12`, true},
		{`0X77777`, true},
	} {
		got := isNumber(each.input)
		if got != each.isNumber {
			t.Errorf("[%d] got [%v] want [%v]", i, got, each.isNumber)
		}
	}
}
