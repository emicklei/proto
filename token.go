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
	COLON      // ;
	EQUALS     // =
	QUOTE      // "
	LEFTPAREN  // (
	RIGHTPAREN // )
	LEFTCURLY  // {
	RIGHTCURLY // }

	// Keywords
	SYNTAX
	SERVICE
	RPC
	RETURNS
	MESSAGE
)
