package proto3parser

// token represents a lexical token.
type token int

const (
	// Special tokens
	ILLEGAL token = iota
	EOF
	WS

	// Literals
	IDENT // main
	TRUE
	FALSE

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
	COMMENT     // /

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

	// special fields
	ONEOF
	ONEOFFIELD
	MAP
	RESERVED
	ENUM
)

const typeTokens = "double float int32 int64 uint32 uint64 sint32 sint64 fixed32 sfixed32 sfixed64 bool string bytes"
