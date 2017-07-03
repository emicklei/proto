package main

import "github.com/emicklei/proto"
import "encoding/xml"

type XSDComplexType struct {
	XMLName  xml.Name    `xml:"complexType"`
	Name     string      `xml:"name,attr"`
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
	Type      string `xml:"type,attr"`
	MinOccurs string `xml:"minOccurs,attr,omitempty"`
	MaxOccurs string `xml:"maxOccurs,attr,omitempty"`
}

func buildXSDTypes(def *proto.Proto) (list []XSDComplexType, err error) {

	for _, each := range def.Elements {
		if msg, ok := each.(*proto.Message); ok {

			ct := XSDComplexType{}
			ct.Name = msg.Name
			sq := XSDSequence{}
			for _, other := range msg.Elements {
				if field, ok := other.(*proto.NormalField); ok {
					sq = withNormalFieldToSequence(field, sq)
				}
			}
			ct.Sequence = sq
			list = append(list, ct)
		}
	}
	return list, nil
}

func withNormalFieldToSequence(f *proto.NormalField, s XSDSequence) XSDSequence {
	el := XSDElement{}
	el.Name = f.Name
	el.Type = mapProtoSimpleTypeToXSDSimpleType(f.Type)
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
		return pt
	}
}
