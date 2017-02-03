package proto

import "testing"

func TestGroup(t *testing.T) {
	oto := `message M {
        optional group OptionalGroup = 16 {
            optional int32 a = 17;
        }
    }`
	p := newParserOn(oto)
	p.scanIgnoreWhitespace() // consume first token
	m := new(Message)
	err := m.parse(p)
	if err != nil {
		t.Error(err)
	}
	if got, want := len(m.Elements), 1; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	g := m.Elements[0].(*Group)
	if got, want := len(g.Elements), 1; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	f := g.Elements[0].(*NormalField)
	if got, want := f.Name, "a"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Optional, true; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
