package proto3

import "testing"

func TestService(t *testing.T) {
	proto := `service AccountService {
		rpc CreateAccount (CreateAccount) returns    (ServiceFault) {}
		rpc GetAccount 	  (Int64)           returns (Account) {}	
	}`
	stmt, err := newParserOn(proto).Parse()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(collect(stmt).Services()[0].RPCalls), 2; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
