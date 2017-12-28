package proto

type ParentAwareVisitor struct {
	Parent Visitee
}

func (v *ParentAwareVisitor) VisitMessage(m *Message) {
	v.Parent = m
	for _, each := range m.Elements {
		each.Accept(v)
	}
}
func (v *ParentAwareVisitor) VisitService(s *Service) {
	v.Parent = s
	for _, each := range s.Elements {
		each.Accept(v)
	}
}
func (v *ParentAwareVisitor) VisitSyntax(s *Syntax)   {}
func (v *ParentAwareVisitor) VisitPackage(p *Package) {}
func (v *ParentAwareVisitor) VisitOption(o *Option)   {}
func (v *ParentAwareVisitor) VisitImport(i *Import)   {}
func (v *ParentAwareVisitor) VisitNormalField(i *NormalField) {
	v.Parent = i
	for _, each := range i.Options {
		each.Accept(v)
	}
}
func (v *ParentAwareVisitor) VisitEnumField(i *EnumField) {
	v.Parent = i
	i.ValueOption.Accept(v)
}
func (v *ParentAwareVisitor) VisitEnum(e *Enum) {
	v.Parent = e
	for _, each := range e.Elements {
		each.Accept(v)
	}
}
func (v *ParentAwareVisitor) VisitComment(e *Comment) {}
func (v *ParentAwareVisitor) VisitOneof(o *Oneof) {
	v.Parent = o
	for _, each := range o.Elements {
		each.Accept(v)
	}
}
func (v *ParentAwareVisitor) VisitOneofField(o *OneOfField) {
	v.Parent = o
	for _, each := range o.Options {
		each.Accept(v)
	}
}
func (v *ParentAwareVisitor) VisitReserved(rs *Reserved) {}
func (v *ParentAwareVisitor) VisitRPC(rpc *RPC)          {}
func (v *ParentAwareVisitor) VisitMapField(f *MapField) {
	v.Parent = f
	for _, each := range f.Options {
		each.Accept(v)
	}
}
func (v *ParentAwareVisitor) VisitGroup(g *Group) {
	v.Parent = g
	for _, each := range g.Elements {
		each.Accept(v)
	}
}
func (v *ParentAwareVisitor) VisitExtensions(e *Extensions) {}
