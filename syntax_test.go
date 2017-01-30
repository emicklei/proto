package proto3

import (
	"strconv"
	"testing"
)

func TestSyntax(t *testing.T) {
	proto := `syntax = "proto3";`
	p := newParserOn(proto)
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

func TestCommentAroundSyntax(t *testing.T) {
	proto := `
	// comment1
	// comment2
	syntax = "proto3"; // comment3
	// comment4
`
	p := newParserOn(proto)
	r, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	comments := collect(r).Comments()
	if got, want := len(comments), 4; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	for i := 1; i <= 4; i++ {
		if got, want := comments[i-1].Message, " comment"+strconv.Itoa(i); got != want {
			t.Errorf("got [%v] want [%v]", got, want)
		}
	}
}
