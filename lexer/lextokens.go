package lexer

import (
	lex "github.com/iNamik/go_lexer"
)

type TokenType lex.TokenType

const (
	EOF        TokenType = TokenType(lex.TokenTypeEOF)
	META_ERROR           = EOF + iota
	Keyword
	Symbol
	Name
	Comment
	Number
	Float
	HexNumber
	String
	LongString
	Whitespace

	_meta_KeywordRangeStart
	KW_And
	KW_Break
	KW_Do
	KW_Else
	KW_Elseif
	KW_End
	KW_False
	KW_For
	KW_Function
	KW_Goto
	KW_If
	KW_In
	KW_Local
	KW_Nil
	KW_Not
	KW_Or
	KW_Repeat
	KW_Return
	KW_Then
	KW_True
	KW_Until
	KW_While
	_meta_KeywordRangeEnd

	_meta_SymbolRangeStart
	S_PLUS
	S_MINUS
	S_MULTIPLY_STAR
	S_DIVISION_BAR
	S_POWER_CHEVRON
	S_PERCENTAGE
	S_COMMA
	S_CURLY_OPEN
	S_CURLY_CLOSE
	S_BRACKET_OPEN
	S_BRACKER_CLOSE
	S_PAREN_OPEN
	S_PAREN_CLOSE
	S_SEMICOLON
	S_POUND
	S_EQUALS

	S_DOUBLE_EQUALS
	S_NOTEQUALS
	S_LESSEQUALS
	S_GREATEREQUALS
	S_LESS_THAN
	S_GREATER_THAN

	S_COLON
	S_DOUBLE_COLON
	S_PERIOD
	S_PERIOD_2
	S_PERIOD_3
	_meta_SymbolRangeEnd
)

var TokenNames = map[TokenType]string{
	EOF:        "EOF",
	META_ERROR: "META_ERROR",
	Keyword:    "Keyword",
	Symbol:     "Symbol",
	Name:       "Name",
	Comment:    "Comment",
	Number:     "Number",
	Float:      "Float",
	HexNumber:  "HexNumber",
	String:     "String",
	LongString: "LongString",
	Whitespace: "Whitespace",

	KW_And:      "Keyword 'and'",
	KW_Break:    "Keyword 'break'",
	KW_Do:       "Keyword 'do'",
	KW_Else:     "Keyword 'else'",
	KW_Elseif:   "Keyword 'elseif'",
	KW_End:      "Keyword 'end'",
	KW_False:    "Keyword 'false'",
	KW_For:      "Keyword 'for'",
	KW_Function: "Keyword 'function'",
	KW_Goto:     "Keyword 'goto'",
	KW_If:       "Keyword 'if'",
	KW_In:       "Keyword 'in'",
	KW_Local:    "Keyword 'local'",
	KW_Nil:      "Keyword 'nil'",
	KW_Not:      "Keyword 'not'",
	KW_Or:       "Keyword 'or'",
	KW_Repeat:   "Keyword 'repeat'",
	KW_Return:   "Keyword 'return'",
	KW_Then:     "Keyword 'then'",
	KW_True:     "Keyword 'true'",
	KW_Until:    "Keyword 'until'",
	KW_While:    "Keyword 'while'",

	S_PLUS:          "S_PLUS",
	S_MINUS:         "S_MINUS",
	S_MULTIPLY_STAR: "S_MULTIPLY_STAR",
	S_DIVISION_BAR:  "S_DIVISION_BAR",
	S_POWER_CHEVRON: "S_POWER_CHEVRON",
	S_PERCENTAGE:    "S_PERCENTAGE",
	S_COMMA:         "S_COMMA",
	S_CURLY_OPEN:    "S_CURLY_OPEN",
	S_CURLY_CLOSE:   "S_CURLY_CLOSE",
	S_BRACKET_OPEN:  "S_BRACKET_OPEN",
	S_BRACKER_CLOSE: "S_BRACKER_CLOSE",
	S_PAREN_OPEN:    "S_PAREN_OPEN",
	S_PAREN_CLOSE:   "S_PAREN_CLOSE",
	S_SEMICOLON:     "S_SEMICOLON",
	S_POUND:         "S_POUND",
	S_EQUALS:        "S_EQUALS",

	S_DOUBLE_EQUALS: "S_DOUBLE_EQUALS",
	S_NOTEQUALS:     "S_NOTEQUALS",
	S_LESSEQUALS:    "S_LESSEQUALS",
	S_GREATEREQUALS: "S_GREATEREQUALS",
	S_LESS_THAN:     "S_LESS_THAN",
	S_GREATER_THAN:  "S_GREATER_THAN",

	S_COLON:        "S_COLON,",
	S_DOUBLE_COLON: "S_DOUBLE_COLON,",
	S_PERIOD:       "S_PERIOD,",
	S_PERIOD_2:     "S_PERIOD_2,",
	S_PERIOD_3:     "S_PERIOD_3,",
}

// RangeKeywords returns the numberical range of keywords
// start < keywords < stop
func RangeKeywords() (start TokenType, stop TokenType) {
	return TokenType(_meta_KeywordRangeStart), TokenType(_meta_KeywordRangeEnd)
}
