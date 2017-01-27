package proto3parser

import (
	"strconv"
	"strings"
	"testing"
)

func TestSyntax(t *testing.T) {
	proto := `syntax = "proto3";`
	p := NewParser(strings.NewReader(proto))
	p.scanIgnoreWhitespace() // consume first token
	syntax, err := ParseSyntax(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := syntax, "proto3"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestService(t *testing.T) {
	proto := `service AccountService {}`
	stmt, err := NewParser(strings.NewReader(proto)).Parse()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := stmt.Services[0].Name, "AccountService"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestServiceWithRPCs(t *testing.T) {
	proto := `service AccountService {
		rpc CreateAccount (CreateAccount) returns    (ServiceFault) {}
		rpc GetAccount 	  (Int64)           returns (Account) {}	
	}`
	stmt, err := NewParser(strings.NewReader(proto)).Parse()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(stmt.Services[0].RPCalls), 2; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestOption(t *testing.T) {
	proto := `option java_package = "com.example.foo";`
	p := NewParser(strings.NewReader(proto))
	p.scanIgnoreWhitespace() // consume first token
	o, err := ParseOption(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := o.Name, "java_package"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := o.Constant, "com.example.foo"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestOptionFull(t *testing.T) {
	proto := `option (full.java_package) = "com.example.foo";`
	p := NewParser(strings.NewReader(proto))
	p.scanIgnoreWhitespace() // consume first token
	o, err := ParseOption(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := o.Name, "full.java_package"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := o.Constant, "com.example.foo"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestCommentAroundSyntax(t *testing.T) {
	proto := `
	// comment1
	// comment2
	syntax = "proto3"; // comment3
	// comment4
`
	p := NewParser(strings.NewReader(proto))
	r, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(r.Comments), 4; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	for i := 1; i <= 4; i++ {
		if got, want := r.Comments[i-1].Message, " comment"+strconv.Itoa(i); got != want {
			t.Errorf("got [%v] want [%v]", got, want)
		}
		if got, want := r.Comments[i-1].Line, i+1; got != want {
			t.Errorf("got [%v] want [%v]", got, want)
		}
	}
}
