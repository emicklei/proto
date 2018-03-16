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
	"testing"
)

func TestOptionCases(t *testing.T) {
	for i, each := range []struct {
		proto     string
		name      string
		strLit    string
		nonStrLit string
	}{{
		`option (full).java_package = "com.example.foo";`,
		"(full).java_package",
		"com.example.foo",
		"",
	}, {
		`option Bool = true;`,
		"Bool",
		"",
		"true",
	}, {
		`option Float = -3.14E1;`,
		"Float",
		"",
		"-3.14E1",
	}, {
		`option (foo_options) = { opt1: 123 opt2: "baz" };`,
		"(foo_options)",
		"",
		"",
	}, {
		`option optimize_for = SPEED;`,
		"optimize_for",
		"",
		"SPEED",
	}, {
		"option (my.enum.service.is.like).rpc = 1;",
		"(my.enum.service.is.like).rpc",
		"",
		"1",
	}} {
		p := newParserOn(each.proto)
		pr, err := p.Parse()
		if err != nil {
			t.Fatal("testcase failed:", i, err)
		}
		if got, want := len(pr.Elements), 1; got != want {
			t.Fatalf("[%d] got [%v] want [%v]", i, got, want)
		}
		o := pr.Elements[0].(*Option)
		if got, want := o.Name, each.name; got != want {
			t.Errorf("[%d] got [%v] want [%v]", i, got, want)
		}
		if len(each.strLit) > 0 {
			if got, want := o.Constant.Source, each.strLit; got != want {
				t.Errorf("[%d] got [%v] want [%v]", i, got, want)
			}
		}
		if len(each.nonStrLit) > 0 {
			if got, want := o.Constant.Source, each.nonStrLit; got != want {
				t.Errorf("[%d] got [%v] want [%v]", i, got, want)
			}
		}
		if got, want := o.IsEmbedded, false; got != want {
			t.Errorf("[%d] got [%v] want [%v]", i, got, want)
		}
	}
}

func TestLiteralString(t *testing.T) {
	proto := `"string"`
	p := newParserOn(proto)
	l := new(Literal)
	if err := l.parse(p); err != nil {
		t.Fatal(err)
	}
	if got, want := l.IsString, true; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := l.Source, "string"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestOptionComments(t *testing.T) {
	proto := `
// comment
option Help = "me"; // inline`
	p := newParserOn(proto)
	pr, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	o := pr.Elements[0].(*Option)
	if got, want := o.IsEmbedded, false; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := o.Comment != nil, true; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := o.Comment.Lines[0], " comment"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := o.InlineComment != nil, true; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := o.InlineComment.Lines[0], " inline"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := o.Position.Line, 3; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := o.Comment.Position.Line, 2; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := o.InlineComment.Position.Line, 3; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
}

func TestAggregateSyntax(t *testing.T) {
	proto := `
// usage:
message Bar {
  // alternative aggregate syntax (uses TextFormat):
  int32 b = 2 [(foo_options) = {
    opt1: 123,
    opt2: "baz"
  }];
}
	`
	p := newParserOn(proto)
	pr, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	o := pr.Elements[0].(*Message)
	f := o.Elements[0].(*NormalField)
	if got, want := len(f.Options), 1; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	ac := f.Options[0].AggregatedConstants
	if got, want := len(ac), 2; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := ac[0].Name, "opt1"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := ac[1].Name, "opt2"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := ac[0].Source, "123"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := ac[1].Source, "baz"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := o.Position.Line, 3; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := o.Comment.Position.String(), "<input>:2:1"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Position.String(), "<input>:5:3"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := ac[0].Position.Line, 6; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := ac[1].Position.Line, 7; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
}

func TestNonPrimitiveOptionComment(t *testing.T) {
	proto := `
// comment
option Help = { string_field: "value" }; // inline`
	p := newParserOn(proto)
	pr, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	o := pr.Elements[0].(*Option)
	if got, want := o.Comment != nil, true; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := o.Comment.Lines[0], " comment"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := o.InlineComment != nil, true; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := o.InlineComment.Lines[0], " inline"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
}

func TestFieldCustomOptions(t *testing.T) {
	proto := `foo.bar lots = 1 [foo={hello:1}, bar=2];`
	p := newParserOn(proto)
	f := newNormalField()
	err := f.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := f.Type, "foo.bar"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Name, "lots"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := len(f.Options), 2; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[0].Name, "foo"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[1].Name, "bar"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[1].Constant.Source, "2"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestFieldCustomOptionExtendedIdent(t *testing.T) {
	proto := `Type field = 1 [(validate.rules).enum.defined_only = true];`
	p := newParserOn(proto)
	f := newNormalField()
	err := f.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := f.Options[0].Name, "(validate.rules).enum.defined_only"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

// issue #50
func TestNestedAggregateConstants(t *testing.T) {
	src := `syntax = "proto3";

	package baz;

	option (foo.bar) = {
	  woot: 100
	  foo {
		hello: 200
		hello2: 300
		bar {
			hello3: 400
		}
	  }
	};`
	p := newParserOn(src)
	proto, err := p.Parse()
	if err != nil {
		t.Error(err)
	}
	option := proto.Elements[2].(*Option)
	if got, want := option.Name, "(foo.bar)"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := len(option.AggregatedConstants), 4; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := option.AggregatedConstants[0].Name, "woot"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := option.AggregatedConstants[1].Name, "foo.hello"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := option.AggregatedConstants[2].Name, "foo.hello2"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := option.AggregatedConstants[3].Name, "foo.bar.hello3"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := option.AggregatedConstants[1].Literal.SourceRepresentation(), "200"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := option.AggregatedConstants[2].Literal.SourceRepresentation(), "300"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := option.AggregatedConstants[3].Literal.SourceRepresentation(), "400"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

// Issue #59
func TestMultiLineOptionAggregateValue(t *testing.T) {
	src := `rpc ListTransferLogs(ListTransferLogsRequest)
	returns (ListTransferLogsResponse) {
		option (google.api.http) = {
		get: "/v1/{parent=projects/*/locations/*/transferConfigs/*/runs/*}/"
			"transferLogs"
		};
}`
	p := newParserOn(src)
	rpc := new(RPC)
	p.next()
	err := rpc.parse(p)
	if err != nil {
		t.Error(err)
	}
	get := rpc.Options[0].AggregatedConstants[0]
	if got, want := get.Name, "get"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := get.Literal.Source, "/v1/{parent=projects/*/locations/*/transferConfigs/*/runs/*}/transferLogs"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
