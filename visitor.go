// Copyright (c) 2017 Ernest Micklei
//
// MIT License
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package proto

import "log"

// Visitor is for dispatching Proto elements.
type Visitor interface {
	VisitProto(p *Proto)
	VisitMessage(m *Message)
	VisitService(v *Service)
	VisitSyntax(s *Syntax)
	VisitPackage(p *Package)
	VisitOption(o *Option)
	VisitImport(i *Import)
	VisitNormalField(i *NormalField)
	VisitEnumField(i *EnumField)
	VisitEnum(e *Enum)
	VisitComment(e *Comment)
	VisitOneof(o *Oneof)
	VisitOneofField(o *OneOfField)
	VisitReserved(r *Reserved)
	VisitRPC(r *RPC)
	VisitMapField(f *MapField)
	// proto2
	VisitGroup(g *Group)
	VisitExtensions(e *Extensions)
}

// Visitee is implemented by all Proto elements.
type Visitee interface {
	Accept(v Visitor)
}

// Documented is for types that may have an associated comment (not inlined).
type Documented interface {
	Doc() *Comment
}

// reflector is a Visitor that can tell the short type name of a Visitee.
type reflector struct {
	name string
}

// sole instance of reflector
var namer = new(reflector)

func (r *reflector) VisitMessage(m *Message)         { r.name = "Message" }
func (r *reflector) VisitService(v *Service)         { r.name = "Service" }
func (r *reflector) VisitSyntax(s *Syntax)           { r.name = "Syntax" }
func (r *reflector) VisitPackage(p *Package)         { r.name = "Package" }
func (r *reflector) VisitOption(o *Option)           { r.name = "Option" }
func (r *reflector) VisitImport(i *Import)           { r.name = "Import" }
func (r *reflector) VisitNormalField(i *NormalField) { r.name = "NormalField" }
func (r *reflector) VisitEnumField(i *EnumField)     { r.name = "EnumField" }
func (r *reflector) VisitEnum(e *Enum)               { r.name = "Enum" }
func (r *reflector) VisitComment(e *Comment)         { r.name = "Comment" }
func (r *reflector) VisitOneof(o *Oneof)             { r.name = "Oneof" }
func (r *reflector) VisitOneofField(o *OneOfField)   { r.name = "OneOfField" }
func (r *reflector) VisitReserved(rs *Reserved)      { r.name = "Reserved" }
func (r *reflector) VisitRPC(rpc *RPC)               { r.name = "RPC" }
func (r *reflector) VisitMapField(f *MapField)       { r.name = "MapField" }
func (r *reflector) VisitGroup(g *Group)             { r.name = "Group" }
func (r *reflector) VisitExtensions(e *Extensions)   { r.name = "Extensions" }
func (r *reflector) VisitProto(p *Proto)             { r.name = "Proto" }

// nameOfVisitee returns the short type name of a Visitee.
func nameOfVisitee(e Visitee) string {
	e.Accept(namer)
	return namer.name
}

func setParent(child Visitee, parent Visitee) {
	if child == nil {
		log.Fatal("child is nil")
	}
	if parent == nil {
		log.Fatal("parent is nil")
	}
	child.Accept(&parentAccessor{isGet: false, parent: parent})
}

func getParent(child Visitee) Visitee {
	if child == nil {
		log.Fatal("child is nil")
	}
	pa := &parentAccessor{isGet: true}
	child.Accept(pa)
	return pa.parent
}

type parentAccessor struct {
	isGet  bool
	parent Visitee
}

func (p *parentAccessor) VisitMessage(m *Message) {
	if p.isGet {
		p.parent = m.Parent
	} else {
		m.Parent = p.parent
	}
}
func (p *parentAccessor) VisitService(v *Service) {
	if p.isGet {
		p.parent = v.Parent
	} else {
		v.Parent = p.parent
	}
}
func (p *parentAccessor) VisitSyntax(s *Syntax) {
	if p.isGet {
		p.parent = s.Parent
	} else {
		s.Parent = p.parent
	}
}
func (p *parentAccessor) VisitPackage(pkg *Package) {
	if p.isGet {
		p.parent = pkg.Parent
	} else {
		pkg.Parent = p.parent
	}
}
func (p *parentAccessor) VisitOption(o *Option) {
	if p.isGet {
		p.parent = o.Parent
	} else {
		o.Parent = p.parent
	}
}
func (p *parentAccessor) VisitImport(i *Import) {
	if p.isGet {
		p.parent = i.Parent
	} else {
		i.Parent = p.parent
	}
}
func (p *parentAccessor) VisitNormalField(i *NormalField) {
	if p.isGet {
		p.parent = i.Parent
	} else {
		i.Parent = p.parent
	}
}
func (p *parentAccessor) VisitEnumField(i *EnumField) {
	if p.isGet {
		p.parent = i.Parent
	} else {
		i.Parent = p.parent
	}
}
func (p *parentAccessor) VisitEnum(e *Enum) {
	if p.isGet {
		p.parent = e.Parent
	} else {
		e.Parent = p.parent
	}
}
func (p *parentAccessor) VisitComment(e *Comment) {}
func (p *parentAccessor) VisitOneof(o *Oneof) {
	if p.isGet {
		p.parent = o.Parent
	} else {
		o.Parent = p.parent
	}
}
func (p *parentAccessor) VisitOneofField(o *OneOfField) {
	if p.isGet {
		p.parent = o.Parent
	} else {
		o.Parent = p.parent
	}
}
func (p *parentAccessor) VisitReserved(rs *Reserved) {
	if p.isGet {
		p.parent = rs.Parent
	} else {
		rs.Parent = p.parent
	}
}
func (p *parentAccessor) VisitRPC(rpc *RPC) {
	if p.isGet {
		p.parent = rpc.Parent
	} else {
		rpc.Parent = p.parent
	}
}
func (p *parentAccessor) VisitMapField(f *MapField) {
	if p.isGet {
		p.parent = f.Parent
	} else {
		f.Parent = p.parent
	}
}
func (p *parentAccessor) VisitGroup(g *Group) {
	if p.isGet {
		p.parent = g.Parent
	} else {
		g.Parent = p.parent
	}
}
func (p *parentAccessor) VisitExtensions(e *Extensions) {
	if p.isGet {
		p.parent = e.Parent
	} else {
		e.Parent = p.parent
	}
}
func (p *parentAccessor) VisitProto(*Proto) {}
