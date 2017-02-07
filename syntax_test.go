package proto

import "testing"

func TestSyntax(t *testing.T) {
	proto := `syntax = "proto";`
	p := newParserOn(proto)
	p.scanIgnoreWhitespace() // consume first token
	s := new(Syntax)
	err := s.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := s.Value, "proto"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestCommentAroundSyntax(t *testing.T) {
	proto := `
	// comment1
	// comment2
	syntax = 'proto'; // comment3
	// comment4
`
	p := newParserOn(proto)
	r, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	comments := collect(r).Comments()
	if got, want := len(comments), 3; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
}
