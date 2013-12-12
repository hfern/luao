package lexer

import (
	lex "github.com/iNamik/go_lexer"
)

type Token struct {
	lex.Token
}

// TODO(hunter): Maybe cache this for performance?
func (t *Token) Type() TokenType {
	tt := TokenType(t.Token.Type())
	// Check for keywords or symbols
	if tt == Name || tt == Symbol {
		text := string(t.Bytes())
		if tt == Name {
			if inMap(text, keywords) {
				return keywords[text].(TokenType)
			}
		} else if tt == Symbol {
			if inMap(text, symbols) {
				return symbols[text].(TokenType)
			}
		}
	}
	return tt
}
