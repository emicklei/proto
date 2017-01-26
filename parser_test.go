package proto3parser

import (
	"fmt"
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
	proto := `service AccountService {
		rpc CreateAccount (CreateAccount) returns    (ServiceFault) {}
		rpc GetAccount 	  (Int64)           returns (Account) {}	
	}`
	stmt, err := NewParser(strings.NewReader(proto)).Parse()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(stmt)
}

func TestMessage(t *testing.T) {
	proto := `message AccountOut {
  optional ServiceFault 	fault 		= 1;
  required int64 		account_id 	= 2;
}`
	stmt, err := NewParser(strings.NewReader(proto)).Parse()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(stmt)
}
