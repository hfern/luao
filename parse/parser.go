package parse

import (
	"github.com/hfern/luao/lexer"
)

type Parser interface {
	Parse(lexer.TokenStream) Tree
}

func NewParser() Parser {
	return &parser{}
}

type parser struct {
}

func (p *parser) Parse(lexer.TokenStream) Tree {
	return Tree{}
}
