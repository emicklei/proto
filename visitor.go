package proto3

// Visitor is for dispatching Proto elements.
type Visitor interface {
	VisitMessage(m *Message)
	VisitService(v *Service)
	VisitSyntax(s *Syntax)
	VisitPackage(p *Package)
	VisitOption(o *Option)
}

// visitee is implemented by all Proto elements.
type visitee interface {
	accept(v Visitor)
}
