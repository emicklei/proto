package proto3

import "testing"

func TestOption(t *testing.T) {
	proto := `option (full.java_package) = "com.example.foo";`
	p := newParserOn(proto)
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
	if got, want := o.PartOfFieldOrEnum, false; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
