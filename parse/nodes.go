package parse

import (
	"github.com/hfern/luao/lexer"
)

type branch interface{}

type production struct {
	Name     RuleId
	Children []branch
}

type terminal lexer.Token
