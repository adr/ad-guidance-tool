// Code generated from ADE.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ADE

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ADEParser struct {
	*antlr.BaseParser
}

var ADEParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func adeParserInit() {
	staticData := &ADEParserStaticData
	staticData.LiteralNames = []string{
		"", "'adr'", "'file'", "'code'", "'component'", "'class'", "'interface'",
		"'path'", "'must'", "'not'", "'only'", "'depend'", "'exist'", "'contain'",
		"'implement'", "'extend'", "'annotated'", "'accessed'", "'acyclic'",
		"'on'", "'be'", "'by'", "'with'", "'match'", "'exclude'", "'implementing'",
		"'in'", "'public'", "'internal'", "'private'", "'abstract'", "'sealed'",
		"'static'", "'severity'", "'error'", "'warning'", "'='", "','", "'{'",
		"'}'",
	}
	staticData.SymbolicNames = []string{
		"", "ADR", "FILE", "CODE", "COMPONENT", "CLASS", "INTERFACE", "PATH",
		"MUST", "NOT", "ONLY", "DEPEND", "EXIST", "CONTAIN", "IMPLEMENT", "EXTEND",
		"ANNOTATED", "ACCESSED", "ACYCLIC", "ON", "BE", "BY", "WITH", "MATCH",
		"EXCLUDE", "IMPLEMENTING", "IN", "PUBLIC", "INTERNAL", "PRIVATE", "ABSTRACT",
		"SEALED", "STATIC", "SEVERITY", "ERROR", "WARNING", "EQ", "COMMA", "LBRACE",
		"RBRACE", "IDENTIFIER", "STRING", "COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"file", "adrDecl", "selectorDecl", "ruleDecl", "ruleType", "ruleStmt",
		"assertionStmt", "subjectExpr", "targetExpr", "selectorType", "mustExpr",
		"verbPhrase", "visibility", "typeConstraint", "excludeStmt", "severityStmt",
		"severityValue",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 43, 213, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 1, 0, 1, 0, 1, 0, 5, 0, 38, 8, 0, 10, 0, 12, 0, 41, 9, 0,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3,
		1, 3, 1, 3, 1, 3, 5, 3, 58, 8, 3, 10, 3, 12, 3, 61, 9, 3, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 5, 3, 5, 70, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7,
		100, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8,
		111, 8, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 120, 8,
		10, 1, 11, 1, 11, 3, 11, 124, 8, 11, 1, 11, 1, 11, 1, 11, 5, 11, 129, 8,
		11, 10, 11, 12, 11, 132, 9, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11,
		139, 8, 11, 1, 11, 1, 11, 1, 11, 3, 11, 144, 8, 11, 1, 11, 1, 11, 3, 11,
		148, 8, 11, 1, 11, 1, 11, 3, 11, 152, 8, 11, 1, 11, 1, 11, 3, 11, 156,
		8, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 163, 8, 11, 10, 11, 12,
		11, 166, 9, 11, 1, 11, 3, 11, 169, 8, 11, 1, 11, 1, 11, 3, 11, 173, 8,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 180, 8, 11, 1, 11, 1, 11,
		3, 11, 184, 8, 11, 1, 11, 3, 11, 187, 8, 11, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 3, 14, 206, 8, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16,
		1, 16, 0, 0, 17, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
		30, 32, 0, 5, 1, 0, 4, 7, 1, 0, 2, 3, 1, 0, 27, 29, 1, 0, 30, 32, 1, 0,
		34, 35, 237, 0, 34, 1, 0, 0, 0, 2, 44, 1, 0, 0, 0, 4, 48, 1, 0, 0, 0, 6,
		53, 1, 0, 0, 0, 8, 64, 1, 0, 0, 0, 10, 69, 1, 0, 0, 0, 12, 71, 1, 0, 0,
		0, 14, 99, 1, 0, 0, 0, 16, 110, 1, 0, 0, 0, 18, 112, 1, 0, 0, 0, 20, 119,
		1, 0, 0, 0, 22, 186, 1, 0, 0, 0, 24, 188, 1, 0, 0, 0, 26, 190, 1, 0, 0,
		0, 28, 205, 1, 0, 0, 0, 30, 207, 1, 0, 0, 0, 32, 210, 1, 0, 0, 0, 34, 39,
		3, 2, 1, 0, 35, 38, 3, 4, 2, 0, 36, 38, 3, 6, 3, 0, 37, 35, 1, 0, 0, 0,
		37, 36, 1, 0, 0, 0, 38, 41, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 39, 40, 1,
		0, 0, 0, 40, 42, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 42, 43, 5, 0, 0, 1, 43,
		1, 1, 0, 0, 0, 44, 45, 5, 1, 0, 0, 45, 46, 5, 41, 0, 0, 46, 47, 5, 41,
		0, 0, 47, 3, 1, 0, 0, 0, 48, 49, 7, 0, 0, 0, 49, 50, 5, 41, 0, 0, 50, 51,
		5, 36, 0, 0, 51, 52, 5, 41, 0, 0, 52, 5, 1, 0, 0, 0, 53, 54, 3, 8, 4, 0,
		54, 55, 5, 41, 0, 0, 55, 59, 5, 38, 0, 0, 56, 58, 3, 10, 5, 0, 57, 56,
		1, 0, 0, 0, 58, 61, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0,
		60, 62, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 62, 63, 5, 39, 0, 0, 63, 7, 1,
		0, 0, 0, 64, 65, 7, 1, 0, 0, 65, 9, 1, 0, 0, 0, 66, 70, 3, 12, 6, 0, 67,
		70, 3, 28, 14, 0, 68, 70, 3, 30, 15, 0, 69, 66, 1, 0, 0, 0, 69, 67, 1,
		0, 0, 0, 69, 68, 1, 0, 0, 0, 70, 11, 1, 0, 0, 0, 71, 72, 3, 14, 7, 0, 72,
		73, 3, 20, 10, 0, 73, 74, 3, 22, 11, 0, 74, 13, 1, 0, 0, 0, 75, 100, 5,
		40, 0, 0, 76, 77, 3, 18, 9, 0, 77, 78, 5, 41, 0, 0, 78, 79, 5, 26, 0, 0,
		79, 80, 3, 16, 8, 0, 80, 100, 1, 0, 0, 0, 81, 82, 3, 18, 9, 0, 82, 83,
		5, 23, 0, 0, 83, 84, 5, 41, 0, 0, 84, 85, 5, 26, 0, 0, 85, 86, 3, 16, 8,
		0, 86, 100, 1, 0, 0, 0, 87, 88, 3, 18, 9, 0, 88, 89, 5, 26, 0, 0, 89, 90,
		3, 16, 8, 0, 90, 100, 1, 0, 0, 0, 91, 92, 3, 18, 9, 0, 92, 93, 5, 41, 0,
		0, 93, 100, 1, 0, 0, 0, 94, 95, 3, 18, 9, 0, 95, 96, 5, 23, 0, 0, 96, 97,
		5, 41, 0, 0, 97, 100, 1, 0, 0, 0, 98, 100, 3, 18, 9, 0, 99, 75, 1, 0, 0,
		0, 99, 76, 1, 0, 0, 0, 99, 81, 1, 0, 0, 0, 99, 87, 1, 0, 0, 0, 99, 91,
		1, 0, 0, 0, 99, 94, 1, 0, 0, 0, 99, 98, 1, 0, 0, 0, 100, 15, 1, 0, 0, 0,
		101, 111, 5, 40, 0, 0, 102, 103, 3, 18, 9, 0, 103, 104, 5, 41, 0, 0, 104,
		111, 1, 0, 0, 0, 105, 106, 3, 18, 9, 0, 106, 107, 5, 23, 0, 0, 107, 108,
		5, 41, 0, 0, 108, 111, 1, 0, 0, 0, 109, 111, 5, 41, 0, 0, 110, 101, 1,
		0, 0, 0, 110, 102, 1, 0, 0, 0, 110, 105, 1, 0, 0, 0, 110, 109, 1, 0, 0,
		0, 111, 17, 1, 0, 0, 0, 112, 113, 7, 0, 0, 0, 113, 19, 1, 0, 0, 0, 114,
		115, 5, 8, 0, 0, 115, 120, 5, 9, 0, 0, 116, 117, 5, 8, 0, 0, 117, 120,
		5, 10, 0, 0, 118, 120, 5, 8, 0, 0, 119, 114, 1, 0, 0, 0, 119, 116, 1, 0,
		0, 0, 119, 118, 1, 0, 0, 0, 120, 21, 1, 0, 0, 0, 121, 123, 5, 11, 0, 0,
		122, 124, 5, 19, 0, 0, 123, 122, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124,
		125, 1, 0, 0, 0, 125, 130, 3, 16, 8, 0, 126, 127, 5, 37, 0, 0, 127, 129,
		3, 16, 8, 0, 128, 126, 1, 0, 0, 0, 129, 132, 1, 0, 0, 0, 130, 128, 1, 0,
		0, 0, 130, 131, 1, 0, 0, 0, 131, 187, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0,
		133, 187, 5, 12, 0, 0, 134, 135, 5, 13, 0, 0, 135, 187, 5, 41, 0, 0, 136,
		138, 5, 14, 0, 0, 137, 139, 5, 6, 0, 0, 138, 137, 1, 0, 0, 0, 138, 139,
		1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 187, 3, 16, 8, 0, 141, 143, 5, 15,
		0, 0, 142, 144, 5, 5, 0, 0, 143, 142, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0,
		144, 145, 1, 0, 0, 0, 145, 187, 3, 16, 8, 0, 146, 148, 5, 20, 0, 0, 147,
		146, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 151,
		5, 16, 0, 0, 150, 152, 5, 22, 0, 0, 151, 150, 1, 0, 0, 0, 151, 152, 1,
		0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 187, 5, 41, 0, 0, 154, 156, 5, 20,
		0, 0, 155, 154, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0,
		157, 158, 5, 17, 0, 0, 158, 159, 5, 21, 0, 0, 159, 164, 3, 16, 8, 0, 160,
		161, 5, 37, 0, 0, 161, 163, 3, 16, 8, 0, 162, 160, 1, 0, 0, 0, 163, 166,
		1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 187, 1, 0,
		0, 0, 166, 164, 1, 0, 0, 0, 167, 169, 5, 20, 0, 0, 168, 167, 1, 0, 0, 0,
		168, 169, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 187, 5, 18, 0, 0, 171,
		173, 5, 20, 0, 0, 172, 171, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 174,
		1, 0, 0, 0, 174, 175, 5, 26, 0, 0, 175, 187, 3, 16, 8, 0, 176, 177, 5,
		23, 0, 0, 177, 187, 5, 41, 0, 0, 178, 180, 5, 20, 0, 0, 179, 178, 1, 0,
		0, 0, 179, 180, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 187, 3, 24, 12,
		0, 182, 184, 5, 20, 0, 0, 183, 182, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184,
		185, 1, 0, 0, 0, 185, 187, 3, 26, 13, 0, 186, 121, 1, 0, 0, 0, 186, 133,
		1, 0, 0, 0, 186, 134, 1, 0, 0, 0, 186, 136, 1, 0, 0, 0, 186, 141, 1, 0,
		0, 0, 186, 147, 1, 0, 0, 0, 186, 155, 1, 0, 0, 0, 186, 168, 1, 0, 0, 0,
		186, 172, 1, 0, 0, 0, 186, 176, 1, 0, 0, 0, 186, 179, 1, 0, 0, 0, 186,
		183, 1, 0, 0, 0, 187, 23, 1, 0, 0, 0, 188, 189, 7, 2, 0, 0, 189, 25, 1,
		0, 0, 0, 190, 191, 7, 3, 0, 0, 191, 27, 1, 0, 0, 0, 192, 193, 5, 24, 0,
		0, 193, 194, 5, 5, 0, 0, 194, 206, 5, 41, 0, 0, 195, 196, 5, 24, 0, 0,
		196, 197, 5, 5, 0, 0, 197, 198, 5, 25, 0, 0, 198, 199, 5, 6, 0, 0, 199,
		206, 5, 41, 0, 0, 200, 201, 5, 24, 0, 0, 201, 202, 5, 4, 0, 0, 202, 206,
		5, 41, 0, 0, 203, 204, 5, 24, 0, 0, 204, 206, 5, 41, 0, 0, 205, 192, 1,
		0, 0, 0, 205, 195, 1, 0, 0, 0, 205, 200, 1, 0, 0, 0, 205, 203, 1, 0, 0,
		0, 206, 29, 1, 0, 0, 0, 207, 208, 5, 33, 0, 0, 208, 209, 3, 32, 16, 0,
		209, 31, 1, 0, 0, 0, 210, 211, 7, 4, 0, 0, 211, 33, 1, 0, 0, 0, 21, 37,
		39, 59, 69, 99, 110, 119, 123, 130, 138, 143, 147, 151, 155, 164, 168,
		172, 179, 183, 186, 205,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ADEParserInit initializes any static state used to implement ADEParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewADEParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ADEParserInit() {
	staticData := &ADEParserStaticData
	staticData.once.Do(adeParserInit)
}

// NewADEParser produces a new parser instance for the optional input antlr.TokenStream.
func NewADEParser(input antlr.TokenStream) *ADEParser {
	ADEParserInit()
	this := new(ADEParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ADEParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "ADE.g4"

	return this
}

// ADEParser tokens.
const (
	ADEParserEOF          = antlr.TokenEOF
	ADEParserADR          = 1
	ADEParserFILE         = 2
	ADEParserCODE         = 3
	ADEParserCOMPONENT    = 4
	ADEParserCLASS        = 5
	ADEParserINTERFACE    = 6
	ADEParserPATH         = 7
	ADEParserMUST         = 8
	ADEParserNOT          = 9
	ADEParserONLY         = 10
	ADEParserDEPEND       = 11
	ADEParserEXIST        = 12
	ADEParserCONTAIN      = 13
	ADEParserIMPLEMENT    = 14
	ADEParserEXTEND       = 15
	ADEParserANNOTATED    = 16
	ADEParserACCESSED     = 17
	ADEParserACYCLIC      = 18
	ADEParserON           = 19
	ADEParserBE           = 20
	ADEParserBY           = 21
	ADEParserWITH         = 22
	ADEParserMATCH        = 23
	ADEParserEXCLUDE      = 24
	ADEParserIMPLEMENTING = 25
	ADEParserIN           = 26
	ADEParserPUBLIC       = 27
	ADEParserINTERNAL     = 28
	ADEParserPRIVATE      = 29
	ADEParserABSTRACT     = 30
	ADEParserSEALED       = 31
	ADEParserSTATIC       = 32
	ADEParserSEVERITY     = 33
	ADEParserERROR        = 34
	ADEParserWARNING      = 35
	ADEParserEQ           = 36
	ADEParserCOMMA        = 37
	ADEParserLBRACE       = 38
	ADEParserRBRACE       = 39
	ADEParserIDENTIFIER   = 40
	ADEParserSTRING       = 41
	ADEParserCOMMENT      = 42
	ADEParserWS           = 43
)

// ADEParser rules.
const (
	ADEParserRULE_file           = 0
	ADEParserRULE_adrDecl        = 1
	ADEParserRULE_selectorDecl   = 2
	ADEParserRULE_ruleDecl       = 3
	ADEParserRULE_ruleType       = 4
	ADEParserRULE_ruleStmt       = 5
	ADEParserRULE_assertionStmt  = 6
	ADEParserRULE_subjectExpr    = 7
	ADEParserRULE_targetExpr     = 8
	ADEParserRULE_selectorType   = 9
	ADEParserRULE_mustExpr       = 10
	ADEParserRULE_verbPhrase     = 11
	ADEParserRULE_visibility     = 12
	ADEParserRULE_typeConstraint = 13
	ADEParserRULE_excludeStmt    = 14
	ADEParserRULE_severityStmt   = 15
	ADEParserRULE_severityValue  = 16
)

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AdrDecl() IAdrDeclContext
	EOF() antlr.TerminalNode
	AllSelectorDecl() []ISelectorDeclContext
	SelectorDecl(i int) ISelectorDeclContext
	AllRuleDecl() []IRuleDeclContext
	RuleDecl(i int) IRuleDeclContext

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_file
	return p
}

func InitEmptyFileContext(p *FileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_file
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADEParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) AdrDecl() IAdrDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdrDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdrDeclContext)
}

func (s *FileContext) EOF() antlr.TerminalNode {
	return s.GetToken(ADEParserEOF, 0)
}

func (s *FileContext) AllSelectorDecl() []ISelectorDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISelectorDeclContext); ok {
			len++
		}
	}

	tst := make([]ISelectorDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISelectorDeclContext); ok {
			tst[i] = t.(ISelectorDeclContext)
			i++
		}
	}

	return tst
}

func (s *FileContext) SelectorDecl(i int) ISelectorDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorDeclContext)
}

func (s *FileContext) AllRuleDecl() []IRuleDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRuleDeclContext); ok {
			len++
		}
	}

	tst := make([]IRuleDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRuleDeclContext); ok {
			tst[i] = t.(IRuleDeclContext)
			i++
		}
	}

	return tst
}

func (s *FileContext) RuleDecl(i int) IRuleDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRuleDeclContext)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitFile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADEParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ADEParserRULE_file)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)
		p.AdrDecl()
	}
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&252) != 0 {
		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case ADEParserCOMPONENT, ADEParserCLASS, ADEParserINTERFACE, ADEParserPATH:
			{
				p.SetState(35)
				p.SelectorDecl()
			}

		case ADEParserFILE, ADEParserCODE:
			{
				p.SetState(36)
				p.RuleDecl()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(42)
		p.Match(ADEParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAdrDeclContext is an interface to support dynamic dispatch.
type IAdrDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ADR() antlr.TerminalNode
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode

	// IsAdrDeclContext differentiates from other interfaces.
	IsAdrDeclContext()
}

type AdrDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdrDeclContext() *AdrDeclContext {
	var p = new(AdrDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_adrDecl
	return p
}

func InitEmptyAdrDeclContext(p *AdrDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_adrDecl
}

func (*AdrDeclContext) IsAdrDeclContext() {}

func NewAdrDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdrDeclContext {
	var p = new(AdrDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADEParserRULE_adrDecl

	return p
}

func (s *AdrDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *AdrDeclContext) ADR() antlr.TerminalNode {
	return s.GetToken(ADEParserADR, 0)
}

func (s *AdrDeclContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(ADEParserSTRING)
}

func (s *AdrDeclContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(ADEParserSTRING, i)
}

func (s *AdrDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdrDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdrDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitAdrDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADEParser) AdrDecl() (localctx IAdrDeclContext) {
	localctx = NewAdrDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ADEParserRULE_adrDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Match(ADEParserADR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(45)
		p.Match(ADEParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(46)
		p.Match(ADEParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelectorDeclContext is an interface to support dynamic dispatch.
type ISelectorDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode
	EQ() antlr.TerminalNode
	COMPONENT() antlr.TerminalNode
	CLASS() antlr.TerminalNode
	INTERFACE() antlr.TerminalNode
	PATH() antlr.TerminalNode

	// IsSelectorDeclContext differentiates from other interfaces.
	IsSelectorDeclContext()
}

type SelectorDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorDeclContext() *SelectorDeclContext {
	var p = new(SelectorDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_selectorDecl
	return p
}

func InitEmptySelectorDeclContext(p *SelectorDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_selectorDecl
}

func (*SelectorDeclContext) IsSelectorDeclContext() {}

func NewSelectorDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorDeclContext {
	var p = new(SelectorDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADEParserRULE_selectorDecl

	return p
}

func (s *SelectorDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorDeclContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(ADEParserSTRING)
}

func (s *SelectorDeclContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(ADEParserSTRING, i)
}

func (s *SelectorDeclContext) EQ() antlr.TerminalNode {
	return s.GetToken(ADEParserEQ, 0)
}

func (s *SelectorDeclContext) COMPONENT() antlr.TerminalNode {
	return s.GetToken(ADEParserCOMPONENT, 0)
}

func (s *SelectorDeclContext) CLASS() antlr.TerminalNode {
	return s.GetToken(ADEParserCLASS, 0)
}

func (s *SelectorDeclContext) INTERFACE() antlr.TerminalNode {
	return s.GetToken(ADEParserINTERFACE, 0)
}

func (s *SelectorDeclContext) PATH() antlr.TerminalNode {
	return s.GetToken(ADEParserPATH, 0)
}

func (s *SelectorDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitSelectorDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADEParser) SelectorDecl() (localctx ISelectorDeclContext) {
	localctx = NewSelectorDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ADEParserRULE_selectorDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&240) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(49)
		p.Match(ADEParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
		p.Match(ADEParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(51)
		p.Match(ADEParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRuleDeclContext is an interface to support dynamic dispatch.
type IRuleDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RuleType() IRuleTypeContext
	STRING() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllRuleStmt() []IRuleStmtContext
	RuleStmt(i int) IRuleStmtContext

	// IsRuleDeclContext differentiates from other interfaces.
	IsRuleDeclContext()
}

type RuleDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleDeclContext() *RuleDeclContext {
	var p = new(RuleDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_ruleDecl
	return p
}

func InitEmptyRuleDeclContext(p *RuleDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_ruleDecl
}

func (*RuleDeclContext) IsRuleDeclContext() {}

func NewRuleDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleDeclContext {
	var p = new(RuleDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADEParserRULE_ruleDecl

	return p
}

func (s *RuleDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleDeclContext) RuleType() IRuleTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRuleTypeContext)
}

func (s *RuleDeclContext) STRING() antlr.TerminalNode {
	return s.GetToken(ADEParserSTRING, 0)
}

func (s *RuleDeclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ADEParserLBRACE, 0)
}

func (s *RuleDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ADEParserRBRACE, 0)
}

func (s *RuleDeclContext) AllRuleStmt() []IRuleStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRuleStmtContext); ok {
			len++
		}
	}

	tst := make([]IRuleStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRuleStmtContext); ok {
			tst[i] = t.(IRuleStmtContext)
			i++
		}
	}

	return tst
}

func (s *RuleDeclContext) RuleStmt(i int) IRuleStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRuleStmtContext)
}

func (s *RuleDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitRuleDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADEParser) RuleDecl() (localctx IRuleDeclContext) {
	localctx = NewRuleDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ADEParserRULE_ruleDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.RuleType()
	}
	{
		p.SetState(54)
		p.Match(ADEParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(55)
		p.Match(ADEParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1108118339824) != 0 {
		{
			p.SetState(56)
			p.RuleStmt()
		}

		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(62)
		p.Match(ADEParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRuleTypeContext is an interface to support dynamic dispatch.
type IRuleTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FILE() antlr.TerminalNode
	CODE() antlr.TerminalNode

	// IsRuleTypeContext differentiates from other interfaces.
	IsRuleTypeContext()
}

type RuleTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleTypeContext() *RuleTypeContext {
	var p = new(RuleTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_ruleType
	return p
}

func InitEmptyRuleTypeContext(p *RuleTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_ruleType
}

func (*RuleTypeContext) IsRuleTypeContext() {}

func NewRuleTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleTypeContext {
	var p = new(RuleTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADEParserRULE_ruleType

	return p
}

func (s *RuleTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleTypeContext) FILE() antlr.TerminalNode {
	return s.GetToken(ADEParserFILE, 0)
}

func (s *RuleTypeContext) CODE() antlr.TerminalNode {
	return s.GetToken(ADEParserCODE, 0)
}

func (s *RuleTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitRuleType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADEParser) RuleType() (localctx IRuleTypeContext) {
	localctx = NewRuleTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ADEParserRULE_ruleType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ADEParserFILE || _la == ADEParserCODE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRuleStmtContext is an interface to support dynamic dispatch.
type IRuleStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AssertionStmt() IAssertionStmtContext
	ExcludeStmt() IExcludeStmtContext
	SeverityStmt() ISeverityStmtContext

	// IsRuleStmtContext differentiates from other interfaces.
	IsRuleStmtContext()
}

type RuleStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleStmtContext() *RuleStmtContext {
	var p = new(RuleStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_ruleStmt
	return p
}

func InitEmptyRuleStmtContext(p *RuleStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_ruleStmt
}

func (*RuleStmtContext) IsRuleStmtContext() {}

func NewRuleStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleStmtContext {
	var p = new(RuleStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADEParserRULE_ruleStmt

	return p
}

func (s *RuleStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleStmtContext) AssertionStmt() IAssertionStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssertionStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssertionStmtContext)
}

func (s *RuleStmtContext) ExcludeStmt() IExcludeStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExcludeStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExcludeStmtContext)
}

func (s *RuleStmtContext) SeverityStmt() ISeverityStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISeverityStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISeverityStmtContext)
}

func (s *RuleStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitRuleStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADEParser) RuleStmt() (localctx IRuleStmtContext) {
	localctx = NewRuleStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ADEParserRULE_ruleStmt)
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ADEParserCOMPONENT, ADEParserCLASS, ADEParserINTERFACE, ADEParserPATH, ADEParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(66)
			p.AssertionStmt()
		}

	case ADEParserEXCLUDE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(67)
			p.ExcludeStmt()
		}

	case ADEParserSEVERITY:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(68)
			p.SeverityStmt()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssertionStmtContext is an interface to support dynamic dispatch.
type IAssertionStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SubjectExpr() ISubjectExprContext
	MustExpr() IMustExprContext
	VerbPhrase() IVerbPhraseContext

	// IsAssertionStmtContext differentiates from other interfaces.
	IsAssertionStmtContext()
}

type AssertionStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssertionStmtContext() *AssertionStmtContext {
	var p = new(AssertionStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_assertionStmt
	return p
}

func InitEmptyAssertionStmtContext(p *AssertionStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_assertionStmt
}

func (*AssertionStmtContext) IsAssertionStmtContext() {}

func NewAssertionStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssertionStmtContext {
	var p = new(AssertionStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADEParserRULE_assertionStmt

	return p
}

func (s *AssertionStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AssertionStmtContext) SubjectExpr() ISubjectExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubjectExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubjectExprContext)
}

func (s *AssertionStmtContext) MustExpr() IMustExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMustExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMustExprContext)
}

func (s *AssertionStmtContext) VerbPhrase() IVerbPhraseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVerbPhraseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVerbPhraseContext)
}

func (s *AssertionStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssertionStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssertionStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitAssertionStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADEParser) AssertionStmt() (localctx IAssertionStmtContext) {
	localctx = NewAssertionStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ADEParserRULE_assertionStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.SubjectExpr()
	}
	{
		p.SetState(72)
		p.MustExpr()
	}
	{
		p.SetState(73)
		p.VerbPhrase()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISubjectExprContext is an interface to support dynamic dispatch.
type ISubjectExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSubjectExprContext differentiates from other interfaces.
	IsSubjectExprContext()
}

type SubjectExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubjectExprContext() *SubjectExprContext {
	var p = new(SubjectExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_subjectExpr
	return p
}

func InitEmptySubjectExprContext(p *SubjectExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_subjectExpr
}

func (*SubjectExprContext) IsSubjectExprContext() {}

func NewSubjectExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubjectExprContext {
	var p = new(SubjectExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADEParserRULE_subjectExpr

	return p
}

func (s *SubjectExprContext) GetParser() antlr.Parser { return s.parser }

func (s *SubjectExprContext) CopyAll(ctx *SubjectExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SubjectExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubjectExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SubsetMatchContext struct {
	SubjectExprContext
}

func NewSubsetMatchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubsetMatchContext {
	var p = new(SubsetMatchContext)

	InitEmptySubjectExprContext(&p.SubjectExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*SubjectExprContext))

	return p
}

func (s *SubsetMatchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubsetMatchContext) SelectorType() ISelectorTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorTypeContext)
}

func (s *SubsetMatchContext) MATCH() antlr.TerminalNode {
	return s.GetToken(ADEParserMATCH, 0)
}

func (s *SubsetMatchContext) STRING() antlr.TerminalNode {
	return s.GetToken(ADEParserSTRING, 0)
}

func (s *SubsetMatchContext) IN() antlr.TerminalNode {
	return s.GetToken(ADEParserIN, 0)
}

func (s *SubsetMatchContext) TargetExpr() ITargetExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetExprContext)
}

func (s *SubsetMatchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitSubsetMatch(s)

	default:
		return t.VisitChildren(s)
	}
}

type SubsetAllContext struct {
	SubjectExprContext
}

func NewSubsetAllContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubsetAllContext {
	var p = new(SubsetAllContext)

	InitEmptySubjectExprContext(&p.SubjectExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*SubjectExprContext))

	return p
}

func (s *SubsetAllContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubsetAllContext) SelectorType() ISelectorTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorTypeContext)
}

func (s *SubsetAllContext) IN() antlr.TerminalNode {
	return s.GetToken(ADEParserIN, 0)
}

func (s *SubsetAllContext) TargetExpr() ITargetExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetExprContext)
}

func (s *SubsetAllContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitSubsetAll(s)

	default:
		return t.VisitChildren(s)
	}
}

type InlineTypeContext struct {
	SubjectExprContext
}

func NewInlineTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InlineTypeContext {
	var p = new(InlineTypeContext)

	InitEmptySubjectExprContext(&p.SubjectExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*SubjectExprContext))

	return p
}

func (s *InlineTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineTypeContext) SelectorType() ISelectorTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorTypeContext)
}

func (s *InlineTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitInlineType(s)

	default:
		return t.VisitChildren(s)
	}
}

type SelectorRefContext struct {
	SubjectExprContext
}

func NewSelectorRefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectorRefContext {
	var p = new(SelectorRefContext)

	InitEmptySubjectExprContext(&p.SubjectExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*SubjectExprContext))

	return p
}

func (s *SelectorRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorRefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ADEParserIDENTIFIER, 0)
}

func (s *SelectorRefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitSelectorRef(s)

	default:
		return t.VisitChildren(s)
	}
}

type InlineLiteralContext struct {
	SubjectExprContext
}

func NewInlineLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InlineLiteralContext {
	var p = new(InlineLiteralContext)

	InitEmptySubjectExprContext(&p.SubjectExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*SubjectExprContext))

	return p
}

func (s *InlineLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineLiteralContext) SelectorType() ISelectorTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorTypeContext)
}

func (s *InlineLiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(ADEParserSTRING, 0)
}

func (s *InlineLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitInlineLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type SubsetLiteralContext struct {
	SubjectExprContext
}

func NewSubsetLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubsetLiteralContext {
	var p = new(SubsetLiteralContext)

	InitEmptySubjectExprContext(&p.SubjectExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*SubjectExprContext))

	return p
}

func (s *SubsetLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubsetLiteralContext) SelectorType() ISelectorTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorTypeContext)
}

func (s *SubsetLiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(ADEParserSTRING, 0)
}

func (s *SubsetLiteralContext) IN() antlr.TerminalNode {
	return s.GetToken(ADEParserIN, 0)
}

func (s *SubsetLiteralContext) TargetExpr() ITargetExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetExprContext)
}

func (s *SubsetLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitSubsetLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type InlineMatchContext struct {
	SubjectExprContext
}

func NewInlineMatchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InlineMatchContext {
	var p = new(InlineMatchContext)

	InitEmptySubjectExprContext(&p.SubjectExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*SubjectExprContext))

	return p
}

func (s *InlineMatchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineMatchContext) SelectorType() ISelectorTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorTypeContext)
}

func (s *InlineMatchContext) MATCH() antlr.TerminalNode {
	return s.GetToken(ADEParserMATCH, 0)
}

func (s *InlineMatchContext) STRING() antlr.TerminalNode {
	return s.GetToken(ADEParserSTRING, 0)
}

func (s *InlineMatchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitInlineMatch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADEParser) SubjectExpr() (localctx ISubjectExprContext) {
	localctx = NewSubjectExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ADEParserRULE_subjectExpr)
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSelectorRefContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(75)
			p.Match(ADEParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewSubsetLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(76)
			p.SelectorType()
		}
		{
			p.SetState(77)
			p.Match(ADEParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)
			p.Match(ADEParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)
			p.TargetExpr()
		}

	case 3:
		localctx = NewSubsetMatchContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(81)
			p.SelectorType()
		}
		{
			p.SetState(82)
			p.Match(ADEParserMATCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.Match(ADEParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(84)
			p.Match(ADEParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
			p.TargetExpr()
		}

	case 4:
		localctx = NewSubsetAllContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(87)
			p.SelectorType()
		}
		{
			p.SetState(88)
			p.Match(ADEParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(89)
			p.TargetExpr()
		}

	case 5:
		localctx = NewInlineLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(91)
			p.SelectorType()
		}
		{
			p.SetState(92)
			p.Match(ADEParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewInlineMatchContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(94)
			p.SelectorType()
		}
		{
			p.SetState(95)
			p.Match(ADEParserMATCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(96)
			p.Match(ADEParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewInlineTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(98)
			p.SelectorType()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITargetExprContext is an interface to support dynamic dispatch.
type ITargetExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTargetExprContext differentiates from other interfaces.
	IsTargetExprContext()
}

type TargetExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetExprContext() *TargetExprContext {
	var p = new(TargetExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_targetExpr
	return p
}

func InitEmptyTargetExprContext(p *TargetExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_targetExpr
}

func (*TargetExprContext) IsTargetExprContext() {}

func NewTargetExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetExprContext {
	var p = new(TargetExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADEParserRULE_targetExpr

	return p
}

func (s *TargetExprContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetExprContext) CopyAll(ctx *TargetExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TargetExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TargetStringLiteralContext struct {
	TargetExprContext
}

func NewTargetStringLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TargetStringLiteralContext {
	var p = new(TargetStringLiteralContext)

	InitEmptyTargetExprContext(&p.TargetExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*TargetExprContext))

	return p
}

func (s *TargetStringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetStringLiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(ADEParserSTRING, 0)
}

func (s *TargetStringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitTargetStringLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type TargetSelectorRefContext struct {
	TargetExprContext
}

func NewTargetSelectorRefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TargetSelectorRefContext {
	var p = new(TargetSelectorRefContext)

	InitEmptyTargetExprContext(&p.TargetExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*TargetExprContext))

	return p
}

func (s *TargetSelectorRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetSelectorRefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ADEParserIDENTIFIER, 0)
}

func (s *TargetSelectorRefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitTargetSelectorRef(s)

	default:
		return t.VisitChildren(s)
	}
}

type TargetInlineLiteralContext struct {
	TargetExprContext
}

func NewTargetInlineLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TargetInlineLiteralContext {
	var p = new(TargetInlineLiteralContext)

	InitEmptyTargetExprContext(&p.TargetExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*TargetExprContext))

	return p
}

func (s *TargetInlineLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetInlineLiteralContext) SelectorType() ISelectorTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorTypeContext)
}

func (s *TargetInlineLiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(ADEParserSTRING, 0)
}

func (s *TargetInlineLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitTargetInlineLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type TargetInlineMatchContext struct {
	TargetExprContext
}

func NewTargetInlineMatchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TargetInlineMatchContext {
	var p = new(TargetInlineMatchContext)

	InitEmptyTargetExprContext(&p.TargetExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*TargetExprContext))

	return p
}

func (s *TargetInlineMatchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetInlineMatchContext) SelectorType() ISelectorTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorTypeContext)
}

func (s *TargetInlineMatchContext) MATCH() antlr.TerminalNode {
	return s.GetToken(ADEParserMATCH, 0)
}

func (s *TargetInlineMatchContext) STRING() antlr.TerminalNode {
	return s.GetToken(ADEParserSTRING, 0)
}

func (s *TargetInlineMatchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitTargetInlineMatch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADEParser) TargetExpr() (localctx ITargetExprContext) {
	localctx = NewTargetExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ADEParserRULE_targetExpr)
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTargetSelectorRefContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(101)
			p.Match(ADEParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewTargetInlineLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(102)
			p.SelectorType()
		}
		{
			p.SetState(103)
			p.Match(ADEParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewTargetInlineMatchContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(105)
			p.SelectorType()
		}
		{
			p.SetState(106)
			p.Match(ADEParserMATCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
			p.Match(ADEParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewTargetStringLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(109)
			p.Match(ADEParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelectorTypeContext is an interface to support dynamic dispatch.
type ISelectorTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMPONENT() antlr.TerminalNode
	CLASS() antlr.TerminalNode
	INTERFACE() antlr.TerminalNode
	PATH() antlr.TerminalNode

	// IsSelectorTypeContext differentiates from other interfaces.
	IsSelectorTypeContext()
}

type SelectorTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorTypeContext() *SelectorTypeContext {
	var p = new(SelectorTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_selectorType
	return p
}

func InitEmptySelectorTypeContext(p *SelectorTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_selectorType
}

func (*SelectorTypeContext) IsSelectorTypeContext() {}

func NewSelectorTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorTypeContext {
	var p = new(SelectorTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADEParserRULE_selectorType

	return p
}

func (s *SelectorTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorTypeContext) COMPONENT() antlr.TerminalNode {
	return s.GetToken(ADEParserCOMPONENT, 0)
}

func (s *SelectorTypeContext) CLASS() antlr.TerminalNode {
	return s.GetToken(ADEParserCLASS, 0)
}

func (s *SelectorTypeContext) INTERFACE() antlr.TerminalNode {
	return s.GetToken(ADEParserINTERFACE, 0)
}

func (s *SelectorTypeContext) PATH() antlr.TerminalNode {
	return s.GetToken(ADEParserPATH, 0)
}

func (s *SelectorTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitSelectorType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADEParser) SelectorType() (localctx ISelectorTypeContext) {
	localctx = NewSelectorTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ADEParserRULE_selectorType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&240) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMustExprContext is an interface to support dynamic dispatch.
type IMustExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MUST() antlr.TerminalNode
	NOT() antlr.TerminalNode
	ONLY() antlr.TerminalNode

	// IsMustExprContext differentiates from other interfaces.
	IsMustExprContext()
}

type MustExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMustExprContext() *MustExprContext {
	var p = new(MustExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_mustExpr
	return p
}

func InitEmptyMustExprContext(p *MustExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_mustExpr
}

func (*MustExprContext) IsMustExprContext() {}

func NewMustExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MustExprContext {
	var p = new(MustExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADEParserRULE_mustExpr

	return p
}

func (s *MustExprContext) GetParser() antlr.Parser { return s.parser }

func (s *MustExprContext) MUST() antlr.TerminalNode {
	return s.GetToken(ADEParserMUST, 0)
}

func (s *MustExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(ADEParserNOT, 0)
}

func (s *MustExprContext) ONLY() antlr.TerminalNode {
	return s.GetToken(ADEParserONLY, 0)
}

func (s *MustExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MustExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MustExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitMustExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADEParser) MustExpr() (localctx IMustExprContext) {
	localctx = NewMustExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ADEParserRULE_mustExpr)
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(114)
			p.Match(ADEParserMUST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(115)
			p.Match(ADEParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(116)
			p.Match(ADEParserMUST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)
			p.Match(ADEParserONLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(118)
			p.Match(ADEParserMUST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVerbPhraseContext is an interface to support dynamic dispatch.
type IVerbPhraseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVerbPhraseContext differentiates from other interfaces.
	IsVerbPhraseContext()
}

type VerbPhraseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVerbPhraseContext() *VerbPhraseContext {
	var p = new(VerbPhraseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_verbPhrase
	return p
}

func InitEmptyVerbPhraseContext(p *VerbPhraseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_verbPhrase
}

func (*VerbPhraseContext) IsVerbPhraseContext() {}

func NewVerbPhraseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VerbPhraseContext {
	var p = new(VerbPhraseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADEParserRULE_verbPhrase

	return p
}

func (s *VerbPhraseContext) GetParser() antlr.Parser { return s.parser }

func (s *VerbPhraseContext) CopyAll(ctx *VerbPhraseContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *VerbPhraseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VerbPhraseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DependOnPhraseContext struct {
	VerbPhraseContext
}

func NewDependOnPhraseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DependOnPhraseContext {
	var p = new(DependOnPhraseContext)

	InitEmptyVerbPhraseContext(&p.VerbPhraseContext)
	p.parser = parser
	p.CopyAll(ctx.(*VerbPhraseContext))

	return p
}

func (s *DependOnPhraseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DependOnPhraseContext) DEPEND() antlr.TerminalNode {
	return s.GetToken(ADEParserDEPEND, 0)
}

func (s *DependOnPhraseContext) AllTargetExpr() []ITargetExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITargetExprContext); ok {
			len++
		}
	}

	tst := make([]ITargetExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITargetExprContext); ok {
			tst[i] = t.(ITargetExprContext)
			i++
		}
	}

	return tst
}

func (s *DependOnPhraseContext) TargetExpr(i int) ITargetExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetExprContext)
}

func (s *DependOnPhraseContext) ON() antlr.TerminalNode {
	return s.GetToken(ADEParserON, 0)
}

func (s *DependOnPhraseContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ADEParserCOMMA)
}

func (s *DependOnPhraseContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ADEParserCOMMA, i)
}

func (s *DependOnPhraseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitDependOnPhrase(s)

	default:
		return t.VisitChildren(s)
	}
}

type MatchPhraseContext struct {
	VerbPhraseContext
}

func NewMatchPhraseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MatchPhraseContext {
	var p = new(MatchPhraseContext)

	InitEmptyVerbPhraseContext(&p.VerbPhraseContext)
	p.parser = parser
	p.CopyAll(ctx.(*VerbPhraseContext))

	return p
}

func (s *MatchPhraseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchPhraseContext) MATCH() antlr.TerminalNode {
	return s.GetToken(ADEParserMATCH, 0)
}

func (s *MatchPhraseContext) STRING() antlr.TerminalNode {
	return s.GetToken(ADEParserSTRING, 0)
}

func (s *MatchPhraseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitMatchPhrase(s)

	default:
		return t.VisitChildren(s)
	}
}

type AnnotatedPhraseContext struct {
	VerbPhraseContext
}

func NewAnnotatedPhraseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AnnotatedPhraseContext {
	var p = new(AnnotatedPhraseContext)

	InitEmptyVerbPhraseContext(&p.VerbPhraseContext)
	p.parser = parser
	p.CopyAll(ctx.(*VerbPhraseContext))

	return p
}

func (s *AnnotatedPhraseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotatedPhraseContext) ANNOTATED() antlr.TerminalNode {
	return s.GetToken(ADEParserANNOTATED, 0)
}

func (s *AnnotatedPhraseContext) STRING() antlr.TerminalNode {
	return s.GetToken(ADEParserSTRING, 0)
}

func (s *AnnotatedPhraseContext) BE() antlr.TerminalNode {
	return s.GetToken(ADEParserBE, 0)
}

func (s *AnnotatedPhraseContext) WITH() antlr.TerminalNode {
	return s.GetToken(ADEParserWITH, 0)
}

func (s *AnnotatedPhraseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitAnnotatedPhrase(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeConstraintPhraseContext struct {
	VerbPhraseContext
}

func NewTypeConstraintPhraseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeConstraintPhraseContext {
	var p = new(TypeConstraintPhraseContext)

	InitEmptyVerbPhraseContext(&p.VerbPhraseContext)
	p.parser = parser
	p.CopyAll(ctx.(*VerbPhraseContext))

	return p
}

func (s *TypeConstraintPhraseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeConstraintPhraseContext) TypeConstraint() ITypeConstraintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeConstraintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeConstraintContext)
}

func (s *TypeConstraintPhraseContext) BE() antlr.TerminalNode {
	return s.GetToken(ADEParserBE, 0)
}

func (s *TypeConstraintPhraseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitTypeConstraintPhrase(s)

	default:
		return t.VisitChildren(s)
	}
}

type ImplementPhraseContext struct {
	VerbPhraseContext
}

func NewImplementPhraseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImplementPhraseContext {
	var p = new(ImplementPhraseContext)

	InitEmptyVerbPhraseContext(&p.VerbPhraseContext)
	p.parser = parser
	p.CopyAll(ctx.(*VerbPhraseContext))

	return p
}

func (s *ImplementPhraseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImplementPhraseContext) IMPLEMENT() antlr.TerminalNode {
	return s.GetToken(ADEParserIMPLEMENT, 0)
}

func (s *ImplementPhraseContext) TargetExpr() ITargetExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetExprContext)
}

func (s *ImplementPhraseContext) INTERFACE() antlr.TerminalNode {
	return s.GetToken(ADEParserINTERFACE, 0)
}

func (s *ImplementPhraseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitImplementPhrase(s)

	default:
		return t.VisitChildren(s)
	}
}

type AccessedByPhraseContext struct {
	VerbPhraseContext
}

func NewAccessedByPhraseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AccessedByPhraseContext {
	var p = new(AccessedByPhraseContext)

	InitEmptyVerbPhraseContext(&p.VerbPhraseContext)
	p.parser = parser
	p.CopyAll(ctx.(*VerbPhraseContext))

	return p
}

func (s *AccessedByPhraseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessedByPhraseContext) ACCESSED() antlr.TerminalNode {
	return s.GetToken(ADEParserACCESSED, 0)
}

func (s *AccessedByPhraseContext) BY() antlr.TerminalNode {
	return s.GetToken(ADEParserBY, 0)
}

func (s *AccessedByPhraseContext) AllTargetExpr() []ITargetExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITargetExprContext); ok {
			len++
		}
	}

	tst := make([]ITargetExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITargetExprContext); ok {
			tst[i] = t.(ITargetExprContext)
			i++
		}
	}

	return tst
}

func (s *AccessedByPhraseContext) TargetExpr(i int) ITargetExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetExprContext)
}

func (s *AccessedByPhraseContext) BE() antlr.TerminalNode {
	return s.GetToken(ADEParserBE, 0)
}

func (s *AccessedByPhraseContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ADEParserCOMMA)
}

func (s *AccessedByPhraseContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ADEParserCOMMA, i)
}

func (s *AccessedByPhraseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitAccessedByPhrase(s)

	default:
		return t.VisitChildren(s)
	}
}

type ContainPhraseContext struct {
	VerbPhraseContext
}

func NewContainPhraseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContainPhraseContext {
	var p = new(ContainPhraseContext)

	InitEmptyVerbPhraseContext(&p.VerbPhraseContext)
	p.parser = parser
	p.CopyAll(ctx.(*VerbPhraseContext))

	return p
}

func (s *ContainPhraseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContainPhraseContext) CONTAIN() antlr.TerminalNode {
	return s.GetToken(ADEParserCONTAIN, 0)
}

func (s *ContainPhraseContext) STRING() antlr.TerminalNode {
	return s.GetToken(ADEParserSTRING, 0)
}

func (s *ContainPhraseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitContainPhrase(s)

	default:
		return t.VisitChildren(s)
	}
}

type InPhraseContext struct {
	VerbPhraseContext
}

func NewInPhraseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InPhraseContext {
	var p = new(InPhraseContext)

	InitEmptyVerbPhraseContext(&p.VerbPhraseContext)
	p.parser = parser
	p.CopyAll(ctx.(*VerbPhraseContext))

	return p
}

func (s *InPhraseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InPhraseContext) IN() antlr.TerminalNode {
	return s.GetToken(ADEParserIN, 0)
}

func (s *InPhraseContext) TargetExpr() ITargetExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetExprContext)
}

func (s *InPhraseContext) BE() antlr.TerminalNode {
	return s.GetToken(ADEParserBE, 0)
}

func (s *InPhraseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitInPhrase(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExtendPhraseContext struct {
	VerbPhraseContext
}

func NewExtendPhraseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExtendPhraseContext {
	var p = new(ExtendPhraseContext)

	InitEmptyVerbPhraseContext(&p.VerbPhraseContext)
	p.parser = parser
	p.CopyAll(ctx.(*VerbPhraseContext))

	return p
}

func (s *ExtendPhraseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtendPhraseContext) EXTEND() antlr.TerminalNode {
	return s.GetToken(ADEParserEXTEND, 0)
}

func (s *ExtendPhraseContext) TargetExpr() ITargetExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetExprContext)
}

func (s *ExtendPhraseContext) CLASS() antlr.TerminalNode {
	return s.GetToken(ADEParserCLASS, 0)
}

func (s *ExtendPhraseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitExtendPhrase(s)

	default:
		return t.VisitChildren(s)
	}
}

type VisibilityPhraseContext struct {
	VerbPhraseContext
}

func NewVisibilityPhraseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VisibilityPhraseContext {
	var p = new(VisibilityPhraseContext)

	InitEmptyVerbPhraseContext(&p.VerbPhraseContext)
	p.parser = parser
	p.CopyAll(ctx.(*VerbPhraseContext))

	return p
}

func (s *VisibilityPhraseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VisibilityPhraseContext) Visibility() IVisibilityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVisibilityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVisibilityContext)
}

func (s *VisibilityPhraseContext) BE() antlr.TerminalNode {
	return s.GetToken(ADEParserBE, 0)
}

func (s *VisibilityPhraseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitVisibilityPhrase(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExistPhraseContext struct {
	VerbPhraseContext
}

func NewExistPhraseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExistPhraseContext {
	var p = new(ExistPhraseContext)

	InitEmptyVerbPhraseContext(&p.VerbPhraseContext)
	p.parser = parser
	p.CopyAll(ctx.(*VerbPhraseContext))

	return p
}

func (s *ExistPhraseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExistPhraseContext) EXIST() antlr.TerminalNode {
	return s.GetToken(ADEParserEXIST, 0)
}

func (s *ExistPhraseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitExistPhrase(s)

	default:
		return t.VisitChildren(s)
	}
}

type AcyclicPhraseContext struct {
	VerbPhraseContext
}

func NewAcyclicPhraseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AcyclicPhraseContext {
	var p = new(AcyclicPhraseContext)

	InitEmptyVerbPhraseContext(&p.VerbPhraseContext)
	p.parser = parser
	p.CopyAll(ctx.(*VerbPhraseContext))

	return p
}

func (s *AcyclicPhraseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AcyclicPhraseContext) ACYCLIC() antlr.TerminalNode {
	return s.GetToken(ADEParserACYCLIC, 0)
}

func (s *AcyclicPhraseContext) BE() antlr.TerminalNode {
	return s.GetToken(ADEParserBE, 0)
}

func (s *AcyclicPhraseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitAcyclicPhrase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADEParser) VerbPhrase() (localctx IVerbPhraseContext) {
	localctx = NewVerbPhraseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ADEParserRULE_verbPhrase)
	var _la int

	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDependOnPhraseContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.Match(ADEParserDEPEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ADEParserON {
			{
				p.SetState(122)
				p.Match(ADEParserON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(125)
			p.TargetExpr()
		}
		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ADEParserCOMMA {
			{
				p.SetState(126)
				p.Match(ADEParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(127)
				p.TargetExpr()
			}

			p.SetState(132)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		localctx = NewExistPhraseContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(133)
			p.Match(ADEParserEXIST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewContainPhraseContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(134)
			p.Match(ADEParserCONTAIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.Match(ADEParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewImplementPhraseContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(136)
			p.Match(ADEParserIMPLEMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(138)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(137)
				p.Match(ADEParserINTERFACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(140)
			p.TargetExpr()
		}

	case 5:
		localctx = NewExtendPhraseContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(141)
			p.Match(ADEParserEXTEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(142)
				p.Match(ADEParserCLASS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(145)
			p.TargetExpr()
		}

	case 6:
		localctx = NewAnnotatedPhraseContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ADEParserBE {
			{
				p.SetState(146)
				p.Match(ADEParserBE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(149)
			p.Match(ADEParserANNOTATED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ADEParserWITH {
			{
				p.SetState(150)
				p.Match(ADEParserWITH)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(153)
			p.Match(ADEParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewAccessedByPhraseContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ADEParserBE {
			{
				p.SetState(154)
				p.Match(ADEParserBE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(157)
			p.Match(ADEParserACCESSED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)
			p.Match(ADEParserBY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(159)
			p.TargetExpr()
		}
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ADEParserCOMMA {
			{
				p.SetState(160)
				p.Match(ADEParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(161)
				p.TargetExpr()
			}

			p.SetState(166)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 8:
		localctx = NewAcyclicPhraseContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ADEParserBE {
			{
				p.SetState(167)
				p.Match(ADEParserBE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(170)
			p.Match(ADEParserACYCLIC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewInPhraseContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ADEParserBE {
			{
				p.SetState(171)
				p.Match(ADEParserBE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(174)
			p.Match(ADEParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(175)
			p.TargetExpr()
		}

	case 10:
		localctx = NewMatchPhraseContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(176)
			p.Match(ADEParserMATCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(177)
			p.Match(ADEParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		localctx = NewVisibilityPhraseContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ADEParserBE {
			{
				p.SetState(178)
				p.Match(ADEParserBE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(181)
			p.Visibility()
		}

	case 12:
		localctx = NewTypeConstraintPhraseContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ADEParserBE {
			{
				p.SetState(182)
				p.Match(ADEParserBE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(185)
			p.TypeConstraint()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVisibilityContext is an interface to support dynamic dispatch.
type IVisibilityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PUBLIC() antlr.TerminalNode
	INTERNAL() antlr.TerminalNode
	PRIVATE() antlr.TerminalNode

	// IsVisibilityContext differentiates from other interfaces.
	IsVisibilityContext()
}

type VisibilityContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVisibilityContext() *VisibilityContext {
	var p = new(VisibilityContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_visibility
	return p
}

func InitEmptyVisibilityContext(p *VisibilityContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_visibility
}

func (*VisibilityContext) IsVisibilityContext() {}

func NewVisibilityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VisibilityContext {
	var p = new(VisibilityContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADEParserRULE_visibility

	return p
}

func (s *VisibilityContext) GetParser() antlr.Parser { return s.parser }

func (s *VisibilityContext) PUBLIC() antlr.TerminalNode {
	return s.GetToken(ADEParserPUBLIC, 0)
}

func (s *VisibilityContext) INTERNAL() antlr.TerminalNode {
	return s.GetToken(ADEParserINTERNAL, 0)
}

func (s *VisibilityContext) PRIVATE() antlr.TerminalNode {
	return s.GetToken(ADEParserPRIVATE, 0)
}

func (s *VisibilityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VisibilityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VisibilityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitVisibility(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADEParser) Visibility() (localctx IVisibilityContext) {
	localctx = NewVisibilityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ADEParserRULE_visibility)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&939524096) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeConstraintContext is an interface to support dynamic dispatch.
type ITypeConstraintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ABSTRACT() antlr.TerminalNode
	SEALED() antlr.TerminalNode
	STATIC() antlr.TerminalNode

	// IsTypeConstraintContext differentiates from other interfaces.
	IsTypeConstraintContext()
}

type TypeConstraintContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeConstraintContext() *TypeConstraintContext {
	var p = new(TypeConstraintContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_typeConstraint
	return p
}

func InitEmptyTypeConstraintContext(p *TypeConstraintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_typeConstraint
}

func (*TypeConstraintContext) IsTypeConstraintContext() {}

func NewTypeConstraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeConstraintContext {
	var p = new(TypeConstraintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADEParserRULE_typeConstraint

	return p
}

func (s *TypeConstraintContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeConstraintContext) ABSTRACT() antlr.TerminalNode {
	return s.GetToken(ADEParserABSTRACT, 0)
}

func (s *TypeConstraintContext) SEALED() antlr.TerminalNode {
	return s.GetToken(ADEParserSEALED, 0)
}

func (s *TypeConstraintContext) STATIC() antlr.TerminalNode {
	return s.GetToken(ADEParserSTATIC, 0)
}

func (s *TypeConstraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeConstraintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeConstraintContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitTypeConstraint(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADEParser) TypeConstraint() (localctx ITypeConstraintContext) {
	localctx = NewTypeConstraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ADEParserRULE_typeConstraint)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7516192768) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExcludeStmtContext is an interface to support dynamic dispatch.
type IExcludeStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExcludeStmtContext differentiates from other interfaces.
	IsExcludeStmtContext()
}

type ExcludeStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExcludeStmtContext() *ExcludeStmtContext {
	var p = new(ExcludeStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_excludeStmt
	return p
}

func InitEmptyExcludeStmtContext(p *ExcludeStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_excludeStmt
}

func (*ExcludeStmtContext) IsExcludeStmtContext() {}

func NewExcludeStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExcludeStmtContext {
	var p = new(ExcludeStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADEParserRULE_excludeStmt

	return p
}

func (s *ExcludeStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ExcludeStmtContext) CopyAll(ctx *ExcludeStmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExcludeStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExcludeStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExcludeClassContext struct {
	ExcludeStmtContext
}

func NewExcludeClassContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExcludeClassContext {
	var p = new(ExcludeClassContext)

	InitEmptyExcludeStmtContext(&p.ExcludeStmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExcludeStmtContext))

	return p
}

func (s *ExcludeClassContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExcludeClassContext) EXCLUDE() antlr.TerminalNode {
	return s.GetToken(ADEParserEXCLUDE, 0)
}

func (s *ExcludeClassContext) CLASS() antlr.TerminalNode {
	return s.GetToken(ADEParserCLASS, 0)
}

func (s *ExcludeClassContext) STRING() antlr.TerminalNode {
	return s.GetToken(ADEParserSTRING, 0)
}

func (s *ExcludeClassContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitExcludeClass(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExcludePatternContext struct {
	ExcludeStmtContext
}

func NewExcludePatternContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExcludePatternContext {
	var p = new(ExcludePatternContext)

	InitEmptyExcludeStmtContext(&p.ExcludeStmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExcludeStmtContext))

	return p
}

func (s *ExcludePatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExcludePatternContext) EXCLUDE() antlr.TerminalNode {
	return s.GetToken(ADEParserEXCLUDE, 0)
}

func (s *ExcludePatternContext) STRING() antlr.TerminalNode {
	return s.GetToken(ADEParserSTRING, 0)
}

func (s *ExcludePatternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitExcludePattern(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExcludeComponentContext struct {
	ExcludeStmtContext
}

func NewExcludeComponentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExcludeComponentContext {
	var p = new(ExcludeComponentContext)

	InitEmptyExcludeStmtContext(&p.ExcludeStmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExcludeStmtContext))

	return p
}

func (s *ExcludeComponentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExcludeComponentContext) EXCLUDE() antlr.TerminalNode {
	return s.GetToken(ADEParserEXCLUDE, 0)
}

func (s *ExcludeComponentContext) COMPONENT() antlr.TerminalNode {
	return s.GetToken(ADEParserCOMPONENT, 0)
}

func (s *ExcludeComponentContext) STRING() antlr.TerminalNode {
	return s.GetToken(ADEParserSTRING, 0)
}

func (s *ExcludeComponentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitExcludeComponent(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExcludeClassImplementingContext struct {
	ExcludeStmtContext
}

func NewExcludeClassImplementingContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExcludeClassImplementingContext {
	var p = new(ExcludeClassImplementingContext)

	InitEmptyExcludeStmtContext(&p.ExcludeStmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExcludeStmtContext))

	return p
}

func (s *ExcludeClassImplementingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExcludeClassImplementingContext) EXCLUDE() antlr.TerminalNode {
	return s.GetToken(ADEParserEXCLUDE, 0)
}

func (s *ExcludeClassImplementingContext) CLASS() antlr.TerminalNode {
	return s.GetToken(ADEParserCLASS, 0)
}

func (s *ExcludeClassImplementingContext) IMPLEMENTING() antlr.TerminalNode {
	return s.GetToken(ADEParserIMPLEMENTING, 0)
}

func (s *ExcludeClassImplementingContext) INTERFACE() antlr.TerminalNode {
	return s.GetToken(ADEParserINTERFACE, 0)
}

func (s *ExcludeClassImplementingContext) STRING() antlr.TerminalNode {
	return s.GetToken(ADEParserSTRING, 0)
}

func (s *ExcludeClassImplementingContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitExcludeClassImplementing(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADEParser) ExcludeStmt() (localctx IExcludeStmtContext) {
	localctx = NewExcludeStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ADEParserRULE_excludeStmt)
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExcludeClassContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(192)
			p.Match(ADEParserEXCLUDE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(193)
			p.Match(ADEParserCLASS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(194)
			p.Match(ADEParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewExcludeClassImplementingContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(195)
			p.Match(ADEParserEXCLUDE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(196)
			p.Match(ADEParserCLASS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(197)
			p.Match(ADEParserIMPLEMENTING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)
			p.Match(ADEParserINTERFACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(199)
			p.Match(ADEParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewExcludeComponentContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(200)
			p.Match(ADEParserEXCLUDE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(201)
			p.Match(ADEParserCOMPONENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(202)
			p.Match(ADEParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewExcludePatternContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(203)
			p.Match(ADEParserEXCLUDE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(204)
			p.Match(ADEParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISeverityStmtContext is an interface to support dynamic dispatch.
type ISeverityStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SEVERITY() antlr.TerminalNode
	SeverityValue() ISeverityValueContext

	// IsSeverityStmtContext differentiates from other interfaces.
	IsSeverityStmtContext()
}

type SeverityStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySeverityStmtContext() *SeverityStmtContext {
	var p = new(SeverityStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_severityStmt
	return p
}

func InitEmptySeverityStmtContext(p *SeverityStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_severityStmt
}

func (*SeverityStmtContext) IsSeverityStmtContext() {}

func NewSeverityStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SeverityStmtContext {
	var p = new(SeverityStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADEParserRULE_severityStmt

	return p
}

func (s *SeverityStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SeverityStmtContext) SEVERITY() antlr.TerminalNode {
	return s.GetToken(ADEParserSEVERITY, 0)
}

func (s *SeverityStmtContext) SeverityValue() ISeverityValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISeverityValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISeverityValueContext)
}

func (s *SeverityStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SeverityStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SeverityStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitSeverityStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADEParser) SeverityStmt() (localctx ISeverityStmtContext) {
	localctx = NewSeverityStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ADEParserRULE_severityStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(207)
		p.Match(ADEParserSEVERITY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(208)
		p.SeverityValue()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISeverityValueContext is an interface to support dynamic dispatch.
type ISeverityValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ERROR() antlr.TerminalNode
	WARNING() antlr.TerminalNode

	// IsSeverityValueContext differentiates from other interfaces.
	IsSeverityValueContext()
}

type SeverityValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySeverityValueContext() *SeverityValueContext {
	var p = new(SeverityValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_severityValue
	return p
}

func InitEmptySeverityValueContext(p *SeverityValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ADEParserRULE_severityValue
}

func (*SeverityValueContext) IsSeverityValueContext() {}

func NewSeverityValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SeverityValueContext {
	var p = new(SeverityValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADEParserRULE_severityValue

	return p
}

func (s *SeverityValueContext) GetParser() antlr.Parser { return s.parser }

func (s *SeverityValueContext) ERROR() antlr.TerminalNode {
	return s.GetToken(ADEParserERROR, 0)
}

func (s *SeverityValueContext) WARNING() antlr.TerminalNode {
	return s.GetToken(ADEParserWARNING, 0)
}

func (s *SeverityValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SeverityValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SeverityValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADEVisitor:
		return t.VisitSeverityValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADEParser) SeverityValue() (localctx ISeverityValueContext) {
	localctx = NewSeverityValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ADEParserRULE_severityValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ADEParserERROR || _la == ADEParserWARNING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
