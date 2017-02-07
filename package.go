package proto

// Package specifies the namespace for all proto elements.
type Package struct {
	Name    string
	Comment *Comment
}

func (p *Package) parse(pr *Parser) error {
	tok, lit := pr.scanIgnoreWhitespace()
	if tIDENT != tok {
		if !isKeyword(tok) {
			return pr.unexpected(lit, "package identifier", p)
		}
	}
	p.Name = lit
	return nil
}

// Accept dispatches the call to the visitor.
func (p *Package) Accept(v Visitor) {
	v.VisitPackage(p)
}

// inlineComment is part of commentInliner.
func (p *Package) inlineComment(c *Comment) {
	p.Comment = c
}

// columns returns printable source tokens
func (p *Package) columns() (cols []aligned) {
	cols = append(cols, notAligned("package "), notAligned(p.Name), alignedSemicolon)
	if p.Comment != nil {
		cols = append(cols, notAligned(" //"), notAligned(p.Comment.Message))
	}
	return
}
