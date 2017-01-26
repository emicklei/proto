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
	OPTIONAL
	REPEATED
)
