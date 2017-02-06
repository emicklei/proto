package proto

import (
	"fmt"
	"io"
	"strings"
)

// Formatter visits a Proto and writes formatted source.
type Formatter struct {
	w               io.Writer
	indentLevel     int
	lastStmt        string
	indentSeparator string
}

// NewFormatter returns a new Formatter. Only the indentation separator is configurable.
func NewFormatter(writer io.Writer, indentSeparator string) *Formatter {
	return &Formatter{w: writer, indentSeparator: indentSeparator}
}

// Format visits all proto elements and writes formatted source.
func (f *Formatter) Format(p *Proto) {
	f.printAsGroups(p.Elements)
}

// VisitComment formats a Comment.
func (f *Formatter) VisitComment(c *Comment) {
	f.begin("comment")
	if c.IsMultiline() {
		fmt.Fprintln(f.w, "/*")
		lines := strings.Split(c.Message, "\n")
		for i, each := range lines {
			// leading no tab or space
			leftAligned := strings.TrimLeft(each, "\t ")
			// only skip first and last empty lines
			skip := (i == 0 && len(leftAligned) == 0) ||
				(i == len(lines)-1 && len(leftAligned) == 0)
			if !skip {
				f.indent(0)
				fmt.Fprintf(f.w, " %s\n", leftAligned)
			}
		}
		f.indent(0)
		fmt.Fprintf(f.w, " */\n")
	} else {
		fmt.Fprintf(f.w, "//%s\n", c.Message)
	}
}

// VisitEnum formats a Enum.
func (f *Formatter) VisitEnum(e *Enum) {
	f.begin("enum")
	fmt.Fprintf(f.w, "enum %s {", e.Name)
	if len(e.Elements) > 0 {
		f.nl()
		f.indentLevel++
		f.printAsGroups(e.Elements)
		f.indent(-1)
	}
	io.WriteString(f.w, "}\n")
	f.end("enum")
}

// VisitEnumField formats a EnumField.
func (f *Formatter) VisitEnumField(e *EnumField) {}

// VisitImport formats a Import.
func (f *Formatter) VisitImport(i *Import) {}

// VisitMessage formats a Message.
func (f *Formatter) VisitMessage(m *Message) {
	f.begin("message")
	fmt.Fprintf(f.w, "message %s {", m.Name)
	if len(m.Elements) > 0 {
		f.nl()
		f.indentLevel++
		f.printAsGroups(m.Elements)
		f.indent(-1)
	}
	io.WriteString(f.w, "}\n")
	f.end("message")
}

// VisitOption formats a Option.
func (f *Formatter) VisitOption(o *Option) {}

// VisitPackage formats a Package.
func (f *Formatter) VisitPackage(p *Package) {
	f.begin("package")
	fmt.Fprintf(f.w, "package %s;\n\n", p.Name)
}

// VisitService formats a Service.
func (f *Formatter) VisitService(s *Service) {
	f.begin("service")
	fmt.Fprintf(f.w, "service %s {", s.Name)
	if len(s.Elements) > 0 {
		f.nl()
		f.indentLevel++
		f.printAsGroups(s.Elements)
		f.indent(-1)
	}
	io.WriteString(f.w, "}\n")
	f.end("service")
}

// VisitSyntax formats a Syntax.
func (f *Formatter) VisitSyntax(s *Syntax) {
	fmt.Fprintf(f.w, "syntax = %q;\n\n", s.Value)
}

// VisitOneof formats a Oneof.
func (f *Formatter) VisitOneof(o *Oneof) {
	f.begin("oneof")
	fmt.Fprintf(f.w, "oneof %s {", o.Name)
	if len(o.Elements) > 0 {
		f.nl()
		f.indentLevel++
		f.printAsGroups(o.Elements)
		f.indent(-1)
	}
	io.WriteString(f.w, "}\n")
	f.end("oneof")
}

// VisitOneofField formats a OneofField.
func (f *Formatter) VisitOneofField(o *OneOfField) {}

// VisitReserved formats a Reserved.
func (f *Formatter) VisitReserved(r *Reserved) {
	f.begin("reserved")
	io.WriteString(f.w, "reserved ")
	if len(r.Ranges) > 0 {
		io.WriteString(f.w, r.Ranges)
	} else {
		for i, each := range r.FieldNames {
			if i > 0 {
				io.WriteString(f.w, ",")
			}
			fmt.Fprintf(f.w, "%q", each)
		}
	}
	io.WriteString(f.w, ";\n")
}

// VisitRPC formats a RPC.
func (f *Formatter) VisitRPC(r *RPC) {}

// VisitMapField formats a MapField.
func (f *Formatter) VisitMapField(m *MapField) {
	f.begin("map")
	fmt.Fprintf(f.w, "map<%s,%s> %s = %d;\n", m.KeyType, m.Type, m.Name, m.Sequence)
}

// VisitNormalField formats a NormalField.
func (f *Formatter) VisitNormalField(f1 *NormalField) {}

// VisitGroup formats a proto2 Group.
func (f *Formatter) VisitGroup(g *Group) {
	f.begin("group")
	if g.Optional {
		io.WriteString(f.w, "optional ")
	}
	fmt.Fprintf(f.w, "group %s = %d {", g.Name, g.Sequence)
	if len(g.Elements) > 0 {
		f.nl()
		f.indentLevel++
		f.printAsGroups(g.Elements)
		f.indent(-1)
	}
	io.WriteString(f.w, "}\n")
	f.end("group")
}

// VisitExtensions formats a proto2 Extensions.
func (f *Formatter) VisitExtensions(e *Extensions) {
	f.indent(0)
	fmt.Fprintf(f.w, "extensions %s;\n", e.Ranges)
}
