package proto

import "fmt"

// Package specifies the namespace for all proto elements.
type Package struct {
	Name string
}

// Accept dispatches the call to the visitor.
func (p *Package) Accept(v Visitor) {
	v.VisitPackage(p)
}

func (p *Package) parse(pr *Parser) error {
	tok, lit := pr.scanIgnoreWhitespace()
	if tIDENT != tok {
		return fmt.Errorf("found %q, expected identifier", lit)
	}
	p.Name = lit
	return nil
}
