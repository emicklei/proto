package proto3parser

import (
	"strings"
	"testing"
)

func TestSyntax(t *testing.T) {
	proto := `syntax = "proto3";`
	stmt, err := NewParser(strings.NewReader(proto)).Parse()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := stmt.Syntax, "proto3"; got != want {
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

func TestField(t *testing.T) {
	proto := `repeated inner inner_message = 2;`
	p := NewParser(strings.NewReader(proto))
	f, err := ParseField(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := f.Repeated, true; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestMessage(t *testing.T) {
	proto := `message AccountOut {}`
	stmt, err := NewParser(strings.NewReader(proto)).Parse()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := stmt.Messages[0].Name, "AccountOut"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
