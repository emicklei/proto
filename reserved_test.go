package proto3

import "testing"

func TestReservedRanges(t *testing.T) {
	r := new(Reserved)
	p := newParserOn(`reserved 2, 15, 9 to 11;`)
	tok, _ := p.scanIgnoreWhitespace()
	if tRESERVED != tok {
		t.Fail()
	}
	err := r.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := r.Ranges, "2, 15, 9 to 11"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestReservedFieldNames(t *testing.T) {
	r := new(Reserved)
	p := newParserOn(`reserved "foo", "bar";`)
	_, _ = p.scanIgnoreWhitespace()
	err := r.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(r.FieldNames), 2; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := r.FieldNames[0], "foo"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := r.FieldNames[1], "bar"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
