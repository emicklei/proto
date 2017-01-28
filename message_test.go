package proto3

import (
	"strings"
	"testing"
)

func TestMessage(t *testing.T) {
	proto := `message AccountOut {}`
	p := NewParser(strings.NewReader(proto))
	p.scanIgnoreWhitespace() // consume first token
	m := new(Message)
	err := m.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := m.Name, "AccountOut"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestMessageWithFields(t *testing.T) {
	proto := `message AccountOut {
		string id   = 1;
		int64  size = 2;
	}`
	p := NewParser(strings.NewReader(proto))
	p.scanIgnoreWhitespace() // consume first token
	m := new(Message)
	err := m.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := m.Name, "AccountOut"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := len(m.Fields), 2; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
