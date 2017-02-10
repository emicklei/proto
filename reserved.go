package proto

import (
	"fmt"
	"strconv"
)

// Reserved statements declare a range of field numbers or field names that cannot be used in a message.
type Reserved struct {
	Ranges     []Range
	FieldNames []string
	Comment    *Comment
}

// inlineComment is part of commentInliner.
func (r *Reserved) inlineComment(c *Comment) {
	r.Comment = c
}

// Range is to specify number intervals
type Range struct {
	From, To int
}

// String return a single number if from = to. Returns <from> to <to> otherwise.
func (r Range) String() string {
	if r.From == r.To {
		return strconv.Itoa(r.From)
	}
	return fmt.Sprintf("%d to %d", r.From, r.To)
}

// Accept dispatches the call to the visitor.
func (r *Reserved) Accept(v Visitor) {
	v.VisitReserved(r)
}

func (r *Reserved) parse(p *Parser) error {
	seenRangeTo := false
	for {
		tok, lit := p.scanIgnoreWhitespace()
		if len(lit) == 0 {
			return p.unexpected(lit, "reserved string or integer", r)
		}
		// first char that determined tok
		ch := []rune(lit)[0]
		if isDigit(ch) {
			// use unread here because scanInteger does not look at buf
			p.s.unread(ch)
			i, err := p.s.scanInteger()
			if err != nil {
				return p.unexpected(lit, "reserved integer", r)
			}
			if seenRangeTo {
				// replace last two ranges with one
				if len(r.Ranges) < 1 {
					p.unexpected(lit, "reserved integer", r)
				}
				from := r.Ranges[len(r.Ranges)-1]
				r.Ranges = append(r.Ranges[0:len(r.Ranges)-1], Range{From: from.From, To: i})
				seenRangeTo = false
			} else {
				r.Ranges = append(r.Ranges, Range{From: i, To: i})
			}
			continue
		}
		if tIDENT == tok && "to" == lit {
			seenRangeTo = true
			continue
		}
		if tQUOTE == tok || tSINGLEQUOTE == tok {
			// use unread here because scanLiteral does not look at buf
			p.s.unread(ch)
			field, isString := p.s.scanLiteral()
			if !isString {
				return p.unexpected(lit, "reserved string", r)
			}
			r.FieldNames = append(r.FieldNames, field)
			continue
		}
		if tSEMICOLON == tok {
			p.unscan()
			break
		}
	}
	return nil
}
