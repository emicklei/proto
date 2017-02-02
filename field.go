package proto

import "strconv"

// Field is an abstract message field.
type Field struct {
	Name     string
	Type     string
	Sequence int
	Options  []*Option
}

// NormalField
type NormalField struct {
	*Field
	Repeated bool
	Optional bool // proto2
}

func newNormalField() *NormalField { return &NormalField{Field: new(Field)} }

// Accept dispatches the call to the visitor.
func (f *NormalField) Accept(v Visitor) {
	v.VisitNormalField(f)
}

// columns returns printable source tokens
func (f *NormalField) columns() (cols []aligned) {
	if f.Repeated {
		cols = append(cols, leftAligned("repeated "))
	} else {
		cols = append(cols, alignedSpace)
	}
	cols = append(cols, leftAligned(f.Name), rightAligned(f.Type), alignedEquals, rightAligned(strconv.Itoa(f.Sequence)))
	if len(f.Options) > 0 {
		cols = append(cols, leftAligned(" ["))
		for i, each := range f.Options {
			if i > 0 {
				cols = append(cols, alignedComma)
			}
			cols = append(cols, each.keyValuePair()...)
		}
		cols = append(cols, leftAligned("]"))
	}
	return
}

// parse expects:
// [ "repeated" | "optional" ] type fieldName "=" fieldNumber [ "[" fieldOptions "]" ] ";"
func (f *NormalField) parse(p *Parser) error {
	for {
		tok, lit := p.scanIgnoreWhitespace()
		switch tok {
		case tREPEATED:
			f.Repeated = true
			return f.parse(p)
		case tOPTIONAL: // proto2
			f.Optional = true
			return f.parse(p)
		case tIDENT:
			f.Type = lit
			return parseFieldAfterType(f.Field, p)
		default:
			goto done
		}
	}
done:
	return nil
}

// parseFieldAfterType expects:
// fieldName "=" fieldNumber [ "[" fieldOptions "]" ] ";
func parseFieldAfterType(f *Field, p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != tIDENT {
		return p.unexpected(lit, "identifier")
	}
	f.Name = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tok != tEQUALS {
		return p.unexpected(lit, "=")
	}
	i, err := p.s.scanInteger()
	if err != nil {
		return p.unexpected(lit, "sequence number")
	}
	f.Sequence = i
	// see if there are options
	tok, lit = p.scanIgnoreWhitespace()
	if tLEFTSQUARE != tok {
		p.unscan()
		return nil
	}
	// consume options
	for {
		o := new(Option)
		o.IsEmbedded = true
		err := o.parse(p)
		if err != nil {
			return err
		}
		f.Options = append(f.Options, o)

		tok, lit = p.scanIgnoreWhitespace()
		if tRIGHTSQUARE == tok {
			break
		}
		if tCOMMA != tok {
			return p.unexpected(lit, ",")
		}
	}
	return nil
}

// MapField represents a map entry in a message.
type MapField struct {
	*Field
	KeyType string
}

func newMapField() *MapField { return &MapField{Field: new(Field)} }

// Accept dispatches the call to the visitor.
func (f *MapField) Accept(v Visitor) {
	v.VisitMapField(f)
}

// parse expects:
// mapField = "map" "<" keyType "," type ">" mapName "=" fieldNumber [ "[" fieldOptions "]" ] ";"
// keyType = "int32" | "int64" | "uint32" | "uint64" | "sint32" | "sint64" |
//           "fixed32" | "fixed64" | "sfixed32" | "sfixed64" | "bool" | "string"
func (f *MapField) parse(p *Parser) error {
	tok, lit := p.scanIgnoreWhitespace()
	if tLESS != tok {
		return p.unexpected(lit, "<")
	}
	tok, lit = p.scanIgnoreWhitespace()
	if tIDENT != tok {
		return p.unexpected(lit, "identifier")
	}
	f.KeyType = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tCOMMA != tok {
		return p.unexpected(lit, ",")
	}
	tok, lit = p.scanIgnoreWhitespace()
	if tIDENT != tok {
		return p.unexpected(lit, "identifier")
	}
	f.Type = lit
	tok, lit = p.scanIgnoreWhitespace()
	if tGREATER != tok {
		return p.unexpected(lit, ">")
	}
	return parseFieldAfterType(f.Field, p)
}
