package proto

import (
	"strings"
	"testing"
)

func TestParseComment(t *testing.T) {
	proto := `
    // single
    /* 
    multi* 
    */`
	p := newParserOn(proto)
	pr, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(collect(pr).Comments()), 2; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func newParserOn(def string) *Parser {
	p := NewParser(strings.NewReader(def))
	p.debug = true
	return p
}

// TEMPORARY tests
func TestScanIgnoreWhitespace_Digits(t *testing.T) {
	p := newParserOn("1234")
	_, lit := p.scanIgnoreWhitespace()
	if got, want := lit, "1"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestScanIgnoreWhitespace_Minus(t *testing.T) {
	p := newParserOn("-1234")
	_, lit := p.scanIgnoreWhitespace()
	if got, want := lit, "-"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
