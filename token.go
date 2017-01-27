package proto3parser

// Token represents a lexical token.
type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	WS

	// Literals
	IDENT // main

	// Misc characters
	SEMICOLON   // ;
	EQUALS      // =
	QUOTE       // "
	LEFTPAREN   // (
	RIGHTPAREN  // )
	LEFTCURLY   // {
	RIGHTCURLY  // }
	LEFTSQUARE  // [
	RIGHTSQUARE // ]

	// Keywords
	SYNTAX
	SERVICE
	RPC
	RETURNS
	MESSAGE
	IMPORT
	PACKAGE
	OPTION
	REPEATED

	// Types?
	DOUBLE
	FLOAT
	INT32
	INT64
	UINT32
	UINT64
	SINT32
	SINT64
	FIXED32
	FIXED64
	SFIXED32
	SFIXED64
	BOOL
	STRING
	BYTES
	MESSAGETYPE
	ENUMTYPE

	// OneOf OneOfField
)

const TypeTokens = "double float int32 int64 uint32 uint64 sint32 sint64 fixed32 sfixed32 sfixed64 bool string bytes messageType enumType"
