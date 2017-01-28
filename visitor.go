package proto3

// Visitor is for dispatching Proto elements.
type Visitor interface {
	VisitMessage(m *Message)
	VisitService(v *Service)
	VisitSyntax(s *Syntax)
	VisitPackage(p *Package)
	VisitOption(o *Option)
}

// Visitee is implemented by all Proto elements.
type Visitee interface {
	Accept(v Visitor)
}
