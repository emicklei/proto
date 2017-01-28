package proto3

import (
	"strconv"
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

func TestCommentAroundSyntax(t *testing.T) {
	proto := `
	// comment1
	// comment2
	syntax = "proto3"; // comment3
	// comment4
`
	p := NewParser(strings.NewReader(proto))
	r, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(r.Comments), 4; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	for i := 1; i <= 4; i++ {
		if got, want := r.Comments[i-1].Message, " comment"+strconv.Itoa(i); got != want {
			t.Errorf("got [%v] want [%v]", got, want)
		}
		if got, want := r.Comments[i-1].Line, i+1; got != want {
			t.Errorf("got [%v] want [%v]", got, want)
		}
	}
}
