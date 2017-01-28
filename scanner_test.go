package proto3

import (
	"strings"
	"testing"
)

func TestScanUntilLineEnd(t *testing.T) {
	r := strings.NewReader(`hello
world`)
	s := newScanner(r)
	v := s.scanUntil('\n')
	if got, want := v, "hello"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := s.line, 1; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestScanMultilineComment(t *testing.T) {
	r := strings.NewReader(`
	/*
	𝓢𝓱𝓸𝓾𝓵𝓭 𝔽𝕠𝕣𝕞𝕒𝕥𝕥𝕚𝕟𝕘 𝘐𝘯 𝓣𝓲𝓽𝓵𝓮𝓼 𝕭𝖊 *🅿🅴🆁🅼🅸🆃🆃🅴🅳* ?
	*/
`)
	s := newScanner(r)
	s.scanUntil('/') // consume COMMENT token
	if got, want := s.scanComment(), `
	𝓢𝓱𝓸𝓾𝓵𝓭 𝔽𝕠𝕣𝕞𝕒𝕥𝕥𝕚𝕟𝕘 𝘐𝘯 𝓣𝓲𝓽𝓵𝓮𝓼 𝕭𝖊 *🅿🅴🆁🅼🅸🆃🆃🅴🅳* ?
	`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestScanSingleLineComment(t *testing.T) {
	r := strings.NewReader(`
	// dreadful //
`)
	s := newScanner(r)
	s.scanUntil('/') // consume COMMENT token
	if got, want := s.scanComment(), ` dreadful //`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
