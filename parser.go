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
	"errors"
	"fmt"
	"io"
	"runtime"
	"strconv"
	"text/scanner"
)

var startPosition = scanner.Position{Line: 1, Column: 1}

// Parser represents a parser.
type Parser struct {
	debug   bool
	scanner *scanner.Scanner
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	s := new(scanner.Scanner)
	s.Init(r)
	s.Mode = scanner.ScanIdents | scanner.ScanFloats | scanner.ScanStrings | scanner.ScanRawStrings | scanner.ScanComments
	return &Parser{scanner: s}
}

// Parse parses a proto definition.
func (p *Parser) Parse() (*Proto, error) {
	proto := new(Proto)
	return proto, proto.parse(p)
}

// Next returns the next token using the scanner.
func (p *Parser) next() (pos scanner.Position, tok token, lit string) {
	ch := p.scanner.Scan()
	if ch == scanner.EOF {
		return p.scanner.Position, tEOF, ""
	}
	lit = p.scanner.TokenText()
	return p.scanner.Position, asToken(lit), lit
}

func (p *Parser) unexpected(found, expected string, obj interface{}) error {
	debug := ""
	if p.debug {
		_, file, line, _ := runtime.Caller(1)
		debug = fmt.Sprintf(" at %s:%d (with %#v)", file, line, obj)
	}
	return fmt.Errorf("found %q on %v, expected [%s]%s", found, p.scanner.Position, expected, debug)
}

func (p *Parser) nextInteger() (i int, err error) {
	_, tok, lit := p.next()
	if tok != tIDENT {
		return -1, errors.New("non integer") // TODO
	}
	i, err = strconv.Atoi(lit)
	return
}

func (p *Parser) peekNonWhitespace() rune {
	r := p.scanner.Peek()
	if r == scanner.EOF {
		return r
	}
	if isWhitespace(r) {
		// consume it
		p.scanner.Next()
		return p.peekNonWhitespace()
	}
	return r
}
