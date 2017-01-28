package proto3

import (
	"strings"
	"testing"
)

func TestOption(t *testing.T) {
	proto := `option java_package = "com.example.foo";`
	p := NewParser(strings.NewReader(proto))
	p.scanIgnoreWhitespace() // consume first token
	o := new(Option)
	err := o.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := o.Name, "java_package"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := o.String, "com.example.foo"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestOptionFull(t *testing.T) {
	proto := `option (full.java_package) = "com.example.foo";`
	p := NewParser(strings.NewReader(proto))
	p.scanIgnoreWhitespace() // consume first token
	o := new(Option)
	err := o.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := o.Name, "full.java_package"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := o.String, "com.example.foo"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
