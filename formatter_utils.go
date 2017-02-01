package proto

import (
	"io"
	"strings"
)

func (f *Formatter) begin(stmt string) {
	if f.lastStmt != stmt && len(f.lastStmt) > 0 { // not the first line
		// add separator because stmt is changed, unless it nested thingy
		if !strings.Contains("comment", f.lastStmt) {
			io.WriteString(f.w, "\n")
		}
	}
	f.indent(0)
	f.lastStmt = stmt
}

func (f *Formatter) indent(diff int) {
	f.indentLevel += diff
	for i := 0; i < f.indentLevel; i++ {
		io.WriteString(f.w, f.indentSeparator)
	}
}

type columnsPrintable interface {
	columns() (cols []aligned)
}

func (f *Formatter) printListOfColumns(list []columnsPrintable) {
	// collect all column values
	values := [][]aligned{}
	widths := map[int]int{}
	for _, each := range list {
		cols := each.columns()
		values = append(values, cols)
		// update max widths per column
		for i, other := range cols {
			pw := other.preferredWidth()
			w, ok := widths[i]
			if ok {
				if pw > w {
					widths[i] = pw
				}
			} else {
				widths[i] = pw
			}
		}
	}
	// now print all values
	for _, each := range values {
		io.WriteString(f.w, "\n")
		f.indent(0)
		for c := 0; c < len(widths); c++ {
			pw := widths[c]
			// only print if there is a value
			if c < len(each) {
				// using space padding to match the max width
				src := each[c].formatted(pw)
				io.WriteString(f.w, src)
			}
		}
		io.WriteString(f.w, ";")
	}
	io.WriteString(f.w, "\n")
}

func paddedTo(word string, length int) string {
	if len(word) >= length {
		return word
	}
	return word + strings.Repeat(" ", length-len(word))
}

func (f *Formatter) newLineIf(ok bool) {
	if ok {
		io.WriteString(f.w, "\n")
	}
}

func (f *Formatter) printAsGroups(list []Visitee) {
	if len(list) == 0 {
		return
	}
	group := []columnsPrintable{}
	lastGroupName := nameOfVisitee(list[0])
	for i := 1; i < len(list); i++ {
		groupName := nameOfVisitee(list[i])
		printable, isColumnsPrintable := list[i].(columnsPrintable)
		if isColumnsPrintable {
			if lastGroupName == groupName {
				// collect in group
				group = append(group, printable)
			} else {
				// print current group
				f.printListOfColumns(group)
				lastGroupName = groupName
				// begin new group
				group = []columnsPrintable{printable}
			}
		} else {
			// not printable in group
			list[i].Accept(f)
		}
	}
	// print last group
	f.printListOfColumns(group)
}
