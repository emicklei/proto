package proto3

import (
	"strings"
	"testing"
)

func TestEnum(t *testing.T) {
	proto := `enum EnumAllowingAlias {}`
	p := NewParser(strings.NewReader(proto))
	_, tok := p.scanIgnoreWhitespace() // consume first token
	if got, want := tok, "enum"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	e := new(Enum)
	err := e.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := e.Name, "EnumAllowingAlias"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestEnumWithBody(t *testing.T) {
	proto := `
enum EnumAllowingAlias {
  option allow_alias = true;
  UNKNOWN = 0;
  STARTED = 1;
  RUNNING = 2 [(custom_option) = "hello world"];
}`
	p := NewParser(strings.NewReader(proto))
	pr, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(pr.Enums), 1; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := len(pr.Enums[0].EnumFields), 3; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := pr.Enums[0].EnumFields[0].Integer, 0; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
