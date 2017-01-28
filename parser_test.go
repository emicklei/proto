package proto3

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
	p := NewParser(strings.NewReader(proto))
	pr, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(pr.Comments), 2; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
