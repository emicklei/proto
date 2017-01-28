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

func TestMessageWithFieldsAndComments(t *testing.T) {
	proto := `
		message AccountOut {
		// identifier
		string id   = 1;
		// size
		int64 size = 2;
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
	if got, want := len(m.Elements), 4; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestOneOf(t *testing.T) {
	proto := `
	message Sample {
		oneof foo {
			string name = 4;
			SubMessage sub_message = 9;
		}	
	}
`
	p := NewParser(strings.NewReader(proto))
	p.scanIgnoreWhitespace() // consume first token
	m := new(Message)
	err := m.parse(p)
	if err != nil {
		t.Fatal(err)
	}
}
