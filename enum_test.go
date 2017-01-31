package proto3

import "testing"

func TestEnum(t *testing.T) {
	proto := `	
enum EnumAllowingAlias {
  option allow_alias = true;
  UNKNOWN = 0;
  STARTED = 1;
  RUNNING = 2 [(custom_option) = "hello world"];
}`
	p := newParserOn(proto)
	pr, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	enums := collect(pr).Enums()
	if got, want := len(enums), 1; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := len(enums[0].Elements), 4; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	ef1 := enums[0].Elements[1].(*EnumField)
	if got, want := ef1.Integer, 0; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	ef3 := enums[0].Elements[3].(*EnumField)
	if got, want := ef3.Integer, 2; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := ef3.ValueOption.Name, "custom_option"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := ef3.ValueOption.Constant.Source, "hello world"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
