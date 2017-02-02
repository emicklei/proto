package proto

import (
	"bytes"
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
	f.printListOfColumns(list)
	formatted := `A   =  1 [a =1234];
ABC = 12 [ab=1234];
`
	if got, want := b.String(), formatted; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
