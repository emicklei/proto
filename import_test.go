package proto3parser

import (
	"strings"
	"testing"
)

func TestParseImport(t *testing.T) {
	proto := `import public "other.proto";`
	p := NewParser(strings.NewReader(proto))
	p.scanIgnoreWhitespace() // consume first token
	i := new(Import)
	err := i.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := i.Filename, "other.proto"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := i.Kind, "public"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
