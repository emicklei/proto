package proto3

// Visitor is for dispatching Proto elements.
type Visitor interface {
	VisitMessage(m *Message)
	VisitService(v *Service)
	VisitSyntax(s *Syntax)
	VisitPackage(p *Package)
	VisitOption(o *Option)
	VisitImport(i *Import)
	VisitField(i *Field)
	VisitEnumField(i *EnumField)
	VisitEnum(e *Enum)
	VisitComment(e *Comment)
	VisitOneof(o *Oneof)
	VisitOneofField(o *OneOfField)
	VisitReserved(r *Reserved)
	VisitRPcall(r *RPcall)
}

// Visitee is implemented by all Proto elements.
type Visitee interface {
	Accept(v Visitor)
}
