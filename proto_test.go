package proto

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestParseFormattedProto2UnitTest(t *testing.T) {
	parseFormattedParsed(t, "unittest_proto2.proto")
}

func TestParseFormattedProto3UnitTest(t *testing.T) {
	parseFormattedParsed(t, "unittest_proto3.proto")
}

func TestParseFormattedProto3ArenaUnitTest(t *testing.T) {
	parseFormattedParsed(t, "unittest_proto3_arena.proto")
}

func parseFormattedParsed(t *testing.T, filename string) {
	// open it
	f, err := os.Open(filepath.Join("cmd", "protofmt", filename))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	// parse it
	p := NewParser(f)
	def, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	// format it
	out := new(bytes.Buffer)
	fmt := NewFormatter(out, "  ")
	fmt.Format(def)
	// parse the formatted content
	fp := NewParser(bytes.NewReader(out.Bytes()))
	_, err = fp.Parse()
	if err != nil {
		t.Fatal(err)
	}
}
