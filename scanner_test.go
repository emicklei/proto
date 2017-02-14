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
	"testing"
)

func TestScanUntilLineEnd(t *testing.T) {
	r := strings.NewReader(`hello
world`)
	s := newScanner(r)
	v := s.scanUntil('\n')
	if got, want := v, "hello"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := s.line, 2; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestScanMultilineComment(t *testing.T) {
	r := strings.NewReader(`
	/*
	𝓢𝓱𝓸𝓾𝓵𝓭 𝔽𝕠𝕣𝕞𝕒𝕥𝕥𝕚𝕟𝕘 𝘐𝘯 𝓣𝓲𝓽𝓵𝓮𝓼 𝕭𝖊 *🅿🅴🆁🅼🅸🆃🆃🅴🅳* ?
	*/
`)
	s := newScanner(r)
	s.scanUntil('/') // consume COMMENT token
	if got, want := s.scanComment(), `
	𝓢𝓱𝓸𝓾𝓵𝓭 𝔽𝕠𝕣𝕞𝕒𝕥𝕥𝕚𝕟𝕘 𝘐𝘯 𝓣𝓲𝓽𝓵𝓮𝓼 𝕭𝖊 *🅿🅴🆁🅼🅸🆃🆃🅴🅳* ?
	`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestScanSingleLineComment(t *testing.T) {
	r := strings.NewReader(`
	// dreadful //
`)
	s := newScanner(r)
	s.scanUntil('/') // consume COMMENT token
	if got, want := s.scanComment(), ` dreadful //`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestScanIntegerString(t *testing.T) {
	r := strings.NewReader("-1234;")
	s := newScanner(r)
	i, _ := s.scanInteger()
	if got, want := i, -1234; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestScanLiteral_string(t *testing.T) {
	r := strings.NewReader(` "string" `)
	s := newScanner(r)
	v, is := s.scanLiteral()
	if got, want := v, "string"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := is, true; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

// TODO table driven
func TestScanLiteral_string2(t *testing.T) {
	r := strings.NewReader(`'string'`)
	s := newScanner(r)
	v, is := s.scanLiteral()
	if got, want := v, "string"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := is, true; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

// TODO table driven
func TestScanLiteral_float(t *testing.T) {
	r := strings.NewReader(`-3.14e10`)
	s := newScanner(r)
	v, is := s.scanLiteral()
	if got, want := v, "-3.14e10"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := is, false; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
