// Copyright (c) 2017 Ernest Micklei
// 
// MIT License
// 
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
// 
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
	formatted := `
A   =  1 [a  = 1234];
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

func logformatted(t *testing.T, v Visitee) {
	b := new(bytes.Buffer)
	f := NewFormatter(b, " ")
	v.Accept(f)
	t.Log(b.String())
}
