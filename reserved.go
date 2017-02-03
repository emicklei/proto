package proto

import (
	"fmt"
	"strconv"
	"strings"
)

// Reserved statements declare a range of field numbers or field names that cannot be used in a message.
type Reserved struct {
	Ranges     string
	FieldNames []string
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
	// TODO proper parsing ranges
	content := strings.TrimSpace(p.s.scanUntil(';'))
	if strings.Contains(content, "\"") {
		quoted := strings.Split(content, ",")
		for _, each := range quoted {
			r.FieldNames = append(r.FieldNames, strings.Trim(strings.TrimSpace(each), "\""))
		}
	} else {
		r.Ranges = content
	}
	return nil
}
