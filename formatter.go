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
	for _, each := range p.Elements {
		each.Accept(f)
	}
}

// VisitComment formats a Comment.
func (f *Formatter) VisitComment(c *Comment) {
	f.begin("comment")
	if c.IsMultiline() {
		fmt.Fprintln(f.w, "/*")
		fmt.Fprint(f.w, strings.TrimSpace(c.Message))
		fmt.Fprintf(f.w, "\n*/\n")
	} else {
		fmt.Fprintf(f.w, "//%s\n", c.Message)
	}
}

// VisitEnum formats a Enum.
func (f *Formatter) VisitEnum(e *Enum) {
	f.begin("enum")
	fmt.Fprintf(f.w, "enum %s {", e.Name)
	f.indentLevel++
	f.printAsGroups(e.Elements)
	f.indent(-1)
	io.WriteString(f.w, "}\n")
}

// VisitEnumField formats a EnumField.
func (f *Formatter) VisitEnumField(e *EnumField) {
	f.begin("field")
	io.WriteString(f.w, paddedTo(e.Name, 10))
	fmt.Fprintf(f.w, " = %d", e.Integer)
	if e.ValueOption != nil {
		io.WriteString(f.w, " ")
		e.ValueOption.Accept(f)
	} else {
		io.WriteString(f.w, ";\n")
	}
}

// VisitImport formats a Import.
func (f *Formatter) VisitImport(i *Import) {
	f.begin("import")
	if len(i.Kind) > 0 {
		fmt.Fprintf(f.w, "import %s ", i.Kind)
	}
	fmt.Fprintf(f.w, "import %q;\n", i.Filename)
}

// VisitMessage formats a Message.
func (f *Formatter) VisitMessage(m *Message) {
	f.begin("message")
	fmt.Fprintf(f.w, "message %s {", m.Name)
	f.newLineIf(len(m.Elements) > 0)
	f.indentLevel++
	for _, each := range m.Elements {
		each.Accept(f)
	}
	f.indent(-1)
	io.WriteString(f.w, "}\n")
}

// VisitOption formats a Option.
func (f *Formatter) VisitOption(o *Option) {
	if o.IsEmbedded {
		io.WriteString(f.w, "[(")
	} else {
		f.begin("option")
		io.WriteString(f.w, "option ")
	}
	if len(o.Name) > 0 {
		io.WriteString(f.w, o.Name)
	}
	if o.IsEmbedded {
		io.WriteString(f.w, ")")
	}
	io.WriteString(f.w, " = ")
	io.WriteString(f.w, o.Constant.String())
	if o.IsEmbedded {
		io.WriteString(f.w, "];\n")
	} else {
		io.WriteString(f.w, ";\n")
	}
}

// VisitPackage formats a Package.
func (f *Formatter) VisitPackage(p *Package) {
	f.begin("package")
	fmt.Fprintf(f.w, "package %s;\n", p.Name)
}

// VisitService formats a Service.
func (f *Formatter) VisitService(s *Service) {
	f.begin("service")
	fmt.Fprintf(f.w, "service %s {", s.Name)
	f.indentLevel++
	for _, each := range s.Elements {
		each.Accept(f)
	}
	f.indent(-1)
	io.WriteString(f.w, "}\n")
}

// VisitSyntax formats a Syntax.
func (f *Formatter) VisitSyntax(s *Syntax) {
	fmt.Fprintf(f.w, "syntax = %q;\n\n", s.Value)
}

// VisitOneof formats a Oneof.
func (f *Formatter) VisitOneof(o *Oneof) {
	f.begin("oneof")
	fmt.Fprintf(f.w, "oneof %s {", o.Name)
	f.indentLevel++
	for _, each := range o.Elements {
		each.Accept(f)
	}
	f.indent(-1)
	io.WriteString(f.w, "}\n")
}

// VisitOneofField formats a OneofField.
func (f *Formatter) VisitOneofField(o *OneOfField) {
	f.begin("oneoffield")
	fmt.Fprintf(f.w, "%s %s = %d", o.Type, o.Name, o.Sequence)
	for _, each := range o.Options {
		f.VisitOption(each)
	}
	io.WriteString(f.w, ";\n")
}

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
func (f *Formatter) VisitRPC(r *RPC) {
	f.begin("rpc")
	fmt.Fprintf(f.w, "rpc %s (", r.Name)
	if r.StreamsRequest {
		io.WriteString(f.w, "stream ")
	}
	io.WriteString(f.w, r.RequestType)
	io.WriteString(f.w, ") returns (")
	if r.StreamsReturns {
		io.WriteString(f.w, "stream ")
	}
	io.WriteString(f.w, r.ReturnsType)
	io.WriteString(f.w, ");\n")
}

// VisitMapField formats a MapField.
func (f *Formatter) VisitMapField(m *MapField) {
	f.begin("map")
	fmt.Fprintf(f.w, "map<%s,%s> %s = %d;\n", m.KeyType, m.Type, m.Name, m.Sequence)
}

// VisitNormalField formats a NormalField.
func (f *Formatter) VisitNormalField(f1 *NormalField) {
	f.begin("field")
	if f1.Repeated {
		io.WriteString(f.w, "repeated ")
	}
	if f1.Optional {
		io.WriteString(f.w, "optional ")
	}
	fmt.Fprintf(f.w, "%s %s = %d;\n", f1.Type, f1.Name, f1.Sequence)
}
