package lexer

import (
	"fmt"
	lex "github.com/iNamik/go_lexer"
	"unicode"
)

func emit(l lex.Lexer, t TokenType) {
	l.EmitTokenWithBytes(lex.TokenType(t))
}

func stStart(l lex.Lexer) lex.StateFn {
	char := l.PeekRune(0)

	switch true {
	case char == lex.RuneEOF:
		l.EmitEOF()
		return nil
	// Read comments
	case char == '-' && l.PeekRune(1) == '-':
		return stComment
	// Double-quote string
	case char == '"' || char == '\'':
		return stReadString
	// Read 3char symbols
	case inMapR(char, symbols_3c_1) && inMap(concatRunes(char, l.PeekRune(1), l.PeekRune(2)), symbols_3c):
		l.NextRune()
		l.NextRune()
		l.NextRune()
		emit(l, Symbol)
		return stStart
	// 2char symbols
	case inMapR(char, symbols_2c_1) && inMap(concatRunes(char, l.PeekRune(1)), symbols_2c):
		l.NextRune()
		l.NextRune()
		emit(l, Symbol)
		return stStart
	// 1char symbols
	case inMapR(char, symbols_1c):
		l.NextRune()
		emit(l, Symbol)
		return stStart
	// Read a Name identifier
	case inMapR(char, nameFirstChar):
		return stReadName
	// Read whitespace
	case inMapR(char, whitespaceChars):
		emit(l, Whitespace)
		return stWhitespace
	// Read Long strings
	case longStringEquals('[', l) >= 0:
		return stLongString
	// Read hex numbers
	case char == '0' && (unicode.ToLower(l.PeekRune(1)) == 'x'):
		return stHex
	// Read normal numbers
	case inMapR(char, digits):
		return stNumber
	}
	return nil
}

func stWhitespace(l lex.Lexer) lex.StateFn {
	for next := l.PeekRune(0); inMapR(next, whitespaceChars); next = l.PeekRune(0) {
		l.NextRune()
		if next == '\n' {
			l.NewLine()
		}
	}
	emit(l, Whitespace)
	return stStart
}

func stNumber(l lex.Lexer) lex.StateFn {
	if !l.MatchOneOrMoreBytes(digits_b) {
		emitError("Malformatted number: no principle number.")
	}
	if l.MatchOneRune('.') {
		if !l.MatchOneOrMoreBytes(digits_b) {
			emitError("Malformatted number: no fractional part in float.")
		}
	}
	if l.MatchOneRunes([]rune{'e', 'E'}) {
		l.MatchZeroOrOneRunes([]rune{'-', '+'})
		if !l.MatchOneOrMoreBytes(digits_b) {
			emitError("Malformatted number: no exponential part in float.")
		}
	}
	emit(l, Number)
	return stStart
}

func stHex(l lex.Lexer) lex.StateFn {
	if l.NextRune() != '0' {
		emitError("Hex numbers should start with 0")
		return nil
	}
	if unicode.ToLower(l.NextRune()) != 'x' {
		emitError("Malformatted hex number (no x)")
		return nil
	}

	baseChars := 0
	for ; inMapR(l.PeekRune(0), hexdigits); baseChars++ {
		l.NextRune()
	}

	fmt.Println("Base Chars:", baseChars)

	if l.PeekRune(0) == '.' {
		l.NextRune()
		if !l.MatchOneOrMoreBytes(hexdigits_b) {
			emitError("Malformed hex number.")
			return nil
		}
	} else if baseChars == 0 {
		emitError("Malformed hex number (no base chars or decimal).")
		return nil
	}

	if unicode.ToLower(l.PeekRune(0)) == 'p' {
		l.NextRune()

		next := l.PeekRune(0)
		if next == '+' || next == '-' {
			l.NextRune()

		}
		if !l.MatchOneOrMoreBytes(digits_b) {
			emitError("Malformed hex number (no digits after p)")
			return nil
		}
	}

	emit(l, HexNumber)
	return stStart
}

func stComment(l lex.Lexer) lex.StateFn {
	isComment := '-' == l.NextRune() && '-' == l.NextRune()
	if !isComment {
		emitError("Comment parser called without comment.")
		return nil
	}

	// check if this is going to be a block comment
	if longStringEquals('[', l) >= 0 {
		// consume the string portion
		if !readLongString(l) {
			emitError("Error reading block-comment contents")
			return nil
		}
	} else {
		// read til \n is found
		l.NonMatchZeroOrMoreRunes([]rune{'\n', lex.RuneEOF})
	}

	emit(l, Comment)
	return stStart
}

func stReadString(l lex.Lexer) lex.StateFn {
	delim := l.NextRune()
	if !(delim == '\'' || delim == '"') {
		emitError("Expected `\"` or `'` to start quote-string.")
	}
	for {
		char := l.NextRune()
		switch char {
		case delim:
			// end of string
			fmt.Println("End of the string!")
			emit(l, String)
			return stStart
		case '\\':
			l.NextRune()
		case lex.RuneEOF:
			l.EmitEOF()
		case '\n':
			l.NewLine()
		}
	}
	return stStart
}

// TODO(hunter): add keywords for reserved name
func stReadName(l lex.Lexer) lex.StateFn {
	if !inMapR(l.NextRune(), nameFirstChar) {
		emitError("Name does not start with valid character")
		return nil
	}
	l.MatchZeroOrMoreBytes(nameAfterChars_b)
	emit(l, Name)
	return stStart
}

// longStringEquals peeks the bytestream to see if
// a long string beginning is directly ahead. It returns the amount
// of equals the long string has, or -1 if
// no string found
func longStringEquals(delim rune, l lex.Lexer) int {
	if l.PeekRune(0) != delim {
		return -1
	}
	equalsCt := 0
	for l.PeekRune(equalsCt+1) == '=' {
		equalsCt++
	}

	if l.PeekRune(equalsCt+1) != delim {
		return -1
	}
	return equalsCt
}

func stLongString(l lex.Lexer) lex.StateFn {
	if !readLongString(l) {
		emitError("Error reading string.")
	}
	emit(l, LongString)
	return stStart

}

func readLongString(l lex.Lexer) bool {
	numEquals := longStringEquals('[', l)
	if numEquals < 0 {
		emitError("Expecting [ to start string.")
		return false
	}
	// eat first [
	l.NextRune()

	for i := 0; i < numEquals; i++ {
		if l.NextRune() != '=' {
			emitError("Expected string `=` padding.")
			return false
		}
	}

	// gobble string contents
	for {
		char := l.NextRune()
		if char == lex.RuneEOF {
			emitError("Unexpected EOF. Expecting long-string close.")
			return false
		}

		if char == '\n' {
			l.NewLine()
		}

		// potential close
		if char == ']' {
			// ungobble first ]
			l.BackupRune()
			closingCt := longStringEquals(']', l)
			if closingCt != numEquals {
				// re-gobble the ]
				l.NextRune()
			} else {
				// it was our close!
				// eat the closing bits
				l.NextRune()
				for i := 0; i < numEquals; i++ {
					if l.NextRune() != '=' {
						emitError("Expected close-string `=` padding.")
					}
				}
				l.NextRune()
				return true
			}
		}

	}

	return false
}
