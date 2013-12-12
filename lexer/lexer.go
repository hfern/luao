package lexer

import (
	lex "github.com/iNamik/go_lexer"
	"io"
)

type TokenStream <-chan Token

type Lexer interface {
	lex.Lexer
	Stream() TokenStream
}

type lexer struct {
	lex.Lexer
	tokstream chan Token
}

// New prepares a new Lua code lexer.
func New(reader io.Reader) Lexer {
	ll := lex.New(stStart, reader, 1)
	return wrapLexer(ll)
}

func NewFromString(sourceCode string) Lexer {
	ll := lex.NewFromString(stStart, sourceCode, 1)
	return wrapLexer(ll)
}

func NewFromBytes(sourceCode []byte) Lexer {
	ll := lex.NewFromBytes(stStart, sourceCode, 1)
	return wrapLexer(ll)
}

func wrapLexer(ll lex.Lexer) Lexer {
	return &lexer{Lexer: ll}
}

func (ll *lexer) Stream() TokenStream {
	if ll.tokstream == nil {
		ll.tokstream = make(chan Token, 1)
		go func() {

			for t := ll.NextToken(); &t != nil; t = ll.NextToken() {

				ll.tokstream <- Token{Token: *t}
				if TokenType(t.Type()) == EOF {
					break
				}
			}
			close(ll.tokstream)
		}()
	}
	return TokenStream(ll.tokstream)
}

func (ll *lexer) EmitReal(t TokenType) {
	ll.Lexer.EmitTokenWithBytes(lex.TokenType(t))
}
