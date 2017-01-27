package proto3parser

import (
	"strings"
	"testing"
)

func TestSyntax(t *testing.T) {
	proto := `syntax = "proto3";`
	p := NewParser(strings.NewReader(proto))
	p.scanIgnoreWhitespace() // consume first token
	s := new(Syntax)
	err := s.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := s.Value, "proto3"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
