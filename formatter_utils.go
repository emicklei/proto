package proto

import (
	"io"
	"strings"
)

func (f *Formatter) begin(stmt string) {
	if f.lastStmt != stmt && len(f.lastStmt) > 0 { // not the first line
		// add separator because stmt is changed, unless it nested thingy
		if !strings.Contains("comment enum service", f.lastStmt) {
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
	if len(list) == 0 {
		return
	}
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
	for i, each := range values {
		if i > 0 {
			f.nl()
		}
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
	f.nl()
}

// paddedTo return the word padding with spaces to match the length.
func paddedTo(word string, length int) string {
	if len(word) >= length {
		return word
	}
	return word + strings.Repeat(" ", length-len(word))
}

func (f *Formatter) nl() *Formatter {
	io.WriteString(f.w, "\n")
	return f
}

func (f *Formatter) printAsGroups(list []Visitee) {
	if len(list) == 0 {
		return
	}
	group := []columnsPrintable{}
	lastGroupName := ""
	for i := 0; i < len(list); i++ {
		each := list[i]
		groupName := nameOfVisitee(each)
		printable, isColumnsPrintable := each.(columnsPrintable)
		if isColumnsPrintable {
			if lastGroupName != groupName {
				// print current group
				if len(group) > 0 {
					f.printListOfColumns(group)
					lastGroupName = groupName
					// begin new group
					group = []columnsPrintable{}
				}
			}
			group = append(group, printable)
		} else {
			// not printable in group
			// print current group
			if len(group) > 0 {
				f.printListOfColumns(group)
				lastGroupName = groupName
				// begin new group
				group = []columnsPrintable{}
			}
			each.Accept(f)
		}
	}
	// print last group
	f.printListOfColumns(group)
}
