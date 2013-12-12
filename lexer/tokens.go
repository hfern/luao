package lexer

import (
	"bytes"
	"unicode/utf8"
)

type charMap map[string]interface{}

var whitespaceChars = toMap(" ", "\n", "\t", "\r")
var whitespaceChars_b = toBArray(whitespaceChars)

var escapeMap map[string]string = map[string]string{
	"\r": "\\r",
	"\t": "\\t",
	"\n": "\\n",
	"\"": "\\\"",
	"'":  "\\'",
}

var lowerChars charMap = toMap(
	"a", "b", "c", "d", "e", "f", "g", "h", "i",
	"j", "k", "l", "m", "n", "o", "p", "q", "r",
	"s", "t", "u", "v", "w", "x", "y", "z",
)

var upperChars charMap = toMap(
	"A", "B", "C", "D", "E", "F", "G", "H", "I",
	"J", "K", "L", "M", "N", "O", "P", "Q", "R",
	"S", "T", "U", "V", "W", "X", "Y", "Z",
)

var digits charMap = toMap(
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
)
var digits_b = toBArray(digits)

var hexdigits charMap = unionMap(toMap(
	"A", "a", "B", "b", "C", "c", "D", "d", "E", "e", "F", "f",
), digits)
var hexdigits_b = toBArray(hexdigits)

// single character symbols
var symbols_1c charMap = charMap{
	"+": S_PLUS,
	"-": S_MINUS,
	"*": S_MULTIPLY_STAR,
	"/": S_DIVISION_BAR,
	"^": S_POWER_CHEVRON,
	"%": S_PERCENTAGE,
	",": S_COMMA,
	"{": S_CURLY_OPEN,
	"}": S_CURLY_CLOSE,
	"[": S_BRACKET_OPEN,
	"]": S_BRACKER_CLOSE,
	"(": S_PAREN_OPEN,
	")": S_PAREN_CLOSE,
	";": S_SEMICOLON,
	"#": S_POUND,
	"=": S_EQUALS,
	"<": S_LESS_THAN,
	">": S_GREATER_THAN,
	":": S_COLON,
	".": S_PERIOD,
}

// double character symbols
var symbols_2c charMap = charMap{
	"==": S_DOUBLE_EQUALS,
	"~=": S_NOTEQUALS,
	"<=": S_LESSEQUALS,
	">=": S_GREATEREQUALS,
	"::": S_DOUBLE_COLON,
	"..": S_PERIOD_2,
}
var symbols_2c_1 charMap = cacheFirstChar(symbols_2c)

// triple character symbols
var symbols_3c charMap = charMap{
	"...": S_PERIOD_3,
}
var symbols_3c_1 charMap = cacheFirstChar(symbols_3c)

var symbols charMap = unionMap(symbols_1c, symbols_2c, symbols_3c)

var keywords charMap = charMap{
	"and":      KW_And,
	"break":    KW_Break,
	"do":       KW_Do,
	"else":     KW_Else,
	"elseif":   KW_Elseif,
	"end":      KW_End,
	"false":    KW_False,
	"for":      KW_For,
	"function": KW_Function,
	"goto":     KW_Goto,
	"if":       KW_If,
	"in":       KW_In,
	"local":    KW_Local,
	"nil":      KW_Nil,
	"not":      KW_Not,
	"or":       KW_Or,
	"repeat":   KW_Repeat,
	"return":   KW_Return,
	"then":     KW_Then,
	"true":     KW_True,
	"until":    KW_Until,
	"while":    KW_While,
}

var nameFirstChar = unionMap(lowerChars, upperChars, toMap("_"))
var nameFirstChar_b = toBArray(nameFirstChar)
var nameAfterChars = unionMap(lowerChars, upperChars, digits, toMap("_"))
var nameAfterChars_b = toBArray(nameAfterChars)

func toMap(tokens ...string) charMap {
	cmap := make(charMap, len(tokens))
	for _, tok := range tokens {
		cmap[tok] = true
	}
	return cmap
}

func unionMap(maps ...charMap) charMap {
	base := make(charMap)

	for _, mp := range maps {
		for key, val := range mp {
			base[key] = val
		}
	}
	return base
}

func inMap(needle string, haystack charMap) bool {
	_, ok := haystack[needle]
	return ok
}

func inMapR(needle rune, haystack charMap) bool {
	return inMap(string(needle), haystack)
}

func toArray(mp charMap) []string {
	arr := make([]string, 0, len(mp))
	for key, _ := range mp {
		arr = append(arr, key)
	}
	return arr
}

func toBArray(mp charMap) []byte {
	arr := make([]byte, 0, len(mp))
	for key, _ := range mp {
		firstByte := ([]byte(key))[0]
		arr = append(arr, firstByte)
	}
	return arr
}

// Cache the first first character of each
// element in the charmap to allow quick lookups
func cacheFirstChar(mp charMap) charMap {
	cache := make(charMap, len(mp))
	for key, _ := range mp {
		rn, _ := utf8.DecodeRuneInString(key)
		cache[string(rn)] = true
	}
	return cache
}

func concatRunes(rns ...rune) string {
	var buffer bytes.Buffer
	for _, rn := range rns {
		buffer.WriteRune(rn)
	}
	return buffer.String()
}
