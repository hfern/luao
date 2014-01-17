package parse

import (
	"github.com/hfern/luao/lexer"
)

type readerState struct {
	read int
}

type tReader struct {
	state  readerState
	source *[]lexer.Token
}

func (r tReader) Save() readerState {
	return r.state
}

func (r *tReader) Restore(s readerState) {
	r.state = s
}

func (r *tReader) Next() lexer.Token {
	if len(r.source) <= r.state.read {
		return lexer.Token
	}
}
