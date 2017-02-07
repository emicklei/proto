package proto

import "testing"

func TestOption(t *testing.T) {
	for i, each := range []struct {
		proto     string
		name      string
		strLit    string
		nonStrLit string
	}{{
		`option (full).java_package = "com.example.foo";`,
		"(full).java_package",
		"com.example.foo",
		"",
	}, {
		`option Bool = true;`,
		"Bool",
		"",
		"true",
	}, {
		`option Float = -3.14E1;`,
		"Float",
		"",
		"-3.14E1",
	}} {
		p := newParserOn(each.proto)
		pr, err := p.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if got, want := len(pr.Elements), 1; got != want {
			t.Fatalf("[%d] got [%v] want [%v]", i, got, want)
		}
		o := pr.Elements[0].(*Option)
		if got, want := o.Name, each.name; got != want {
			t.Errorf("[%d] got [%v] want [%v]", i, got, want)
		}
		if len(each.strLit) > 0 {
			if got, want := o.Constant.Source, each.strLit; got != want {
				t.Errorf("[%d] got [%v] want [%v]", i, got, want)
			}
		}
		if len(each.nonStrLit) > 0 {
			if got, want := o.Constant.Source, each.nonStrLit; got != want {
				t.Errorf("[%d] got [%v] want [%v]", i, got, want)
			}
		}
		if got, want := o.IsEmbedded, false; got != want {
			t.Errorf("[%d] got [%v] want [%v]", i, got, want)
		}
	}
}
