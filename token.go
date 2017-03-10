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
	tCOLON       // :
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
