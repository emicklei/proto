package proto

// Extensions declare that a range of field numbers in a message are available for third-party extensions.
// proto2 only
type Extensions struct {
	Ranges  string
	Comment *Comment
}

// inlineComment is part of commentInliner.
func (e *Extensions) inlineComment(c *Comment) {
	e.Comment = c
}

// Accept dispatches the call to the visitor.
func (e *Extensions) Accept(v Visitor) {
	v.VisitExtensions(e)
}

// parse expects ident { messageBody
func (e *Extensions) parse(p *Parser) error {
	// TODO proper range parsing
	e.Ranges = p.s.scanUntil(';')
	p.s.unread(';') // for reading inline comment
	return nil
}

// columns returns printable source tokens
func (e *Extensions) columns() (cols []aligned) {
	cols = append(cols,
		notAligned("extensions "),
		leftAligned(e.Ranges),
		alignedSemicolon)
	if e.Comment != nil {
		cols = append(cols, notAligned(" //"), notAligned(e.Comment.Message))
	}
	return
}
