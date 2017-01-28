package proto3

// Package specifies the namespace for all proto elements.
type Package struct {
	Name string
}

// accept dispatches the call to the visitor.
func (p *Package) accept(v Visitor) {
	v.VisitPackage(p)
}

func (p *Package) parse(pr *Parser) error {
	return nil
}
