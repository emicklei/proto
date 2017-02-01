package proto

// Visitor is for dispatching Proto elements.
type Visitor interface {
	VisitMessage(m *Message)
	VisitService(v *Service)
	VisitSyntax(s *Syntax)
	VisitPackage(p *Package)
	VisitOption(o *Option)
	VisitImport(i *Import)
	VisitNormalField(i *NormalField)
	VisitEnumField(i *EnumField)
	VisitEnum(e *Enum)
	VisitComment(e *Comment)
	VisitOneof(o *Oneof)
	VisitOneofField(o *OneOfField)
	VisitReserved(r *Reserved)
	VisitRPC(r *RPC)
	VisitMapField(f *MapField)
}

// Visitee is implemented by all Proto elements.
type Visitee interface {
	Accept(v Visitor)
}

// reflector is a Visitor that can tell the short type name of a Visitee.
type reflector struct {
	name string
}

// sole instance of reflector
var namer = new(reflector)

func (r *reflector) VisitMessage(m *Message)         { r.name = "Message" }
func (r *reflector) VisitService(v *Service)         { r.name = "Service" }
func (r *reflector) VisitSyntax(s *Syntax)           { r.name = "Syntax" }
func (r *reflector) VisitPackage(p *Package)         { r.name = "Package" }
func (r *reflector) VisitOption(o *Option)           { r.name = "Option" }
func (r *reflector) VisitImport(i *Import)           { r.name = "Import" }
func (r *reflector) VisitNormalField(i *NormalField) { r.name = "NormalField" }
func (r *reflector) VisitEnumField(i *EnumField)     { r.name = "EnumField" }
func (r *reflector) VisitEnum(e *Enum)               { r.name = "Enum" }
func (r *reflector) VisitComment(e *Comment)         { r.name = "Comment" }
func (r *reflector) VisitOneof(o *Oneof)             { r.name = "Oneof" }
func (r *reflector) VisitOneofField(o *OneOfField)   { r.name = "OneOfField" }
func (r *reflector) VisitReserved(rs *Reserved)      { r.name = "Reserved" }
func (r *reflector) VisitRPC(rpc *RPC)               { r.name = "RPC" }
func (r *reflector) VisitMapField(f *MapField)       { r.name = "MapField" }

// nameOfVisitee returns the short type name of a Visitee.
func nameOfVisitee(e Visitee) string {
	e.Accept(namer)
	return namer.name
}
