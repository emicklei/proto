package main

import (
	"encoding/xml"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/emicklei/proto"
)

// XSDSchema represents a schema
type XSDSchema struct {
	XMLName            xml.Name `xml:"schema"`
	StandardNamespace  string   `xml:"xmlns,attr"`
	TargetNamespace    string   `xml:"targetNamespace,attr"`
	TargetAlias        string   `xml:"xmlns:target,attr"`
	Version            string   `xml:"xmlns:version,attr"`
	ElementFormDefault string   `xml:"elementFormDefault,attr"`
	Types              []XSDComplexType
}

func buildXSDSchema(target string) XSDSchema {
	return XSDSchema{
		StandardNamespace:  "http://www.w3.org/2001/XMLSchema",
		TargetNamespace:    target,
		TargetAlias:        target,
		Version:            "v1",
		ElementFormDefault: "qualified",
	}
}

// XSDComplexType represents a complexType
type XSDComplexType struct {
	XMLName  xml.Name    `xml:"complexType"`
	Name     string      `xml:"name,attr"`
	Comment  string      `xml:",comment"`
	Sequence XSDSequence `xml:"sequence"`
}

// XSDSequence represents a sequence as part of e.g. complexType
type XSDSequence struct {
	Elements []XSDElement `xml:"element"`
}

// XSDElement represents an element as part of e.g. sequence
type XSDElement struct {
	XMLName   xml.Name
	Name      string `xml:"name,attr"`
	Comment   string `xml:",comment"`
	Type      string `xml:"type,attr"`
	MinOccurs string `xml:"minOccurs,attr,omitempty"`
	MaxOccurs string `xml:"maxOccurs,attr,omitempty"`
}

func buildXSDTypes(def *proto.Proto) (list []XSDComplexType, err error) {
	for _, each := range def.Elements {
		if msg, ok := each.(*proto.Message); ok {
			list = append(list, buildComplexType(msg))
		} else {
			// if *oDebug {
			// 	log.Printf("skipped a %T\n", each)
			// }
		}
	}
	return list, nil
}

func buildComplexType(msg *proto.Message) XSDComplexType {
	ct := XSDComplexType{}
	ct.Name = msg.Name
	if msg.Comment != nil {
		ct.Comment = msg.Comment.Message()
	}
	sq := XSDSequence{}
	for _, other := range msg.Elements {
		if field, ok := other.(*proto.NormalField); ok {
			sq = withNormalFieldToSequence(field, sq)
		} else {
			// if *oDebug {
			// 	log.Printf("skipped a %T\n", other)
			// }
		}
	}
	ct.Sequence = sq
	return ct
}

func withNormalFieldToSequence(f *proto.NormalField, s XSDSequence) XSDSequence {
	el := XSDElement{}
	el.Name = f.Name
	if f.Comment != nil {
		el.Comment = strings.Join(f.Comment.Lines, "\n")
	}
	el.Type = mapProtoSimpleTypeToXSDSimpleType(f.Type)
	// proto 3 fields are always optional. TODO check proto version
	el.MinOccurs = "0"
	if f.Repeated {
		el.MaxOccurs = "unbounded"
	}
	s.Elements = append(s.Elements, el)
	return s
}

func mapProtoSimpleTypeToXSDSimpleType(pt string) string {
	switch pt {
	case "fixed32", "uint32", "int32", "sfixed32", "sint32":
		return "integer"
	case "uint64", "int64", "fixed64", "sfixed64", "sint64":
		return "long"
	default:
		// assume that target types are in uppercase
		r, _ := utf8.DecodeRuneInString(pt)
		if unicode.IsUpper(r) {
			return "target:" + pt
		}
		return pt
	}
}
