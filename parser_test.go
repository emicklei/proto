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

func TestScanIdent(t *testing.T) {
	p := NewParser(strings.NewReader(" message "))
	tok, lit := p.scanIdent()
	if got, want := tok, tIDENT; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := lit, "message"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
