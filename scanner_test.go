package proto3parser

import (
	"strings"
	"testing"
)

func TestScanUntilLineEnd(t *testing.T) {
	r := strings.NewReader(`hello
world`)
	s := NewScanner(r)
	v := s.scanUntilLineEnd()
	if got, want := v, "hello"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := s.line, 1; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
