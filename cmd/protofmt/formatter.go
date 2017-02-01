package main

import (
	"fmt"
	"io"

	"strings"

	"github.com/emicklei/proto"
)

type formatter struct {
	w               io.Writer
	indentLevel     int
	lastStmt        string
	indentSeparator string
}

func (f *formatter) VisitComment(c *proto.Comment) {
	f.begin("comment")
	if c.IsMultiline() {
		fmt.Fprintln(f.w, "/*")
		fmt.Fprint(f.w, strings.TrimSpace(c.Message))
		fmt.Fprintf(f.w, "\n*/\n")
	} else {
		fmt.Fprintf(f.w, "//%s\n", c.Message)
	}
}

func (f *formatter) VisitEnum(e *proto.Enum) {
	f.begin("enum")
	fmt.Fprintf(f.w, "enum %s {", e.Name)
	f.indentLevel++
	for _, each := range e.Elements {
		each.Accept(f)
	}
	f.indent(-1)
	io.WriteString(f.w, "}\n")
}

func (f *formatter) VisitEnumField(e *proto.EnumField) {
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

func (f *formatter) VisitImport(i *proto.Import) {
	f.begin("import")
	if len(i.Kind) > 0 {
		fmt.Fprintf(f.w, "import %s ", i.Kind)
	}
	fmt.Fprintf(f.w, "import %q;\n", i.Filename)
}

func (f *formatter) VisitMessage(m *proto.Message) {
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

func (f *formatter) VisitOption(o *proto.Option) {
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

func (f *formatter) VisitPackage(p *proto.Package) {
	f.begin("package")
	fmt.Fprintf(f.w, "package %s;\n", p.Name)
}

func (f *formatter) VisitService(s *proto.Service) {
	f.begin("service")
	fmt.Fprintf(f.w, "service %s {", s.Name)
	f.indentLevel++
	for _, each := range s.Elements {
		each.Accept(f)
	}
	f.indent(-1)
	io.WriteString(f.w, "}\n")
}

func (f *formatter) VisitSyntax(s *proto.Syntax) {
	fmt.Fprintf(f.w, "syntax = %q;\n\n", s.Value)
}

func (f *formatter) VisitOneof(o *proto.Oneof) {
	f.begin("oneof")
	fmt.Fprintf(f.w, "oneof %s {", o.Name)
	f.indentLevel++
	for _, each := range o.Elements {
		each.Accept(f)
	}
	f.indent(-1)
	io.WriteString(f.w, "}\n")
}

func (f *formatter) VisitOneofField(o *proto.OneOfField) {
	f.begin("oneoffield")
	fmt.Fprintf(f.w, "%s %s = %d", o.Type, o.Name, o.Sequence)
	for _, each := range o.Options {
		f.VisitOption(each)
	}
	io.WriteString(f.w, ";\n")
}

func (f *formatter) VisitReserved(r *proto.Reserved) {
	f.begin("reserved")
	io.WriteString(f.w, "reserved")
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

func (f *formatter) VisitRPC(r *proto.RPC) {
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

func (f *formatter) VisitMapField(m *proto.MapField) {
	f.begin("map")
	fmt.Fprintf(f.w, "map<%s,%s> %s = %d;\n", m.KeyType, m.Type, m.Name, m.Sequence)
}

func (f *formatter) VisitNormalField(f1 *proto.NormalField) {
	f.begin("field")
	if f1.Repeated {
		io.WriteString(f.w, "repeated ")
	}
	if f1.Optional {
		io.WriteString(f.w, "optional ")
	}
	fmt.Fprintf(f.w, "%s %s = %d;\n", f1.Type, f1.Name, f1.Sequence)
}

// Utils

func (f *formatter) begin(stmt string) {
	if f.lastStmt != stmt && len(f.lastStmt) > 0 { // not the first line
		// add separator because stmt is changed, unless it nested thingy
		if !strings.Contains("comment", f.lastStmt) {
			io.WriteString(f.w, "\n")
		}
	}
	f.indent(0)
	f.lastStmt = stmt
}

func (f *formatter) indent(diff int) {
	f.indentLevel += diff
	for i := 0; i < f.indentLevel; i++ {
		io.WriteString(f.w, f.indentSeparator)
	}
}

func paddedTo(word string, length int) string {
	if len(word) >= length {
		return word
	}
	return word + strings.Repeat(" ", length-len(word))
}

func (f *formatter) newLineIf(ok bool) {
	if ok {
		io.WriteString(f.w, "\n")
	}
}
