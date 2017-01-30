package proto3

import "testing"

func TestService(t *testing.T) {
	t.Skip()
	proto := `service AccountService {
		// comment
		rpc CreateAccount (CreateAccount) returns (ServiceFault);
		// comment
		rpc GetAccounts   (stream Int64)  returns (Account);	
	}`
	pr, err := newParserOn(proto).Parse()
	if err != nil {
		t.Fatal(err)
	}
	srv := collect(pr).Services()[0]
	if got, want := len(srv.Elements), 4; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	rpc1 := srv.Elements[1].(*RPcall)
	if got, want := rpc1.Name, "CreateAccount"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
