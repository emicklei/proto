package proto

import (
	"os"
	"testing"
)

type counter struct {
	counts map[string]int
}

func (c counter) handleService(s *Service) {
	c.counts["service"] = c.counts["service"] + 1
}

func (c counter) handleRPC(r *RPC) {
	c.counts["rpc"] = c.counts["rpc"] + 1
}

func (c counter) handleImport(r *Import) {
	c.counts["import"] = c.counts["import"] + 1
}

func (c counter) handleNormalField(r *NormalField) {
	c.counts["normal field"] = c.counts["import"] + 1
}

func TestWalkGoogleApisDLP(t *testing.T) {
	if len(os.Getenv("PB")) == 0 {
		t.Skip("PB test not run")
	}
	proto := fetchAndParse(t, "https://raw.githubusercontent.com/gogo/protobuf/master/test/theproto3/theproto3.proto")
	count := counter{counts: map[string]int{}}
	Walk(proto,
		WithPackage(func(p *Package) {
			t.Log("package:", p.Name)
		}),
		WithService(count.handleService),
		WithRPC(count.handleRPC),
		WithImport(count.handleImport),
		WithNormalField(count.handleNormalField),
	)
	t.Logf("%#v", count)
}
