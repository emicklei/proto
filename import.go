package proto

import "fmt"

// Import holds a filename to another .proto definition.
type Import struct {
	Filename string
	Kind     string // weak, public, <empty>
	Comment  *Comment
}

func (i *Import) parse(p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	switch tok {
	case tWEAK:
		i.Kind = lit
		return i.parse(p)
	case tPUBLIC:
		i.Kind = lit
		return i.parse(p)
	case tQUOTE:
		i.Filename = p.s.scanUntil('"')
	case tSINGLEQUOTE:
		i.Filename = p.s.scanUntil('\'')
	default:
		return p.unexpected(lit, "import classifier weak|public|quoted", i)
	}
	return nil
}

// Accept dispatches the call to the visitor.
func (i *Import) Accept(v Visitor) {
	v.VisitImport(i)
}

// inlineComment is part of commentInliner.
func (i *Import) inlineComment(c *Comment) {
	i.Comment = c
}

// columns returns printable source tokens
func (i *Import) columns() (cols []aligned) {
	cols = append(cols, leftAligned("import "))
	if len(i.Kind) > 0 {
		cols = append(cols, leftAligned(i.Kind))
	} else {
		cols = append(cols, alignedEmpty)
	}
	cols = append(cols, alignedSpace, notAligned(fmt.Sprintf("%q", i.Filename)), alignedSemicolon)
	if i.Comment != nil {
		cols = append(cols, notAligned(" //"), notAligned(i.Comment.Message))
	}
	return
}
