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
	"fmt"
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

func formatted(t *testing.T, v Visitee) string {
	b := new(bytes.Buffer)
	f := NewFormatter(b, "  ") // 2 spaces
	v.Accept(f)
	return b.String()
}

func TestExtendMessage(t *testing.T) {
	proto := `extend google.protobuf.MessageOptions {  optional string my_option = 51234; }`
	p := newParserOn(proto)
	p.scanIgnoreWhitespace() // consume first token
	m := new(Message)
	m.IsExtend = true
	err := m.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := formatted(t, m), `
extend google.protobuf.MessageOptions {
  optional string my_option = 51234;
}
`; got != want {
		fmt.Println(diff(t, got, want))
		t.Fail()
	}
}

func TestAggregatedOptionSyntax(t *testing.T) {
	proto := `rpc Find (  Finder  ) returns ( stream Result ) {
            option (google.api.http) = {
                post: "/v1/finders/1"
                body: "*"
            };
       }`
	p := newParserOn(proto)
	r := new(RPC)
	p.scanIgnoreWhitespace() // consumer rpc
	err := r.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := formatted(t, r), `
rpc Find (Finder) returns (stream Result) {
  option (google.api.http) = {
    post: "/v1/finders/1"
    body: "*"
  };
}
`; got != want {
		fmt.Println(diff(t, got, want))
		fmt.Println("---")
		fmt.Println(got)
		fmt.Println("---")
		fmt.Println(want)
		t.Fail()
	}
}

func diff(t *testing.T, left, right string) string {
	b := new(bytes.Buffer)
	w := func(char rune) {
		if '\n' == char {
			b.WriteString("(n)")
		} else if '\t' == char {
			b.WriteString("(t)")
		} else if ' ' == char {
			b.WriteString("( )")
		} else {
			b.WriteRune(char)
		}
	}
	for _, char := range left {
		w(char)
	}
	if len(left) == 0 {
		b.WriteString("(empty)")
	}
	b.WriteString("\n")
	for _, char := range right {
		w(char)
	}
	return b.String()
}
