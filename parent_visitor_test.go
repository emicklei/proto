package proto

import "testing"

func TestShowParents(t *testing.T) {
	proto := `
	message   Out   {
	// identifier
	string   id  = 1;
	// size
	int64   size = 2;

	oneof foo {
		string     name        = 4;
		SubMessage sub_message = 9;
	}
	message  Inner {   // Level 2
		   int64  ival = 1;
	  }
	map<string, testdata.SubDefaults> proto2_value  =  13;
	option  (my_option).a  =  true;
}`
	p := newParserOn(proto)
	p.next() // consume first token
	m := new(Message)
	err := m.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	collector := newTree()
	m.Accept(collector)
	for _, each := range m.Elements {
		parents := collector.parents[each]
		if got, want := parents[0], m; got != want {
			t.Errorf("got [%v] want [%v]", got, want)
		}
	}
}

type tree struct {
	*ParentAwareVisitor
	parents map[Visitee][]Visitee
}

func newTree() *tree {
	collector := new(tree)
	collector.ParentAwareVisitor = NewParentAwareVisitor(collector)
	collector.parents = map[Visitee][]Visitee{}
	return collector
}

func (v *tree) VisitMessage(e *Message) {
	v.parents[e] = v.ParentAwareVisitor.Parents
	// super
	v.ParentAwareVisitor.VisitMessage(e)
}
func (v *tree) VisitService(e *Service) {
	v.parents[e] = v.ParentAwareVisitor.Parents
	// super
	v.ParentAwareVisitor.VisitService(e)
}
func (v *tree) VisitSyntax(e *Syntax) {
	v.parents[e] = v.ParentAwareVisitor.Parents
	// super
	v.ParentAwareVisitor.VisitSyntax(e)
}
func (v *tree) VisitPackage(e *Package) {
	v.parents[e] = v.ParentAwareVisitor.Parents
	// super
	v.ParentAwareVisitor.VisitPackage(e)
}
func (v *tree) VisitOption(e *Option) {
	v.parents[e] = v.ParentAwareVisitor.Parents
	// super
	v.ParentAwareVisitor.VisitOption(e)
}
func (v *tree) VisitImport(e *Import) {
	v.parents[e] = v.ParentAwareVisitor.Parents
	// super
	v.ParentAwareVisitor.VisitImport(e)
}
func (v *tree) VisitNormalField(e *NormalField) {
	v.parents[e] = v.ParentAwareVisitor.Parents
	// super
	v.ParentAwareVisitor.VisitNormalField(e)
}
func (v *tree) VisitEnumField(e *EnumField) {
	v.parents[e] = v.ParentAwareVisitor.Parents
	// super
	v.ParentAwareVisitor.VisitEnumField(e)
}
func (v *tree) VisitEnum(e *Enum) {
	v.parents[e] = v.ParentAwareVisitor.Parents
	// super
	v.ParentAwareVisitor.VisitEnum(e)
}
func (v *tree) VisitComment(e *Comment) {
	v.parents[e] = v.ParentAwareVisitor.Parents
	// super
	v.ParentAwareVisitor.VisitComment(e)
}
func (v *tree) VisitOneof(e *Oneof) {
	v.parents[e] = v.ParentAwareVisitor.Parents
	// super
	v.ParentAwareVisitor.VisitOneof(e)
}
func (v *tree) VisitOneofField(e *OneOfField) {
	v.parents[e] = v.ParentAwareVisitor.Parents
	// super
	v.ParentAwareVisitor.VisitOneofField(e)
}
func (v *tree) VisitReserved(e *Reserved) {
	v.parents[e] = v.ParentAwareVisitor.Parents
	// super
	v.ParentAwareVisitor.VisitReserved(e)
}
func (v *tree) VisitRPC(e *RPC) {
	v.parents[e] = v.ParentAwareVisitor.Parents
	// super
	v.ParentAwareVisitor.VisitRPC(e)
}
func (v *tree) VisitMapField(e *MapField) {
	v.parents[e] = v.ParentAwareVisitor.Parents
	// super
	v.ParentAwareVisitor.VisitMapField(e)
}
func (v *tree) VisitGroup(e *Group) {
	v.parents[e] = v.ParentAwareVisitor.Parents
	// super
	v.ParentAwareVisitor.VisitGroup(e)
}
func (v *tree) VisitExtensions(e *Extensions) {
	v.parents[e] = v.ParentAwareVisitor.Parents
	// super
	v.ParentAwareVisitor.VisitExtensions(e)
}
