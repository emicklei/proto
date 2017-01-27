package proto3parser

import (
	"strings"
	"testing"
)

func TestRepeatedField(t *testing.T) {
	proto := `repeated string lots = 1;`
	p := NewParser(strings.NewReader(proto))
	f := new(Field)
	err := ParseField(f, p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := f.Repeated, true; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
