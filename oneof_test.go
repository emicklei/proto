package proto

import "testing"

func TestOneof(t *testing.T) {
	proto := `oneof foo {
    string name = 4;
    SubMessage sub_message = 9 [options=none];
}`
	p := newParserOn(proto)
	p.scanIgnoreWhitespace() // consume first token
	o := new(Oneof)
	err := o.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := o.Name, "foo"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := len(o.Elements), 2; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	second := o.Elements[1].(*OneOfField)
	if got, want := second.Name, "sub_message"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := second.Type, "SubMessage"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := second.Sequence, 9; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
