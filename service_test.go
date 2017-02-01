package proto

import "testing"

func TestService(t *testing.T) {
	proto := `service AccountService {
		// comment
		rpc CreateAccount (CreateAccount) returns (ServiceFault);
		rpc GetAccounts   (stream Int64)  returns (Account);	
	}`
	pr, err := newParserOn(proto).Parse()
	if err != nil {
		t.Fatal(err)
	}
	srv := collect(pr).Services()[0]
	if got, want := len(srv.Elements), 3; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	rpc1 := srv.Elements[1].(*RPC)
	if got, want := rpc1.Name, "CreateAccount"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	rpc2 := srv.Elements[2].(*RPC)
	if got, want := rpc2.Name, "GetAccounts"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
