package proto

// Extensions declare that a range of field numbers in a message are available for third-party extensions.
// proto2 only
type Extensions struct {
	Ranges string
}

// Accept dispatches the call to the visitor.
func (e *Extensions) Accept(v Visitor) {
	v.VisitExtensions(e)
}

// parse expects ident { messageBody
func (e *Extensions) parse(p *Parser) error {
	// TODO proper range parsing
	e.Ranges = p.s.scanUntil(';')
	return nil
}
