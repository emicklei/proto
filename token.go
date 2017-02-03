package proto

// token represents a lexical token.
type token int

const (
	// Special tokens
	tILLEGAL token = iota
	tEOF
	tWS

	// Literals
	tIDENT

	// Misc characters
	tSEMICOLON   // ;
	tEQUALS      // =
	tQUOTE       // "
	tSINGLEQUOTE // '
	tLEFTPAREN   // (
	tRIGHTPAREN  // )
	tLEFTCURLY   // {
	tRIGHTCURLY  // }
	tLEFTSQUARE  // [
	tRIGHTSQUARE // ]
	tCOMMENT     // /
	tLESS        // <
	tGREATER     // >
	tCOMMA       // ,
	tDOT         // .

	// Keywords
	keywordsStart
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
	tMAP
	tRESERVED
	tENUM

	// BEGIN proto2
	tOPTIONAL
	tGROUP
	tEXTENSIONS
	tEXTEND
	tREQUIRED
	// END proto2
	keywordsEnd
)

// typeTokens exists for future validation
const typeTokens = "double float int32 int64 uint32 uint64 sint32 sint64 fixed32 sfixed32 sfixed64 bool string bytes"

// context dependent reserved words
const (
	iSTREAM = "stream"
)

// isKeyword returns if tok is in the keywords range
func isKeyword(tok token) bool {
	return keywordsStart < tok && tok < keywordsEnd
}
