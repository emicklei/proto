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

// partial code from https://raw.githubusercontent.com/benbjohnson/sql-parser/master/scanner.go

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

// scanner represents a lexical scanner.
type scanner struct {
	r    *bufio.Reader
	line int
}

// newScanner returns a new instance of Scanner.
func newScanner(r io.Reader) *scanner {
	return &scanner{r: bufio.NewReader(r), line: 1}
}

// scan returns the next token and literal value.
func (s *scanner) scan() (line int, tok token, lit string) {
	// Read the next rune.
	ch := s.read()
	line = s.line
	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as an ident or reserved word.
	// If we see a slash then consume all as a comment (can be multiline)
	if isWhitespace(ch) {
		s.unread(ch)
		tok, lit = s.scanWhitespace()
		return
	} else if isLetter(ch) || ch == '_' {
		s.unread(ch)
		tok, lit = s.scanIdent()
		return
	}

	// Otherwise read the individual character.
	switch ch {
	case eof:
		return line, tEOF, ""
	case ';':
		return line, tSEMICOLON, string(ch)
	case ':':
		return line, tCOLON, string(ch)
	case '=':
		return line, tEQUALS, string(ch)
	case '"':
		return line, tQUOTE, string(ch)
	case '\'':
		return line, tSINGLEQUOTE, string(ch)
	case '(':
		return line, tLEFTPAREN, string(ch)
	case ')':
		return line, tRIGHTPAREN, string(ch)
	case '{':
		return line, tLEFTCURLY, string(ch)
	case '}':
		return line, tRIGHTCURLY, string(ch)
	case '[':
		return line, tLEFTSQUARE, string(ch)
	case ']':
		return line, tRIGHTSQUARE, string(ch)
	case '/':
		return line, tCOMMENT, s.scanComment()
	case '<':
		return line, tLESS, string(ch)
	case ',':
		return line, tCOMMA, string(ch)
	case '.':
		return line, tDOT, string(ch)
	case '>':
		return line, tGREATER, string(ch)
	}
	return line, tILLEGAL, string(ch)
}

// skipWhitespace consumes all whitespace until eof or a non-whitespace rune.
func (s *scanner) skipWhitespace() {
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread(ch)
			break
		}
	}
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *scanner) scanWhitespace() (tok token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread(ch)
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return tWS, buf.String()
}

// scanLiteral returns the current rune and all contiguous non-literal and whether is a string.
func (s *scanner) scanLiteral() (string, bool) {
	var ch rune
	// first skip all whitespace runes
	for {
		if ch = s.read(); ch == eof {
			return "", false
		}
		if !isWhitespace(ch) {
			break
		}
	}
	// is there a single quoted string ahead?
	if '\'' == ch {
		return s.scanUntil('\''), true
	}
	// is there a double quoted string ahead?
	if '"' == ch {
		return s.scanUntil('"'), true
	}
	if isLiteralTerminator(ch) {
		return string(ch), false
	}
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(ch)

	// Read every subsequent non-literal character into the buffer.
	// Whitespace characters , EOF and literal terminators will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if isWhitespace(ch) || isLiteralTerminator(ch) {
			s.unread(ch)
			break
		} else {
			buf.WriteRune(ch)
		}
	}
	return buf.String(), false
}

// scanInteger reads an integer representation.
func (s *scanner) scanInteger() (int, error) {
	var i int
	if _, err := fmt.Fscanf(s.r, "%d", &i); err != nil {
		return i, err
	}
	return i, nil
}

// scanIdent consumes the current rune and all contiguous ident runes.
func (s *scanner) scanIdent() (tok token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) && ch != '_' && ch != '.' { // underscore and dot can be part of identifier
			s.unread(ch)
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// If the string matches a keyword then return that keyword.
	ident := buf.String()
	switch ident {
	case "syntax":
		return tSYNTAX, buf.String()
	case "service":
		return tSERVICE, buf.String()
	case "message":
		return tMESSAGE, buf.String()
	case "rpc":
		return tRPC, buf.String()
	case "returns":
		return tRETURNS, buf.String()
	case "import":
		return tIMPORT, buf.String()
	case "package":
		return tPACKAGE, buf.String()
	case "repeated":
		return tREPEATED, buf.String()
	case "option":
		return tOPTION, buf.String()
	case "enum":
		return tENUM, buf.String()
	case "weak":
		return tWEAK, buf.String()
	case "public":
		return tPUBLIC, buf.String()
	case "map":
		return tMAP, buf.String()
	case "oneof":
		return tONEOF, buf.String()
	case "reserved":
		return tRESERVED, buf.String()
	// BEGIN proto2
	case "optional":
		return tOPTIONAL, buf.String()
	case "group":
		return tGROUP, buf.String()
	case "extensions":
		return tEXTENSIONS, buf.String()
	case "extend":
		return tEXTEND, buf.String()
	case "required":
		return tREQUIRED, buf.String()
		// END proto2
	}
	return tIDENT, buf.String()
}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	if '\n' == ch {
		s.line++
	}
	return ch
}

// unread places the previously read rune back on the reader.
// decrement the line count if it was a newline.
func (s *scanner) unread(ch rune) {
	_ = s.r.UnreadRune()
	if '\n' == ch {
		s.line--
	}
}

// isWhitespace returns true if the rune is a space, tab, or newline.
func isWhitespace(ch rune) bool { return ch == ' ' || ch == '\t' || ch == '\n' }

// isLetter returns true if the rune is a letter.
func isLetter(ch rune) bool { return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') }

// isDigit returns true if the rune is a digit.
func isDigit(ch rune) bool { return (ch >= '0' && ch <= '9') }

// isLiteralTerminator returns true if the rune cannot be part of a literal.
func isLiteralTerminator(ch rune) bool { return strings.ContainsRune("[]();,", ch) }

// eof represents a marker rune for the end of the reader.
var eof = rune(0)

// scanUntil returns the string up to (not including) the terminator or EOF.
func (s *scanner) scanUntil(terminator rune) string {
	var buf bytes.Buffer
	// Read every subsequent character into the buffer.
	// New line character and EOF will cause the loop to exit.
	lastCh := ' '
	for {
		if ch := s.read(); ch == eof {
			break
		} else if ch == terminator && lastCh != '\\' {
			break
		} else {
			buf.WriteRune(ch)
			lastCh = ch
		}
	}
	return buf.String()
}

// peek returns true if a rune is ahead.
func (s *scanner) peek(ch rune) bool {
	r := s.read()
	s.unread(r)
	return r == ch
}

// scanComment returns the string after // or between /* and */. COMMENT token was consumed.
func (s *scanner) scanComment() string {
	next := s.read()
	if '/' == next {
		// single line
		return s.scanUntil('\n')
	}
	if '*' != next {
		s.unread(next)
		return fmt.Sprintf("%d %s", next, string(next))
	}
	var buf bytes.Buffer
	// Read every subsequent character into the buffer.
	// */ and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if ch == '*' && s.peek('/') {
			s.read() // consume /
			break
		} else {
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}
