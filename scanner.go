package proto3

// partial code from https://raw.githubusercontent.com/benbjohnson/sql-parser/master/scanner.go

import (
	"bufio"
	"bytes"
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
	return &scanner{r: bufio.NewReader(r)}
}

// scan returns the next token and literal value.
func (s *scanner) scan() (tok token, lit string) {
	// Read the next rune.
	ch := s.read()

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as an ident or reserved word.
	// If we see a digit then consume as a number.
	// If we see a double slash then consume all as a comment line
	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if isLetter(ch) {
		s.unread()
		return s.scanIdent()
	} else if ch == '/' {
		if ch = s.read(); ch == '/' {
			return COMMENT, s.scanUntil('\n')
		}
		s.unread()
		s.unread()
		return s.scanIdent()
	}

	// Otherwise read the individual character.
	switch ch {
	case eof:
		return EOF, ""
	case ';':
		return SEMICOLON, string(ch)
	case '=':
		return EQUALS, string(ch)
	case '"':
		return QUOTE, string(ch)
	case '(':
		return LEFTPAREN, string(ch)
	case ')':
		return RIGHTPAREN, string(ch)
	case '{':
		return LEFTCURLY, string(ch)
	case '}':
		return RIGHTCURLY, string(ch)
	case '[':
		return LEFTSQUARE, string(ch)
	case ']':
		return RIGHTSQUARE, string(ch)
	}
	return ILLEGAL, string(ch)
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
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return WS, buf.String()
}

func (s *scanner) scanIntegerString() string {
	s.scanWhitespace()
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent digit character into the buffer.
	// Non-digiti characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isDigit(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}
	return buf.String()
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
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// If the string matches a keyword then return that keyword.
	switch strings.ToUpper(buf.String()) {
	case "SYNTAX":
		return SYNTAX, buf.String()
	case "SERVICE":
		return SERVICE, buf.String()
	case "MESSAGE":
		return MESSAGE, buf.String()
	case "RPC":
		return RPC, buf.String()
	case "RETURNS":
		return RETURNS, buf.String()
	case "IMPORT":
		return IMPORT, buf.String()
	case "PACKAGE":
		return PACKAGE, buf.String()
	case "REPEATED":
		return REPEATED, buf.String()
	case "OPTION":
		return OPTION, buf.String()
	case "ENUM":
		return ENUM, buf.String()
	case "TRUE":
		return TRUE, buf.String()
	case "FALSE":
		return FALSE, buf.String()
	}

	// Otherwise return as a regular identifier.
	return IDENT, buf.String()
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
func (s *scanner) unread() { _ = s.r.UnreadRune() }

// isWhitespace returns true if the rune is a space, tab, or newline.
func isWhitespace(ch rune) bool { return ch == ' ' || ch == '\t' || ch == '\n' }

// isLetter returns true if the rune is a letter.
func isLetter(ch rune) bool { return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') }

// isDigit returns true if the rune is a digit.
func isDigit(ch rune) bool { return (ch >= '0' && ch <= '9') }

// eof represents a marker rune for the end of the reader.
var eof = rune(0)

// scanUntil returns the string up to (not including) the terminator or EOF.
func (s *scanner) scanUntil(terminator rune) string {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())
	// Read every subsequent character into the buffer.
	// New line character and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if ch == terminator {
			break
		} else {
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}
