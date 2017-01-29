package proto3

import "testing"

func TestMessage(t *testing.T) {
	proto := `
		message Out {
		// identifier
		string id   = 1;
		// size
		int64 size = 2;
		
		oneof foo {
			string name = 4;
			SubMessage sub_message = 9;
		}
		message Inner {   // Level 2
   			int64 ival = 1;
  		}
		map<string, testdata.SubDefaults> proto2_value = 13;
		option (my_option).a = true;
	}`
	p := newParserOn(proto)
	p.scanIgnoreWhitespace() // consume first token
	m := new(Message)
	err := m.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := m.Name, "Out"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := len(m.Elements), 9; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
