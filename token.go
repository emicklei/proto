package proto3

// token represents a lexical token.
type token int

const (
	// Special tokens
	tILLEGAL token = iota
	tEOF
	tWS

	// Literals
	tIDENT
	tTRUE
	tFALSE

	// Misc characters
	tSEMICOLON   // ;
	tEQUALS      // =
	tQUOTE       // "
	tLEFTPAREN   // (
	tRIGHTPAREN  // )
	tLEFTCURLY   // {
	tRIGHTCURLY  // }
	tLEFTSQUARE  // [
	tRIGHTSQUARE // ]
	tCOMMENT     // /

	// Keywords
	tSYNTAX
	tSERVICE
	tRPC
	tRETURNS
	tMESSAGE
	tIMPORT
	tPACKAGE
	tOPTION
	tREPEATED
	tWEAK
	tPUBLIC

	// special fields
	tONEOF
	tONEOFFIELD
	tMAP
	tRESERVED
	tENUM
)

const typeTokens = "double float int32 int64 uint32 uint64 sint32 sint64 fixed32 sfixed32 sfixed64 bool string bytes"
