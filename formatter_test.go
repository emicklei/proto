package proto

import (
	"bytes"
	"strings"
	"testing"
)

func TestPrintListOfColumns(t *testing.T) {
	e0 := new(EnumField)
	e0.Name = "A"
	e0.Integer = 1
	op0 := new(Option)
	op0.IsEmbedded = true
	op0.Name = "a"
	op0.Constant = Literal{Source: "1234"}
	e0.ValueOption = op0

	e1 := new(EnumField)
	e1.Name = "ABC"
	e1.Integer = 12
	op1 := new(Option)
	op1.IsEmbedded = true
	op1.Name = "ab"
	op1.Constant = Literal{Source: "1234"}
	e1.ValueOption = op1

	list := []columnsPrintable{e0, e1}
	b := new(bytes.Buffer)
	f := NewFormatter(b, " ")
	f.printListOfColumns(list, "enum")
	formatted := `A   =  1 [a = 1234];
ABC = 12 [ab = 1234];
`
	if got, want := b.String(), formatted; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestFormatComment(t *testing.T) {
	proto := `
/*
 * Hello
 * World
 */
  `
	def, _ := NewParser(strings.NewReader(proto)).Parse()
	b := new(bytes.Buffer)
	f := NewFormatter(b, " ")
	f.Format(def)
	if got, want := strings.TrimSpace(b.String()), strings.TrimSpace(proto); got != want {
		t.Errorf("got [%s] want [%s]", got, want)
	}
}

func TestFormatInlineComment(t *testing.T) {
	proto := `
message ConnectRequest {
 string clientID       = 1; // Client name/identifier.
 string heartbeatInbox = 2; // Inbox for server initiated heartbeats.
}	
 `
	def, _ := NewParser(strings.NewReader(proto)).Parse()
	b := new(bytes.Buffer)
	f := NewFormatter(b, " ")
	f.Format(def)
	if got, want := strings.TrimSpace(b.String()), strings.TrimSpace(proto); got != want {
		t.Errorf("got [%s] want [%s]", got, want)
	}
}
