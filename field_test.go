package proto3

import "testing"

func TestField(t *testing.T) {
	proto := `repeated foo.bar lots = 1 [option1=a, option2=b, option3="happy"];`
	p := newParserOn(proto)
	f := new(Field)
	err := f.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := f.Repeated, true; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Type, "foo.bar"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Name, "lots"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := len(f.Options), 3; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[0].Name, "option1"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[0].Identifier, "a"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[1].Name, "option2"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[1].Identifier, "b"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[2].String, "happy"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestFieldSyntaxErrors(t *testing.T) {
	for i, each := range []string{
		`repeatet foo.bar lots = 1;`,
		`string lots == 1;`,
	} {
		if new(Field).parse(newParserOn(each)) == nil {
			t.Errorf("%d: uncaught syntax error", i)
		}
	}
}
