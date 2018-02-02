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

// ParentAwareVisitor keeps track of all the parents of each AST node
// while visiting all the elements and options.
type ParentAwareVisitor struct {
	Parents  []Visitee
	delegate Visitor
}

// NewParentAwareVisitor returns a ParentAwareVisitor that collects the parents
// of each visited and delegates the Visit call to the delegate.
// The delegate typically has a reference to the *ParentAwareVisitor (embbeded or normal field)
func NewParentAwareVisitor(delegate Visitor) *ParentAwareVisitor {
	return &ParentAwareVisitor{delegate: delegate}
}

// LastParent returns the immediat parent of the current Visitee accepted (and visited) by the visitor.
func (v *ParentAwareVisitor) LastParent() Visitee {
	if len(v.Parents) == 0 {
		return nil
	}
	return v.Parents[len(v.Parents)-1]
}

// VisitMessage is part of Visitor
func (v *ParentAwareVisitor) VisitMessage(m *Message) {
	v.push(m)
	defer v.pop()
	for _, each := range m.Elements {
		each.Accept(v.delegate)
	}
}

// VisitService is part of Visitor
func (v *ParentAwareVisitor) VisitService(s *Service) {
	v.push(s)
	defer v.pop()
	for _, each := range s.Elements {
		each.Accept(v.delegate)
	}
}

// VisitSyntax is part of Visitor
func (v *ParentAwareVisitor) VisitSyntax(s *Syntax) {}

// VisitPackage is part of Visitor
func (v *ParentAwareVisitor) VisitPackage(p *Package) {}

// VisitOption is part of Visitor
func (v *ParentAwareVisitor) VisitOption(o *Option) {}

// VisitImport is part of Visitor
func (v *ParentAwareVisitor) VisitImport(i *Import) {}

// VisitNormalField is part of Visitor
func (v *ParentAwareVisitor) VisitNormalField(i *NormalField) {
	v.push(i)
	defer v.pop()
	for _, each := range i.Options {
		each.Accept(v.delegate)
	}
}

// VisitEnumField is part of Visitor
func (v *ParentAwareVisitor) VisitEnumField(i *EnumField) {
	v.push(i)
	defer v.pop()
	i.ValueOption.Accept(v.delegate)
}

// VisitEnum is part of Visitor
func (v *ParentAwareVisitor) VisitEnum(e *Enum) {
	v.push(e)
	defer v.pop()
	for _, each := range e.Elements {
		each.Accept(v)
	}
}

// VisitComment is part of Visitor
func (v *ParentAwareVisitor) VisitComment(e *Comment) {}

// VisitOneof is part of Visitor
func (v *ParentAwareVisitor) VisitOneof(o *Oneof) {
	v.push(o)
	defer v.pop()
	for _, each := range o.Elements {
		each.Accept(v.delegate)
	}
}

// VisitOneofField is part of Visitor
func (v *ParentAwareVisitor) VisitOneofField(o *OneOfField) {
	v.push(o)
	defer v.pop()
	for _, each := range o.Options {
		each.Accept(v.delegate)
	}
}

// VisitReserved is part of Visitor
func (v *ParentAwareVisitor) VisitReserved(rs *Reserved) {}

// VisitRPC is part of Visitor
func (v *ParentAwareVisitor) VisitRPC(rpc *RPC) {}

// VisitMapField is part of Visitor
func (v *ParentAwareVisitor) VisitMapField(f *MapField) {
	v.push(f)
	defer v.pop()
	for _, each := range f.Options {
		each.Accept(v.delegate)
	}
}

// VisitGroup is part of Visitor
func (v *ParentAwareVisitor) VisitGroup(g *Group) {
	v.push(g)
	defer v.pop()
	for _, each := range g.Elements {
		each.Accept(v.delegate)
	}
}

// VisitExtensions is part of Visitor
func (v *ParentAwareVisitor) VisitExtensions(e *Extensions) {}

func (v *ParentAwareVisitor) push(parent Visitee) {
	v.Parents = append(v.Parents, parent)
}

func (v *ParentAwareVisitor) pop() {
	v.Parents = v.Parents[:len(v.Parents)-1]
}
