package proto3parser

type Visitor interface {
	VisitMessage(m *Message)
	VisitService(v *Service)
}

type Visitee interface {
	Accept(v Visitor)
}
