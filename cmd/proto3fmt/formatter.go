package main

import (
	"errors"
	"fmt"
	"io"

	"strings"

	"github.com/emicklei/proto3"
)

type formatter struct {
	w               io.Writer
	indentLevel     int
	lastStmt        string
	indentSeparator string
}

func (f *formatter) VisitComment(c *proto3.Comment) {
	f.begin("comment")
	if c.IsMultiline() {
		fmt.Fprintln(f.w, "/*")
		fmt.Fprint(f.w, strings.TrimSpace(c.Message))
		fmt.Fprintf(f.w, "\n*/\n")
	} else {
		fmt.Fprintf(f.w, "//%s\n", c.Message)
	}
}

func (f *formatter) VisitEnum(e *proto3.Enum) {
	f.begin("enum")
	fmt.Fprintf(f.w, "enum %s {\n", e.Name)
	f.indentLevel++
	for _, each := range e.Elements {
		each.Accept(f)
	}
	f.indent(-1)
	io.WriteString(f.w, "}\n")
}

func (f *formatter) VisitEnumField(e *proto3.EnumField) {
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

func (f *formatter) VisitField(f1 *proto3.Field) {
	f.begin("field")
	if f1.Repeated {
		io.WriteString(f.w, "repeated ")
	}
	fmt.Fprintf(f.w, "%s %s = %d;\n", f1.Type, f1.Name, f1.Sequence)
}

func (f *formatter) VisitImport(i *proto3.Import) {
	f.begin("import")
	if len(i.Kind) > 0 {
		fmt.Fprintf(f.w, "%s ", i.Kind)
	}
	fmt.Fprintf(f.w, "%q;\n", i.Filename)
}

func (f *formatter) VisitMessage(m *proto3.Message) {
	f.begin("message")
	fmt.Fprintf(f.w, "message %s {\n", m.Name)
	f.indentLevel++
	for _, each := range m.Elements {
		each.Accept(f)
	}
	f.indentLevel++
	io.WriteString(f.w, "}\n")
}

func (f *formatter) VisitOption(o *proto3.Option) {
	if o.PartOfFieldOrEnum {
		io.WriteString(f.w, "[(")
	} else {
		f.begin("option")
		io.WriteString(f.w, "option ")
	}
	if len(o.Name) > 0 {
		io.WriteString(f.w, o.Name)
	}
	if o.PartOfFieldOrEnum {
		io.WriteString(f.w, ")")
	}
	io.WriteString(f.w, " = ")
	if len(o.String) > 0 {
		fmt.Fprintf(f.w, "%q", o.String)
	} else {
		fmt.Fprintf(f.w, "%s", o.Identifier)
	}
	if o.PartOfFieldOrEnum {
		io.WriteString(f.w, "];\n")
	} else {
		io.WriteString(f.w, ";\n")
	}
}

func (f *formatter) VisitPackage(p *proto3.Package) {
	f.begin("package")
	fmt.Fprintf(f.w, "package %s;\n", p.Name)
}

func (f *formatter) VisitService(s *proto3.Service) {
	panic(errors.New("Not implemented"))
}

func (f *formatter) VisitSyntax(s *proto3.Syntax) {
	fmt.Fprintf(f.w, "syntax = %q;\n\n", s.Value)
}

func (f *formatter) VisitOneof(o *proto3.Oneof) {
	panic(errors.New("Not implemented"))
}

func (f *formatter) VisitOneofField(o *proto3.OneOfField) {
	panic(errors.New("Not implemented"))
}

func (f *formatter) VisitReserved(r *proto3.Reserved) {
	panic(errors.New("VisitReserved Not implemented"))
}

// Utils

func (f *formatter) begin(stmt string) {
	if f.lastStmt != stmt && len(f.lastStmt) > 0 { // not the first line
		// add separator because stmt is changed, unless it was comment or a nested thingy
		if !strings.Contains("comment message enum", f.lastStmt) {
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
