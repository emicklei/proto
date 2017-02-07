package proto

import "testing"

func TestExtensions(t *testing.T) {
	proto := `message M { 
		extensions 4, 20 to max; // max
	}`
	p := newParserOn(proto)
	p.scanIgnoreWhitespace() // consume extensions
	m := new(Message)
	err := m.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if len(m.Elements) == 0 {
		t.Fatal("extensions expected")
	}
	f := m.Elements[0].(*Extensions)
	if got, want := f.Ranges, " 4, 20 to max"; got != want {
		t.Errorf("got [%s] want [%s]", got, want)
	}
	if f.Comment == nil {
		t.Fatal("comment expected")
	}
	if got, want := f.Comment.Message, " max"; got != want {
		t.Errorf("got [%s] want [%s]", got, want)
	}
}
