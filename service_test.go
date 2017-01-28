package proto3

import (
	"strings"
	"testing"
)

func TestService(t *testing.T) {
	proto := `service AccountService {}`
	stmt, err := NewParser(strings.NewReader(proto)).Parse()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := collect(stmt).Services()[0].Name, "AccountService"; got != want {
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
	if got, want := len(collect(stmt).Services()[0].RPCalls), 2; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
