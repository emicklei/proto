package proto3

import "testing"

func TestOption(t *testing.T) {
	proto := `option (full.java_package) = "com.example.foo";`
	p := newParserOn(proto)
	pr, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(pr.Elements), 1; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	o := pr.Elements[0].(*Option)
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
